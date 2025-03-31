package jobs

import (
	"github.com/makeopensource/leviathan/models"
	. "github.com/makeopensource/leviathan/service/file_manager"
	"os"
	"path/filepath"
	"testing"
	"time"
)

var (
	basePath        = "../../../example/tango"
	tangoDockerFile = basePath + "/tango-Dockerfile"
	autolab0        = basePath + "/tango0"
	autolab1        = basePath + "/tango1"
	//autolab2        = basePath + "/tango2"
	autolab3     = basePath + "/tango3"
	autolab4     = basePath + "/tango4"
	tangoTimeout = 10 * time.Second
)

type testMap = map[string]testCase

var (
	tangoTestCases = map[string]testMap{
		"tango0": {
			"correct": {
				studentFile:    autolab0 + "/handin.py",
				expectedOutput: `{"scores": {"q1": 10, "q2": 10, "q3": 10}}`,
				correctStatus:  models.Complete,
			},
			"cheating1": testCase{
				studentFile:    autolab0 + "/handin_cheating1.py",
				expectedOutput: `{"scores": {"q1": 0, "q2": 0, "q3": 0}}`,
				correctStatus:  models.Complete,
			},
			"cheating2": testCase{
				studentFile:    autolab0 + "/handin_cheating2.py",
				expectedOutput: `Maximum timeout reached for job, job ran for 10s`,
				correctStatus:  models.Failed,
			},
			"incorrect1": testCase{
				studentFile:    autolab0 + "/handin_incorrect1.py",
				expectedOutput: `{"scores": {"q1": 0, "q2": 0, "q3": 0}}`,
				correctStatus:  models.Complete,
			},
			"incorrect2": testCase{
				studentFile:    autolab0 + "/handin_incorrect2.py",
				expectedOutput: `unable to parse log output`,
				correctStatus:  models.Failed,
			},
		},
		"tango1": {
			"correct": {
				studentFile:    autolab1 + "/handin.py",
				expectedOutput: `{"scores": {"q1": 10, "q2": 10, "q3": 10}}`,
				correctStatus:  models.Complete,
			},
			"incorrect1": testCase{
				studentFile:    autolab1 + "/handin_incorrect1.py",
				expectedOutput: `{"scores": {"q1": 10, "q2": 0, "q3": 0}}`,
				correctStatus:  models.Complete,
			},
			"incorrect2": testCase{
				studentFile:    autolab1 + "/handin_incorrect2.py",
				expectedOutput: `{"scores": {"q1": 9, "q2": 3, "q3": 3}}`,
				correctStatus:  models.Complete,
			},
			"incorrect3": testCase{
				studentFile:    autolab1 + "/handin_incorrect3.py",
				expectedOutput: `{"scores": {"q1": 1, "q2": 0, "q3": 0}}`,
				correctStatus:  models.Complete,
			},
			"incorrect4": testCase{
				studentFile:    autolab1 + "/handin_incorrect4.py",
				expectedOutput: `{"scores": {"q1": 0, "q2": 0, "q3": 0}}`,
				correctStatus:  models.Complete,
			},
		},
		"tango3": {
			"correct": {
				studentFile:    autolab3 + "/handin.json",
				expectedOutput: `{"scores": {"q1": 10, "q2": 10, "q3": 99}}`,
				correctStatus:  models.Complete,
			},
		},
		"tango4": {
			"correct": {
				studentFile:    autolab4 + "/handin.py",
				expectedOutput: `{"scores": {"q1": 10}}`,
				correctStatus:  models.Complete,
			},
			"incorrect": {
				studentFile:    autolab4 + "/handin_incorrect.py",
				expectedOutput: `{"scores": {"q1": 0}}`,
				correctStatus:  models.Complete,
			},
		},
	}
	tangoTestFuncs = map[string]func(*testing.T){
		"autolab0": TestTango0,
		"autolab1": TestTango1,
		"autolab3": TestTango3,
		"autolab4": TestTango4,
	}
)

func TestAllTango(t *testing.T) {
	for tCase, test := range tangoTestFuncs {
		t.Run(tCase, func(t *testing.T) {
			t.Parallel()
			test(t)
		})
	}
}

func TestTango0(t *testing.T) {
	setupTest()

	folderName := "tango0"
	cases := tangoTestCases[folderName]

	labId := newLab(t, autolab0)

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			jobid := setupJobProcessTango(t, labId, test.studentFile)
			testJob(t, jobid, test.expectedOutput, test.correctStatus)
		})
	}
}

func TestTango1(t *testing.T) {
	setupTest()

	folderName := "tango1"
	cases := tangoTestCases[folderName]
	labId := newLab(t, autolab1)

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			jobid := setupJobProcessTango(t, labId, test.studentFile)
			testJob(t, jobid, test.expectedOutput, test.correctStatus)
		})
	}
}

func TestTango3(t *testing.T) {
	setupTest()

	folderName := "tango3"
	cases := tangoTestCases[folderName]
	labId := newLab(t, autolab3)

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			jobid := setupJobProcessTango(t, labId, test.studentFile)
			testJob(t, jobid, test.expectedOutput, test.correctStatus)
		})
	}
}

func TestTango4(t *testing.T) {
	setupTest()

	folderName := "tango4"
	cases := tangoTestCases[folderName]
	labId := newLab(t, autolab4)

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			jobid := setupJobProcessTango(t, labId, test.studentFile)
			testJob(t, jobid, test.expectedOutput, test.correctStatus)
		})
	}
}

func newLab(t *testing.T, folderName string) uint {
	tarPath := folderName + "/autograde.tar"
	tangoMakeFilePath := folderName + "/Makefile"
	labId := createLab(t, &models.Lab{
		Name:              "tango-test-lab",
		JobTimeout:        tangoTimeout,
		AutolabCompatible: true,
	}, tangoDockerFile, tarPath, tangoMakeFilePath)
	if labId == 0 {
		t.Fatalf("Failed to create lab")
	}
	return labId
}

func setupJobProcessTango(t *testing.T, labId uint, studentCodePath string) string {
	newJob := &models.Job{LabID: labId}

	studentFileName := filepath.Base(studentCodePath)
	if filepath.Ext(studentFileName) == ".json" {
		studentFileName = "handin.json"
	} else {
		studentFileName = "handin.py"
	}

	studentCode, err := os.Open(studentCodePath)
	if err != nil {
		t.Fatal(err)
	}

	folderId, err := fileManTestService.CreateSubmissionFolder(&FileInfo{
		Reader:   studentCode,
		Filename: studentFileName,
	})
	if err != nil {
		t.Fatal(err)
		return ""
	}

	jobId, err := jobTestService.NewJob(
		newJob,
		folderId,
	)

	if err != nil {
		t.Fatal("unable to create job", err)
	}

	return jobId
}
