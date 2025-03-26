package tests

import (
	"github.com/makeopensource/leviathan/common"
	v1 "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/service/jobs"
	"github.com/makeopensource/leviathan/service/labs"
	"github.com/stretchr/testify/assert"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	dkTestService  *docker.DkService
	jobTestService *jobs.JobService
	labTestService *labs.LabService
	setupOnce      sync.Once
)

const (
	imageName      = "arithmetic-python"
	dockerFilePath = "../../example/simple-addition/ex-Dockerfile"
	makeFilePath   = "../../example/simple-addition/makefile"
	graderFilePath = "../../example/simple-addition/grader.py"
)

// creates a lab and send a job
func TestLabJob(t *testing.T) {
	err := os.RemoveAll("appdata")
	if err != nil {
		t.Fatal(err)
	}
	initDeps()

	labId := createLab(t)
	testJobProcessor(t, "../../example/simple-addition/student_correct.py", labId)
}

func testJobProcessor(t *testing.T, studentCodePath string, labId uint) {
	jobId := setupJobProcess(studentCodePath, labId)
	testJob(t, jobId, "", "")
}

func setupJobProcess(studentCodePath string, labId uint) string {
	studentBytes, err := os.ReadFile(studentCodePath)
	if err != nil {
		log.Fatal("Error reading student", err)
	}

	newJob := &models.Job{LabID: labId}
	jobId, err := jobTestService.NewJobFromLab(
		newJob,
		[]*v1.FileUpload{
			{
				Filename: "student.py",
				Content:  studentBytes,
			},
		},
	)
	if err != nil {
		log.Fatal("Error creating job", err)
	}

	return jobId
}

func testJob(t *testing.T, jobId string, correctOutput string, correctStatus models.JobStatus) {
	jobInfo, logs, err := jobTestService.WaitForJobAndLogs(jobId)
	if err != nil {
		t.Fatalf("Error waiting for job: %v", err)
		return
	}

	t.Log("Job ID: ", jobId, " Logs: ", logs)

	returned := strings.TrimSpace(jobInfo.StatusMessage)
	expected := strings.TrimSpace(correctOutput)

	assert.Equal(t, expected, returned)
	assert.Equal(t, correctStatus, jobInfo.Status)
}

func createLab(t *testing.T) uint {
	graderBytes, err := os.ReadFile(graderFilePath)
	if err != nil {
		t.Fatalf("Error reading grader.py: %v", err)
		return 0
	}
	makefileBytes, err := os.ReadFile(makeFilePath)
	if err != nil {
		t.Fatalf("Error reading grader.py: %v", err)
		return 0
	}
	dockerBytes, err := os.ReadFile(dockerFilePath)
	if err != nil {
		t.Fatalf("Error reading docker file: %v", err)
		return 0
	}

	jobFiles := []*v1.FileUpload{
		{
			Filename: filepath.Base(makeFilePath),
			Content:  makefileBytes,
		},
		{
			Filename: filepath.Base(graderFilePath),
			Content:  graderBytes,
		},
	}

	dockerFile := &v1.FileUpload{
		Filename: filepath.Base(dockerFilePath),
		Content:  dockerBytes,
	}

	lab := models.Lab{
		Name:        "test-lab",
		JobTimeout:  time.Second * 10,
		JobEntryCmd: "make grade",
	}

	createLab, err := labTestService.CreateLab(&lab, dockerFile, jobFiles)
	if err != nil {
		t.Fatalf("Error creating lab: %v", err)
		return 0
	}

	t.Logf("Created Lab: %v", createLab)
	return createLab
}

func initDeps() {
	setupOnce.Do(func() {
		common.InitConfig()
		db, bc := common.InitDB()

		dkTestService = docker.NewDockerServiceWithClients()
		labTestService = labs.NewLabService(db, dkTestService)
		jobTestService = jobs.NewJobService(db, bc, dkTestService, labTestService)

		// no logs on tests
		//log.Logger = log.Logger.Level(zerolog.Disabled)
	})
}
