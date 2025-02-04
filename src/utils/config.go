package utils

import (
	"errors"
	"fmt"
	"github.com/makeopensource/leviathan/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	baseDir := getBaseDir()
	configDir := getConfigDir(baseDir)

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configDir)

	setupDefaultOptions(configDir)

	submissionFolder := getStringEnvOrDefault("SUBMISSION_FOLDER", fmt.Sprintf("%s/%s", baseDir, "submissions"))
	viper.SetDefault("submission_folder", submissionFolder)

	err := makeDirectories([]string{submissionFolder})

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
	// set database path
	viper.SetDefault("db_path", fmt.Sprintf("%s/leviathan.db", configDir))
	// create log directory
	viper.SetDefault("log_file", fmt.Sprintf("%s/logs/leviathan.log", configDir))
	// Set general settings
	viper.SetDefault("server.port", "11200")
	// enable local docker
	viper.SetDefault("clients.enable_local_docker", true)
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
