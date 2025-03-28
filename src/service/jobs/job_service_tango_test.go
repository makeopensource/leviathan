package jobs

import (
	v1 "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/makeopensource/leviathan/models"
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
	autolab2        = basePath + "/tango2"
	autolab3        = basePath + "/tango3"
	autolab4        = basePath + "/tango4"
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
	}
)

func TestTango0(t *testing.T) {
	setupTest()

	folderName := "tango0"
	cases := tangoTestCases[folderName]

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			jobid := setupJobProcessTango(t, basePath+"/"+folderName, test.studentFile, 10*time.Second)
			testJob(t, jobid, test.expectedOutput, test.correctStatus)
		})
	}
}

func TestTango1(t *testing.T) {
	setupTest()

	folderName := "tango1"
	cases := tangoTestCases[folderName]

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			jobid := setupJobProcessTango(t, basePath+"/"+folderName, test.studentFile, 10*time.Second)
			testJob(t, jobid, test.expectedOutput, test.correctStatus)
		})
	}
}

func TestTango3(t *testing.T) {
	setupTest()

	folderName := "tango3"
	cases := tangoTestCases[folderName]

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			jobid := setupJobProcessTango(t, basePath+"/"+folderName, test.studentFile, 10*time.Second)
			testJob(t, jobid, test.expectedOutput, test.correctStatus)
		})
	}
}

func TestTango4(t *testing.T) {
	setupTest()

	folderName := "tango4"
	cases := tangoTestCases[folderName]

	for name, test := range cases {
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			jobid := setupJobProcessTango(t, basePath+"/"+folderName, test.studentFile, 10*time.Second)
			testJob(t, jobid, test.expectedOutput, test.correctStatus)
		})
	}
}

func setupJobProcessTango(t *testing.T, testFolder, studentCodePath string, timeout time.Duration) string {

	tarPath := testFolder + "/autograde.tar"
	graderTar, err := os.ReadFile(tarPath)
	if err != nil {
		t.Fatal(err)
	}
	tangoMakeFilePath := testFolder + "/Makefile"
	tangoMakeFilePathBytes, err := os.ReadFile(tangoMakeFilePath)
	if err != nil {
		t.Fatal(err)
	}

	studentFileName := filepath.Base(studentCodePath)
	if filepath.Ext(studentFileName) == ".json" {
		studentFileName = "handin.json"
	} else {
		studentFileName = "handin.py"
	}

	studentBytes, err := os.ReadFile(studentCodePath)
	if err != nil {
		t.Fatal(err)
	}

	dockerBytes, err := os.ReadFile(tangoDockerFile)
	if err != nil {
		t.Fatal(err)
	}

	newLab := models.Lab{
		JobTimeout: timeout,
		//JobEntryCmd: "ls -la ;while true; do sleep 1; done",
		ImageTag: "tango-test",
	}

	newJob := &models.Job{LabData: &newLab}

	jobId, err := jobTestService.NewJobFromRPC(
		newJob,
		[]*v1.FileUpload{
			{
				Filename: filepath.Base(tarPath),
				Content:  graderTar,
			},
			{
				Filename: filepath.Base(tangoMakeFilePath),
				Content:  tangoMakeFilePathBytes,
			},
			{
				Filename: studentFileName,
				Content:  studentBytes,
			},
		},
		&v1.FileUpload{
			Filename: filepath.Base(tangoDockerFile),
			Content:  dockerBytes,
		},
		true,
	)

	if err != nil {
		t.Fatal("unable to create job", err)
	}

	return jobId
}
