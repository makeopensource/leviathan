package labs

import (
	"github.com/makeopensource/leviathan/common"
	v1 "github.com/makeopensource/leviathan/generated/types/v1"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/service/jobs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"
)

var (
	dkTestService  *docker.DkService
	jobTestService *jobs.JobService
	labTestService *LabService
	setupOnce      sync.Once
)

const (
	imageName      = "arithmetic-python"
	dockerFilePath = "../../../example/simple-addition/ex-Dockerfile"
	makeFilePath   = "../../../example/simple-addition/makefile"
	graderFilePath = "../../../example/simple-addition/grader.py"
)

func TestLabService_CreateLab(t *testing.T) {
	initDeps()

	graderBytes, err := os.ReadFile(graderFilePath)
	if err != nil {
		t.Fatalf("Error reading grader.py: %v", err)
		return
	}
	makefileBytes, err := os.ReadFile(makeFilePath)
	if err != nil {
		t.Fatalf("Error reading grader.py: %v", err)
		return
	}
	dockerBytes, err := os.ReadFile(dockerFilePath)
	if err != nil {
		t.Fatalf("Error reading docker file: %v", err)
		return
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
		ImageTag:    "test-lab:v1",
		JobEntryCmd: "make grade",
	}

	createLab, err := labTestService.CreateLab(&lab, dockerFile, jobFiles)
	if err != nil {
		t.Fatalf("Error creating lab: %v", err)
		return
	}

	t.Logf("Created Lab: %v", createLab)
}

func initDeps() {
	setupOnce.Do(func() {
		common.InitConfig()
		db, bc := common.InitDB()

		dkTestService = docker.NewDockerServiceWithClients()
		labTestService = NewLabService(db, dkTestService)
		jobTestService = jobs.NewJobService(db, bc, dkTestService, labTestService)

		// no logs on tests
		log.Logger = log.Logger.Level(zerolog.Disabled)
	})
}
