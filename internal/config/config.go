package config

import (
	"errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AutomaticEnv()

	isDocker := os.Getenv("IS_DOCKER")

	configDir := ""
	if isDocker == "" {
		log.Info().Msgf("Leviathan is running in dev machine")
		configDir = "." // look for config in the working directory, when developing
	} else {
		log.Info().Msgf("Leviathan is running in docker")
		configDir = "/app/" // look in /app when running in docker
	}
	viper.AddConfigPath(configDir)

	if err := viper.ReadInConfig(); err != nil {
		var configFileNotFoundError viper.ConfigFileNotFoundError
		if errors.As(err, &configFileNotFoundError) {
			log.Warn().Msgf("Config file not found, creating default config in %s", configDir)
			createDefaultConfigFile(configDir)
			err := viper.ReadInConfig()
			if err != nil {
				log.Fatal().Err(err).Msgf("Config file not found, even after creating a default file in %s", configDir)
				return
			}
		} else {
			log.Fatal().Err(err).Msgf("Something went wrong while reading config file")
		}
	} else {
		log.Info().Msgf("Successfully loaded config file from %s", viper.ConfigFileUsed())
	}
}

func createDefaultConfigFile(configPath string) {
	// todo figure out config options
	filename := "config.toml"
	finalPath := filepath.Join(configPath, filename)
	_, err := os.Create(finalPath)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to create default config file at %s", finalPath)
	}

	log.Info().Msgf("Created default config file at %s", finalPath)
}
