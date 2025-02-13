package tests

import (
	"fmt"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/utils"
	"github.com/rs/zerolog/log"
	"golang.org/x/exp/maps"
	"math/rand"
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
	}
)

func TestCorrect(t *testing.T) {
	setupTest()
	correct := testCases["correct"]
	testJobProcessor(t, correct.studentFile, correct.expectedOutput, defaultTimeout)
}

func TestIncorrect(t *testing.T) {
	setupTest()
	incorrect := testCases["incorrect"]
	testJobProcessor(t, incorrect.studentFile, incorrect.expectedOutput, defaultTimeout)
}

func TestTimeout(t *testing.T) {
	setupTest()
	timeLimit := time.Second * 10
	timeout := testCases["timeout"]
	timeout.expectedOutput = fmt.Sprintf("Maximum timeout reached for job, job ran for %s", timeLimit)
	testJobProcessor(t, timeout.studentFile, timeout.expectedOutput, timeLimit)
}

func Test50Jobs(t *testing.T) {
	testBatchJobProcessor(t, 50)
}

func Test100Jobs(t *testing.T) {
	testBatchJobProcessor(t, 100)
}

func Test500Jobs(t *testing.T) {
	testBatchJobProcessor(t, 500)
}

func testBatchJobProcessor(t *testing.T, numJobs int) {
	setupTest()

	testValues := maps.Values(testCases)

	for i := 0; i < numJobs; i++ {
		// Randomly choose from test cases
		testCaseIndex := rand.Intn(len(testValues))
		testCase := testValues[testCaseIndex]

		// Run the test for each job directly in the main test goroutine
		t.Run(fmt.Sprintf("Job_%d", i), func(t *testing.T) {
			// Create subtests for better reporting
			// Enable parallel execution for this subtest
			t.Parallel()
			testJobProcessor(t, testCase.studentFile, testCase.expectedOutput, defaultTimeout)
			fmt.Printf("Job %d finished\n", i)
		})
	}
}

func testJobProcessor(t *testing.T, studentCodePath string, correctOutput string, timeout time.Duration) {
	graderBytes, err := utils.ReadFileBytes(graderFilePath)
	if err != nil {
		t.Fatalf("Error reading grader.py: %v", err)
	}

	makefileBytes, err := utils.ReadFileBytes(makeFilePath)
	if err != nil {
		t.Fatalf("Error reading makefile: %v", err)
	}

	studentBytes, err := utils.ReadFileBytes(studentCodePath)
	if err != nil {
		t.Fatalf("Error reading student: %v", err)
	}

	newJob := &models.Job{
		ImageTag:                  imageName,
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

	jobId, err := jobService.NewJob(newJob)
	if err != nil {
		t.Fatal("Error creating job: ", err)
	}

	jobInfo, err := jobService.WaitForJob(jobId)
	if err != nil {
		t.Fatalf("Error waiting for job: %v", err)
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
