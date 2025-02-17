package common

import (
	"errors"
	"fmt"
	"github.com/makeopensource/leviathan/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	defer func() {
		log.Logger = FileConsoleLogger()
	}()

	baseDir := getBaseDir()
	configDir := getConfigDir(baseDir)

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configDir)

	setupDefaultOptions(configDir)

	submissionFolderPath := getStringEnvOrDefault("SUBMISSION_FOLDER", fmt.Sprintf("%s/%s", baseDir, "submissions"))
	viper.SetDefault(submissionFolderKey, submissionFolderPath)

	outputFolderPath := getStringEnvOrDefault("OUTPUT_FOLDER", fmt.Sprintf("%s/%s", baseDir, "output"))
	viper.SetDefault(outputFolderKey, submissionFolderPath)

	dockerFolderPath := getStringEnvOrDefault("DOCKERFILE_FOLDER", fmt.Sprintf("%s/%s", baseDir, "dockerfiles"))
	viper.SetDefault(dockerFilesFolderKey, dockerFolderPath)

	err := makeDirectories([]string{submissionFolderPath, dockerFolderPath, outputFolderPath})

	if err := viper.SafeWriteConfig(); err != nil {
		var configFileAlreadyExistsError viper.ConfigFileAlreadyExistsError
		if !errors.As(err, &configFileAlreadyExistsError) {
			log.Fatal().Err(err).Msg("viper could not write to config file")
		}
	}

	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("could not read config file")
	}

	log.Info().Msg("Watching config file")
	viper.WatchConfig()

	// maybe create viper instance and return from this function
	// future setup in case https://github.com/spf13/viper/issues/1855 is accepted
}

func getConfigDir(baseDir string) string {
	configDir := fmt.Sprintf("%s/%s", baseDir, "config")
	err := os.MkdirAll(configDir, DefaultFilePerm)
	if err != nil {
		log.Fatal().Err(err).Str("Config dir", configDir).Msgf("could not create config directory")
	}
	return configDir
}

func getStringEnvOrDefault(key, defaultVal string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal
	}
	return value
}

func setupDefaultOptions(configDir string) {
	// misc application files
	viper.SetDefault(dbPathKey, fmt.Sprintf("%s/leviathan.db", configDir))
	viper.SetDefault(logDirKey, fmt.Sprintf("%s/logs/leviathan.log", configDir))
	viper.SetDefault(serverPortKey, "11200")
	viper.SetDefault(enableLocalDockerKey, true)
	viper.SetDefault(concurrentJobsKey, 50)
}

func getBaseDir() string {
	baseDir := "./appdata"
	if os.Getenv("IS_DOCKER") != "" {
		baseDir = "/appdata"
	}
	return baseDir
}

func makeDirectories(dirs []string) error {
	for _, dir := range dirs {
		err := os.MkdirAll(dir, DefaultFilePerm)
		if err != nil {
			return fmt.Errorf("unable to create directory at %s: %v", dir, err)
		}
	}
	return nil
}

func GetClientList() []models.MachineOptions {
	var allMachines []models.MachineOptions

	// Get all settings
	allSettings := viper.AllSettings()

	// Navigate to clients.ssh
	clients, ok := allSettings["clients"].(map[string]interface{})
	if !ok {
		log.Error().Msgf("clients section not found or not configured")
		return nil
	}
	ssh, ok := clients["ssh"].(map[string]interface{})
	if !ok {
		log.Error().Msgf("ssh section not found or not configured")
		return nil
	}

	// Iterate over all keys in clients.ssh
	for clientName, clientConfig := range ssh {
		clientMap, ok := clientConfig.(map[string]interface{})
		if !ok {
			fmt.Printf("  Invalid configuration for %s\n", clientName)
			continue
		}

		log.Info().Msgf("Found machine: %s", clientName)

		options := models.MachineOptions{
			Name:           clientName,
			Host:           clientMap["host"].(string),
			Port:           clientMap["port"].(int64),
			User:           clientMap["user"].(string),
			PrivateKeyFile: clientMap["private_key_file"].(string),
		}
		log.Debug().Any("Machine options", options).Msgf("Loaded: %s", clientName)
		allMachines = append(allMachines, options)
	}

	return allMachines
}
