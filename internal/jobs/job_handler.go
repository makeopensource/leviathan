package jobs

// TODO
func HandleNewJob(imageTag string, courseName string, studentFileTarFile []byte) error {
	//jobId, err := uuid.NewUUID()
	//if err != nil {
	//	log.Error().Err(err).Msgf("Failed to generaete job ID")
	//	return errors.New("failed to generate job ID")
	//}
	//
	//// create directory to for outputfile
	//folder := fmt.Sprintf("%d/%d")

	return nil
}

func HandleJobStatus(jobUuid string) error {
	return nil
}

func HandleCancelJob(jobUuid string) error {
	return nil
}
