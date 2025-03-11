package common

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/makeopensource/leviathan/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
	"strconv"
)

func InitConfig() {
	_, ok := os.LookupEnv("LEVIATHAN_IS_DOCKER")
	if !ok {
		err := godotenv.Load() // load .env file for non docker env
		if err != nil {
			log.Warn().Err(err).Msg(".env not found. ignore this warning if you did not intend to load a .env file")
		}
	}

	defer func() {
		log.Logger = FileConsoleLogger()
	}()

	baseDir, err := getBaseDir()
	if err != nil {
		log.Fatal().Err(err).Msg("unable to get base dir")
	}
	configDir := getConfigDir(baseDir)

	viper.SetConfigName("config")
	viper.SetConfigType("toml")
	viper.AddConfigPath(configDir)

	// ignore any error to setup default vals
	_ = viper.ReadInConfig()

	setIfEnvPresentOrDefault(
		loglevelKey,
		"LEVIATHAN_LOG_LEVEL",
		"debug",
	)

	loadPostgresOptions()
	setupDefaultOptions(configDir)

	submissionFolderPath := setIfEnvPresentOrDefault(
		submissionDirKey,
		"TMP_SUBMISSION_DIR",
		fmt.Sprintf("%s/%s", baseDir, "submissions"),
	)
	outputFolderPath := setIfEnvPresentOrDefault(
		outputDirKey,
		"SUBMISSION_OUTPUT_DIR",
		fmt.Sprintf("%s/%s", baseDir, "output"),
	)

	err = makeDirectories([]string{submissionFolderPath, outputFolderPath})
	if err != nil {
		log.Fatal().Err(err).Msg("unable to make required directories")
	}

	if err := viper.SafeWriteConfig(); err != nil {
		var configFileAlreadyExistsError viper.ConfigFileAlreadyExistsError
		if errors.As(err, &configFileAlreadyExistsError) {
			// merge any new changes
			err := viper.WriteConfig()
			if err != nil {
				log.Fatal().Err(err).Msg("viper could not write to config file")
			}
		} else {
			log.Fatal().Err(err).Msg("viper could not write to config file")
		}
	}
	err = viper.ReadInConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("could not read config file")
	}

	log.Info().Msgf("loaded config from %s", viper.ConfigFileUsed())
}

func loadPostgresOptions() {
	setIfEnvPresentOrDefault(postgresHostKey, "POSTGRES_HOST", "localhost")
	setIfEnvPresentOrDefault(postgresPortKey, "POSTGRES_PORT", "5432")
	setIfEnvPresentOrDefault(postgresUserKey, "POSTGRES_USER", "postgres")
	setIfEnvPresentOrDefault(postgresPassKey, "POSTGRES_PASSWORD", "postgres")
	setIfEnvPresentOrDefault(postgresDBKey, "POSTGRES_DB", "postgres")
	setIfEnvPresentOrDefault(postgresSslKey, "POSTGRES_SSL", "disable")

	val, isDefault := getBoolEnvOrDefault("POSTGRES_ENABLE", false)
	if isDefault {
		viper.SetDefault(enablePostgresKey, val)
	} else {
		viper.Set(enablePostgresKey, val)
	}
}

func getConfigDir(baseDir string) string {
	configDir := fmt.Sprintf("%s/%s", baseDir, "config")
	err := os.MkdirAll(configDir, DefaultFilePerm)
	if err != nil {
		log.Fatal().Err(err).Str("Config dir", configDir).Msgf("could not create config directory")
	}
	return configDir
}

// uses viper.Set if env var was found,
//
// else uses' viper.SetDefault and uses defaultValue
//
// this allows us to overwrite any new configration changes passed via env vars,
// but ignore if no env were passed
func setIfEnvPresentOrDefault(configKey, envKeyName, defaultValue string) string {
	val, isDefault := getStringEnvOrDefault(envKeyName, defaultValue)
	if isDefault {
		viper.SetDefault(configKey, val)
	} else {
		// always overwrite with key
		viper.Set(configKey, val)
	}

	return val
}

func getStringEnvOrDefault(key, defaultVal string) (finalVal string, isDefault bool) {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal, true
	}
	return value, false
}

func getBoolEnvOrDefault(key string, defaultVal bool) (finalVal, isDefault bool) {
	value := os.Getenv(key)
	if value == "" {
		return defaultVal, true
	}
	parseBool, err := strconv.ParseBool(value)
	if err != nil {
		return defaultVal, true
	}
	return parseBool, false
}

func setupDefaultOptions(configDir string) {
	// misc application files
	viper.SetDefault(sqliteDbPathKey, fmt.Sprintf("%s/leviathan.db", configDir))
	viper.SetDefault(logDirKey, fmt.Sprintf("%s/logs/leviathan.log", configDir))
	viper.SetDefault(serverPortKey, "9221")
	viper.SetDefault(enableLocalDockerKey, true)
	viper.SetDefault(concurrentJobsKey, 50)
	viper.SetDefault(ClientSSHKey, map[string]models.MachineOptions{
		"example": {
			Enable: false,
			Name:   "example",
			Host:   "http://localhost:8080",
			User:   "test",
			Port:   22,
		},
	})
}

func getBaseDir() (string, error) {
	return filepath.Abs("./appdata")
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
