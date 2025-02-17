package jobs

import (
	"context"
	"fmt"
	"github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

const (
	makeFilePath   = "../../../example/python/simple-addition/makefile"
	graderFilePath = "../../../example/python/simple-addition/grader.py"
)

var (
	defaultTimeout = time.Second * 10
	testCases      = map[string]struct {
		studentFile    string
		expectedOutput string
	}{
		"correct": {
			studentFile:    "../../../example/python/simple-addition/student_correct.py",
			expectedOutput: `{"addition": {"passed": true, "message": ""}, "subtraction": {"passed": true, "message": ""}, "multiplication": {"passed": true, "message": ""}, "division": {"passed": true, "message": ""}}`,
		},
		"incorrect": {
			studentFile:    "../../../example/python/simple-addition/student_incorrect.py",
			expectedOutput: `{"addition": {"passed": true, "message": ""}, "subtraction": {"passed": true, "message": ""}, "multiplication": {"passed": false, "message": "Multiplication failed. Expected 42, got 48"}, "division": {"passed": false, "message": "Division failed. Expected 4, got 3.3333333333333335"}}`,
		},
		"timeout": {
			studentFile:    "../../../example/python/simple-addition/student_timeout.py",
			expectedOutput: "Maximum timeout reached for job, job ran for 10s",
		},
		//"forkb": {
		//	studentFile:    "../../../example/python/simple-addition/student_fork_bomb.py",
		//	expectedOutput: "",
		//},
	}
)

func TestCorrect(t *testing.T) {
	SetupTest()
	correct := testCases["correct"]
	testJobProcessor(t, correct.studentFile, correct.expectedOutput, defaultTimeout)
}

func TestIncorrect(t *testing.T) {
	SetupTest()
	incorrect := testCases["incorrect"]
	testJobProcessor(t, incorrect.studentFile, incorrect.expectedOutput, defaultTimeout)
}

// TODO
func TestForkBomb(t *testing.T) {
	SetupTest()
	forkBomb := testCases["forkb"]
	testJobProcessor(t, forkBomb.studentFile, forkBomb.expectedOutput, defaultTimeout)
}

func TestTimeout(t *testing.T) {
	SetupTest()
	timeLimit := time.Second * 10
	timeout := testCases["timeout"]
	timeout.expectedOutput = fmt.Sprintf("Maximum timeout reached for job, job ran for %s", timeLimit)
	testJobProcessor(t, timeout.studentFile, timeout.expectedOutput, timeLimit)
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

	testJob(t, jobId, timeout.expectedOutput)

	// verify cancel function was removed from context map
	value := JobTestService.queue.GetJobCancelFunc(jobId)
	if value != nil {
		t.Fatalf("Job was cancelled, but the cancel func was not nil")
	}

}

func testJobProcessor(t *testing.T, studentCodePath string, correctOutput string, timeout time.Duration) {
	jobId := setupJobProcess(studentCodePath, timeout)
	testJob(t, jobId, correctOutput)
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

	newJob := &models.Job{
		ImageTag:                  ImageName,
		StudentSubmissionFileName: "student.py",
		StudentSubmissionFile:     studentBytes,
		LabData: models.LabModel{
			GraderFilename: filepath.Base(graderFilePath),
			GraderFile:     graderBytes,
			MakeFilename:   filepath.Base(makeFilePath),
			MakeFile:       makefileBytes,
		},
		JobTimeout: timeout,
	}

	jobId, err := JobTestService.NewJob(newJob)
	if err != nil {
		log.Fatal().Err(err).Msgf("Error creating job")
	}

	return jobId
}

func testJob(t *testing.T, jobId string, correctOutput string) {
	jobInfo, err := JobTestService.WaitForJob(context.Background(), jobId)
	if err != nil {
		t.Fatalf("Error waiting for job: %v", err)
		return
	}

	returned := strings.TrimSpace(jobInfo.StatusMessage)
	expected := strings.TrimSpace(correctOutput)

	if returned != expected {
		t.Fatal("Expected correct output, got: '", correctOutput, "' instead got: ", jobInfo.StatusMessage)
	} else {
		// delete output file if correct
		err := os.Remove(jobInfo.OutputFilePath)
		if err != nil {
			log.Warn().Msgf("Error while removing file: %v", err)
			return
		}
	}
}
