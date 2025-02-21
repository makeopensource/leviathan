package jobs

import (
	"context"
	"fmt"
	"github.com/makeopensource/leviathan/common"
	v1 "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"path/filepath"
	"strings"
	"sync"
	"testing"
	"time"
)

var (
	dkTestService  *docker.DkService
	jobTestService *JobService
	setupOnce      sync.Once
)

const (
	imageName      = "arithmetic-python"
	dockerFilePath = "../../../example/simple-addition/ex-Dockerfile"
	makeFilePath   = "../../../example/simple-addition/makefile"
	graderFilePath = "../../../example/simple-addition/grader.py"
)

type testCase struct {
	studentFile    string
	expectedOutput string
	correctStatus  models.JobStatus
}

var (
	defaultTimeout = time.Second * 10
	testCases      = map[string]testCase{
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
		"timeout_edge": {
			studentFile:    "../../../example/simple-addition/student_timeout_edge.py",
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
	setupTest()
	correct := testCases["correct"]
	testJobProcessor(t, correct.studentFile, correct.expectedOutput, defaultTimeout, correct.correctStatus)
}

func TestIncorrect(t *testing.T) {
	setupTest()
	incorrect := testCases["incorrect"]
	testJobProcessor(t, incorrect.studentFile, incorrect.expectedOutput, defaultTimeout, incorrect.correctStatus)
}

// TODO
func TestForkBomb(t *testing.T) {
	setupTest()
	forkBomb := testCases["forkb"]
	testJobProcessor(t, forkBomb.studentFile, forkBomb.expectedOutput, defaultTimeout, forkBomb.correctStatus)
}

func TestTimeout(t *testing.T) {
	setupTest()
	timeLimit := time.Second * 10
	timeout := testCases["timeout"]
	timeout.expectedOutput = fmt.Sprintf("Maximum timeout reached for job, job ran for %s", timeLimit)
	testJobProcessor(t, timeout.studentFile, timeout.expectedOutput, timeLimit, timeout.correctStatus)
}

// TestTimeoutEdge takes in a submission that takes 11 seconds to run,
// designed to check if the go scheduler, is correctly selecting the timeout case
// intended to be run in a batch test
func TestTimeoutEdge(t *testing.T) {
	setupTest()
	timeLimit := time.Second * 10
	timeout := testCases["timeout_edge"]
	timeout.expectedOutput = fmt.Sprintf("Maximum timeout reached for job, job ran for %s", timeLimit)
	testJobProcessor(t, timeout.studentFile, timeout.expectedOutput, timeLimit, timeout.correctStatus)
}

func TestCancel(t *testing.T) {
	setupTest()
	timeLimit := time.Second * 10
	timeout := testCases["timeout"]
	timeout.expectedOutput = fmt.Sprintf("Job was cancelled")

	jobId := setupJobProcess(timeout.studentFile, timeLimit)

	// cancel the job after 3 seconds
	time.AfterFunc(3*time.Second, func() {
		jobTestService.CancelJob(jobId)
	})

	testJob(t, jobId, timeout.expectedOutput, models.Canceled)

	// verify cancel function was removed from context map
	value := jobTestService.queue.getJobCancelFunc(jobId)
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

	dockerBytes, err := common.ReadFileBytes(dockerFilePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading docker file")
	}

	newJob := &models.Job{
		LabData:     models.LabModel{ImageTag: imageName},
		JobTimeout:  timeout,
		JobEntryCmd: "make grade",
	}

	jobId, err := jobTestService.NewJob(
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
			Filename: filepath.Base(dockerFilePath),
			Content:  dockerBytes,
		},
	)

	if err != nil {
		log.Fatal().Err(err).Msgf("Error creating job")
	}

	return jobId
}

func testJob(t *testing.T, jobId string, correctOutput string, correctStatus models.JobStatus) {
	jobInfo, logs, err := jobTestService.WaitForJobAndLogs(context.Background(), jobId)
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

func setupTest() {
	setupOnce.Do(func() {
		common.InitConfig()
		initServices()
	})
}

func initServices() {
	// common for services
	db := common.InitDB()
	bc, ctx := models.NewBroadcastChannel()
	// inject broadcast channel to database
	db = db.WithContext(ctx)

	fCache := models.NewLabFilesCache(db)
	clientList := docker.InitDockerClients()

	dkTestService = docker.NewDockerService(clientList)
	jobTestService = NewJobService(db, fCache, bc, dkTestService) // depends on docker service
}
