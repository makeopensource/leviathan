package labs

import (
	"github.com/makeopensource/leviathan/internal/config"
	"github.com/makeopensource/leviathan/internal/database"
	"github.com/makeopensource/leviathan/internal/docker"
	fm "github.com/makeopensource/leviathan/internal/file_manager"
	"github.com/makeopensource/leviathan/pkg/file_utils"
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
	fileMan        *fm.FileManagerService
	labTestService *LabService
	setupOnce      sync.Once
)

const (
	dockerFilePath = "../../../example/simple-addition/ex-Dockerfile"
	makeFilePath   = "../../../example/simple-addition/makefile"
	graderFilePath = "../../../example/simple-addition/grader.py"
)

func TestLabService_CreateLab(t *testing.T) {
	initDeps()

	graderBytes, err := os.Open(graderFilePath)
	if err != nil {
		t.Fatalf("Error reading grader.py: %v", err)
		return
	}
	makefileBytes, err := os.Open(makeFilePath)
	if err != nil {
		t.Fatalf("Error reading makefile: %v", err)
		return
	}
	dockerfile, err := os.Open(dockerFilePath)
	if err != nil {
		t.Fatalf("Error reading docker file: %v", err)
		return
	}

	files := []*fm.FileInfo{
		{
			Reader:   makefileBytes,
			Filename: filepath.Base(makeFilePath),
		},
		{
			Reader:   graderBytes,
			Filename: filepath.Base(graderFilePath),
		},
	}

	tmpFolderID, err := fileMan.CreateTmpLabFolder(dockerfile, files...)
	if err != nil {
		t.Fatalf("Error creating tmp folder: %v", err)
		return
	}

	lab := Lab{
		Name:        "test-lab",
		JobTimeout:  time.Second * 10,
		ImageTag:    "test-lab:v1",
		JobEntryCmd: "make grade",
	}

	createLab, err := labTestService.CreateLab(&lab, tmpFolderID)
	if err != nil {
		t.Fatalf("Error creating lab: %v", err)
		return
	}

	t.Logf("Created Lab: %v", createLab)

	labDta, err := labTestService.GetLabFromDB(createLab)
	if err != nil {
		t.Fatalf("Error retrieving lab: %v", err)
		return
	}

	if !file_utils.FileExists(labDta.JobFilesDirPath) {
		t.Fatalf("Job files dir does not exist")
		return
	}
	if !file_utils.FileExists(labDta.DockerFilePath) {
		t.Fatalf("Dockerfile does not exist")
		return
	}
}

func initDeps() {
	setupOnce.Do(func() {
		config.InitConfig()
		db, _ := database.InitDB()

		dkTestService = docker.NewDockerServiceWithClients()
		fileMan = fm.NewFileManagerService()
		labTestService = NewLabService(db, dkTestService, fileMan)

		// no logs on tests
		log.Logger = log.Logger.Level(zerolog.Disabled)
	})
}
