package jobs

import (
	"context"
	"fmt"
	"github.com/makeopensource/leviathan/common"
	v1 "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const (
	makeFilePath   = "../../../example/simple-addition/makefile"
	graderFilePath = "../../../example/simple-addition/grader.py"
)

var (
	defaultTimeout = time.Second * 10
	testCases      = map[string]struct {
		studentFile    string
		expectedOutput string
		correctStatus  models.JobStatus
	}{
		"correct": {
			studentFile:    "../../../example/simple-addition/student_correct.py",
			expectedOutput: `{"addition": {"passed": true, "message": ""}, "subtraction": {"passed": true, "message": ""}, "multiplication": {"passed": true, "message": ""}, "division": {"passed": true, "message": ""}}`,
			correctStatus:  models.Complete,
		},
		"incorrect": {
			studentFile:    "../../../example/simple-addition/student_incorrect.py",
			expectedOutput: `{"addition": {"passed": true, "message": ""}, "subtraction": {"passed": true, "message": ""}, "multiplication": {"passed": false, "message": "Multiplication failed. Expected 42, got 48"}, "division": {"passed": false, "message": "Division failed. Expected 4, got 3.3333333333333335"}}`,
			correctStatus:  models.Complete,
		},
		"timeout": {
			studentFile:    "../../../example/simple-addition/student_timeout.py",
			expectedOutput: "Maximum timeout reached for job, job ran for 10s",
			correctStatus:  models.Failed,
		},
		//"forkb": {
		//	studentFile:    "../../../example/simple-addition/student_fork_bomb.py",
		//	expectedOutput: "",
		//},
	}
)

func TestCorrect(t *testing.T) {
	SetupTest()
	correct := testCases["correct"]
	testJobProcessor(t, correct.studentFile, correct.expectedOutput, defaultTimeout, correct.correctStatus)
}

func TestIncorrect(t *testing.T) {
	SetupTest()
	incorrect := testCases["incorrect"]
	testJobProcessor(t, incorrect.studentFile, incorrect.expectedOutput, defaultTimeout, incorrect.correctStatus)
}

// TODO
func TestForkBomb(t *testing.T) {
	SetupTest()
	forkBomb := testCases["forkb"]
	testJobProcessor(t, forkBomb.studentFile, forkBomb.expectedOutput, defaultTimeout, forkBomb.correctStatus)
}

func TestTimeout(t *testing.T) {
	SetupTest()
	timeLimit := time.Second * 10
	timeout := testCases["timeout"]
	timeout.expectedOutput = fmt.Sprintf("Maximum timeout reached for job, job ran for %s", timeLimit)
	testJobProcessor(t, timeout.studentFile, timeout.expectedOutput, timeLimit, timeout.correctStatus)
}

func TestCancel(t *testing.T) {
	SetupTest()
	timeLimit := time.Second * 10
	timeout := testCases["timeout"]
	timeout.expectedOutput = fmt.Sprintf("Job was cancelled")

	jobId := setupJobProcess(timeout.studentFile, timeLimit)

	// cancel the job after 3 seconds
	time.AfterFunc(3*time.Second, func() {
		JobTestService.CancelJob(jobId)
	})

	testJob(t, jobId, timeout.expectedOutput, models.Canceled)

	// verify cancel function was removed from context map
	value := JobTestService.queue.GetJobCancelFunc(jobId)
	if value != nil {
		t.Fatalf("Job was cancelled, but the cancel func was not nil")
	}

}

func testJobProcessor(t *testing.T, studentCodePath string, correctOutput string, timeout time.Duration, status models.JobStatus) {
	jobId := setupJobProcess(studentCodePath, timeout)
	testJob(t, jobId, correctOutput, status)
}

func setupJobProcess(studentCodePath string, timeout time.Duration) string {
	graderBytes, err := common.ReadFileBytes(graderFilePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading grader.py")
	}

	makefileBytes, err := common.ReadFileBytes(makeFilePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading grader.py")
	}

	studentBytes, err := common.ReadFileBytes(studentCodePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading student")
	}

	dockerBytes, err := common.ReadFileBytes(DockerFilePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading docker file")
	}

	newJob := &models.Job{
		LabData:    models.LabModel{ImageTag: ImageName},
		JobTimeout: timeout,
	}

	jobId, err := JobTestService.NewJob(
		newJob,
		&v1.FileUpload{
			Filename: filepath.Base(makeFilePath),
			Content:  makefileBytes,
		},
		&v1.FileUpload{
			Filename: filepath.Base(graderFilePath),
			Content:  graderBytes,
		},
		&v1.FileUpload{
			Filename: "student.py",
			Content:  studentBytes,
		},
		&v1.FileUpload{
			Filename: filepath.Base(DockerFilePath),
			Content:  dockerBytes,
		},
	)

	if err != nil {
		log.Fatal().Err(err).Msgf("Error creating job")
	}

	return jobId
}

func testJob(t *testing.T, jobId string, correctOutput string, correctStatus models.JobStatus) {
	jobInfo, logs, err := JobTestService.WaitForJobAndLogs(context.Background(), jobId)
	if err != nil {
		t.Fatalf("Error waiting for job: %v", err)
		return
	}

	log.Debug().Msgf("Job ID: %s, Logs: %s", jobId, logs)

	returned := strings.TrimSpace(jobInfo.StatusMessage)
	expected := strings.TrimSpace(correctOutput)

	assert.Equal(t, expected, returned)
	assert.Equal(t, correctStatus, jobInfo.Status)
}
