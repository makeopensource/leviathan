package service

import (
	"github.com/makeopensource/leviathan/common"
	"github.com/makeopensource/leviathan/models"
	"github.com/makeopensource/leviathan/service/docker"
	"github.com/makeopensource/leviathan/service/jobs"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	"os"
	"path/filepath"
)

func InitServices() (*docker.DkService, *jobs.JobService) {
	// common for services
	db, bc := common.InitDB()
	clientList := docker.InitDockerClients()

	dkService := docker.NewDockerService(clientList)
	jobService := jobs.NewJobService(db, bc, dkService)

	cleanupOrphanJobs(db, dkService)

	return dkService, jobService
}

// removes any job left in an 'active' state before application start,
// fail any jobs that were running before leviathan was able to process them (for whatever reason)
//
// for example machine running leviathan shutdown unexpectedly or leviathan had an unrecoverable error
func cleanupOrphanJobs(db *gorm.DB, dk *docker.DkService) {
	var orphanJobs []*models.Job
	res := db.
		Where("status = ?", string(models.Queued)).
		Or("status = ?", string(models.Running)).
		Or("status = ?", string(models.Preparing)).
		Find(&orphanJobs)
	if res.Error != nil {
		log.Warn().Err(res.Error).Msgf("Failed to query database for orphan jobs")
		return
	}

	for _, orphan := range orphanJobs {
		client, err := dk.ClientManager.GetClientById(orphan.MachineId)
		if err != nil {
			log.Warn().Err(err).
				Msgf("unable to find machine: %s ,job: %s was running on", orphan.MachineId, orphan.JobId)
			continue
		}

		err = client.RemoveContainer(orphan.ContainerId, true, true)
		if err != nil {
			log.Warn().Err(err).Str("containerID", orphan.ContainerId).Msg("unable to remove orphan container")
		}

		tmpFold := filepath.Dir(orphan.TmpJobFolderPath) // get the dir above autolab subdir
		err = os.RemoveAll(tmpFold)
		if err != nil {
			log.Warn().Err(err).Str("dir", tmpFold).Msg("unable to remove orphan tmp job directory")
			return
		}

		orphan.Status = models.Failed
		orphan.StatusMessage = "job was unable to be processed due to an internal server error"
		res = db.Save(orphan)
		if res.Error == nil {
			log.Warn().Err(res.Error).Msg("unable to update orphan job status")
		}
	}

	if len(orphanJobs) != 0 {
		log.Info().Msgf("Cleaned up %d orphan jobs", len(orphanJobs))
	}
}
