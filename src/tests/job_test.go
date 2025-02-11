package tests

import (
	"fmt"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/service/jobs"
	"github.com/makeopensource/leviathan/utils"
	log2 "log"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

const (
	expectedIncorrectOutput = `{"addition": {"passed": true, "message": ""}, "subtraction": {"passed": true, "message": ""}, "multiplication": {"passed": false, "message": "Multiplication failed. Expected 42, got 48"}, "division": {"passed": false, "message": "Division failed. Expected 4, got 3.3333333333333335"}}`

	expectedCorrectOutput = `{"addition": {"passed": true, "message": ""}, "subtraction": {"passed": true, "message": ""}, "multiplication": {"passed": true, "message": ""}, "division": {"passed": true, "message": ""}}`

	makeFilePath   = "../../example/python/simple-addition/makefile"
	graderFilePath = "../../example/python/simple-addition/grader.py"
	dockerFilePath = "../../example/python/simple-addition/ex-Dockerfile"
	imageName      = "arithmetic-python"
)

var (
	dkService  *docker.DockerService
	jobService *jobs.JobService
)

func initServices() {
	fmt.Println("Starting tests")
	// utils for services
	db := utils.InitDB()
	fCache := utils.NewLabFilesCache(db)
	clientList := docker.InitDockerClients()

	dkService = docker.NewDockerService(clientList)
	jobService = jobs.NewJobService(db, fCache, dkService) // depends on docker service
}

// TestMain is the entry point for running tests in this package.
// It's called *once* before any tests are run.
func TestMain(m *testing.M) {
	utils.InitConfig()
	initServices()
	buildImage()
	code := m.Run() // Run the tests
	os.Exit(code)   // Exit with the appropriate code (0 on success, non-zero on failure)
}

func TestJobProcessorCorrect(t *testing.T) {
	fmt.Println("Starting correct code")
	testJobProcessor(t, "../../example/python/simple-addition/student_correct.py", expectedCorrectOutput)
}

func TestJobProcessorIncorrect(t *testing.T) {
	fmt.Println("Starting incorrect code")
	testJobProcessor(t, "../../example/python/simple-addition/student_incorrect.py", expectedIncorrectOutput)
}

func testJobProcessor(t *testing.T, studentCodePath string, correctOutput string) {
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
	}

	jobId, err := jobService.NewJob(newJob)
	if err != nil {
		t.Fatalf("Error creating job: %v", err)
	}

	jobInfo, err := jobService.WaitForJob(jobId)
	if err != nil {
		t.Fatalf("Error waiting for job: %v", err)
	}

	returned := strings.TrimSpace(jobInfo.StatusMessage)
	expected := strings.TrimSpace(correctOutput)

	if returned != expected {
		t.Error("Expected correct output, got", correctOutput, "instead got", jobInfo.StatusMessage)
	}
}

func buildImage() {
	bytes, err := utils.ReadFileBytes(dockerFilePath)
	if err != nil {
		log2.Fatal("Unable to read Dockerfile " + dockerFilePath)
	}
	err = dkService.NewImageReq(filepath.Base(dockerFilePath), bytes, imageName)
	if err != nil {
		log2.Fatal("Unable to build Dockerfile " + dockerFilePath)
	}
}
