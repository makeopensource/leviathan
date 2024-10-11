package config

import (
	"errors"
	"fmt"
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
		configDir = "./appdata" // look for config in the working directory, when developing
	} else {
		log.Info().Msgf("Leviathan is running in docker")
		configDir = "/app/appdata" // look in /app/appdata when running in docker
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
	filename := "config.toml"
	finalPath := filepath.Join(configPath, filename)
	_, err := os.Create(finalPath)
	if err != nil {
		log.Fatal().Err(err).Msgf("Failed to create default config file at %s", finalPath)
	}
	log.Info().Msgf("Created default config file at %s", finalPath)

	createDefaultOptions()
	err = viper.WriteConfigAs(finalPath)
	if err != nil {
		log.Error().Err(err).Msgf("Failed to write default config file at %s", finalPath)
		return
	}
}

func createDefaultOptions() {
	// todo figure out config options
	viper.SetDefault("clients.enable_local_docker", true)
}

func GetClientList() []MachineOptions {
	var allMachines []MachineOptions

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

		options := MachineOptions{
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
