package jobs

import (
	"fmt"
	"github.com/makeopensource/leviathan/common"
	v1 "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/service/labs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/stretchr/testify/assert"
	"os"
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
		"oom": {
			studentFile:    "../../../example/simple-addition/student_oom.py",
			expectedOutput: "unable to parse log output",
			correctStatus:  models.Failed,
		},
		"forkb": {
			studentFile:    "../../../example/simple-addition/student_fork_bomb.py",
			expectedOutput: `{"addition": {"passed": false, "message": "Addition test caused an error: [Errno 11] Resource temporarily unavailable"}, "subtraction": {"passed": true, "message": ""}, "multiplication": {"passed": false, "message": "Multiplication failed. Expected 42, got 48"}, "division": {"passed": false, "message": "Division failed. Expected 4, got 3.3333333333333335"}}`,
			correctStatus:  models.Complete, // job completes since we can parse the last line
		},
	}
	testFuncs = map[string]func(*testing.T){
		"correct":      TestCorrect,
		"incorrect":    TestIncorrect,
		"cancel":       TestCancel,
		"timeout":      TestTimeout,
		"timeout_edge": TestTimeoutEdge,
		"oom":          TestOom,
		"forkb":        TestForkBomb,
	}
)

func TestAll(t *testing.T) {
	for tCase, test := range testFuncs {
		t.Run(tCase, func(t *testing.T) {
			t.Parallel()
			test(t)
		})
	}
}

func TestCorrect(t *testing.T) {
	setupTest()
	correct := testCases["correct"]
	testJobProcessor(t, correct.studentFile, correct.expectedOutput, defaultTimeout, correct.correctStatus)
}

func TestOom(t *testing.T) {
	setupTest()
	correct := testCases["oom"]
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
// designed to check if the go scheduler is correctly selecting the timeout case
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
	timeout.expectedOutput = "Job was cancelled"

	jobId := setupJobProcess(timeout.studentFile, timeLimit)

	// cancel the job after 3 seconds
	time.AfterFunc(3*time.Second, func() {
		jobTestService.CancelJob(jobId)
	})

	testJob(t, jobId, timeout.expectedOutput, models.Canceled)

	// verify cancel function was removed from context map
	_, ok := jobTestService.queue.contextMap.Load(jobId)
	if ok {
		t.Fatalf("Job was cancelled, but the cancel func was not nil")
	}
}

func testJobProcessor(t *testing.T, studentCodePath string, correctOutput string, timeout time.Duration, status models.JobStatus) {
	jobId := setupJobProcess(studentCodePath, timeout)
	testJob(t, jobId, correctOutput, status)
}

func setupJobProcess(studentCodePath string, timeout time.Duration) string {
	graderBytes, err := os.ReadFile(graderFilePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading grader.py")
	}

	makefileBytes, err := os.ReadFile(makeFilePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading grader.py")
	}

	studentBytes, err := os.ReadFile(studentCodePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading student")
	}

	dockerBytes, err := os.ReadFile(dockerFilePath)
	if err != nil {
		log.Fatal().Err(err).Msg("Error reading docker file")
	}

	newLab := models.Lab{
		JobTimeout:  timeout,
		JobEntryCmd: "make grade",
		ImageTag:    imageName,
	}

	newJob := &models.Job{LabData: &newLab}

	jobId, err := jobTestService.NewJobFromRPC(
		newJob,
		[]*v1.FileUpload{
			{
				Filename: filepath.Base(makeFilePath),
				Content:  makefileBytes,
			},
			{
				Filename: filepath.Base(graderFilePath),
				Content:  graderBytes,
			},
			{
				Filename: "student.py",
				Content:  studentBytes,
			},
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
	jobInfo, logs, err := jobTestService.WaitForJobAndLogs(jobId)
	if err != nil {
		t.Fatalf("Error waiting for job: %v", err)
		return
	}

	t.Log("Job ID: ", jobId, " Logs: %s", logs)

	returned := strings.TrimSpace(jobInfo.StatusMessage)
	expected := strings.TrimSpace(correctOutput)

	assert.Equal(t, expected, returned)
	assert.Equal(t, correctStatus, jobInfo.Status)
}

func setupTest() {
	setupOnce.Do(func() {
		initServices()
	})
}

func initServices() {
	common.InitConfig()

	db, bc := common.InitDB()

	dkTestService = docker.NewDockerServiceWithClients()
	lab := labs.NewLabService(db, dkTestService)
	jobTestService = NewJobService(db, bc, dkTestService, lab)

	// no logs on tests
	log.Logger = log.Logger.Level(zerolog.Disabled)
}
