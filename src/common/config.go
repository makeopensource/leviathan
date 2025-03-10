package common

import (
	"errors"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg(".env not found. you can safely ignore this warning if you dont have a .env file")
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
	viper.SetConfigType("json")
	viper.AddConfigPath(configDir)

	setupDefaultOptions(configDir)
	loadPostgresOptions()

	submissionFolderPath := getStringEnvOrDefault("TMP_SUBMISSION_FOLDER", fmt.Sprintf("%s/%s", baseDir, "submissions"))
	viper.SetDefault(submissionDirKey, submissionFolderPath)

	outputFolderPath := getStringEnvOrDefault("LOG_OUTPUT_FOLDER", fmt.Sprintf("%s/%s", baseDir, "output"))
	viper.SetDefault(outputDirKey, outputFolderPath)

	err = makeDirectories([]string{submissionFolderPath, outputFolderPath})
	if err != nil {
		log.Fatal().Err(err).Msg("unable to make required directories")
	}

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

	log.Info().Msgf("watching config file at %s", viper.ConfigFileUsed())
	viper.WatchConfig()

	// maybe create viper instance and return from this function
	// future setup in case https://github.com/spf13/viper/issues/1855 is accepted
}

func loadPostgresOptions() {
	enablePost := false
	if getStringEnvOrDefault("POSTGRES_ENABLE", "false") == "true" {
		enablePost = true
	}
	viper.SetDefault(
		enablePostgresKey,
		enablePost,
	)
	viper.SetDefault(
		postgresHostKey,
		getStringEnvOrDefault("POSTGRES_HOST", "localhost"),
	)
	viper.SetDefault(
		postgresPortKey,
		getStringEnvOrDefault("POSTGRES_PORT", "5432"),
	)
	viper.SetDefault(
		postgresUserKey,
		getStringEnvOrDefault("POSTGRES_USER", "postgres"),
	)
	viper.SetDefault(
		postgresPassKey,
		getStringEnvOrDefault("POSTGRES_PASSWORD", ""),
	)
	viper.SetDefault(
		postgresDBKey,
		getStringEnvOrDefault("POSTGRES_DB", "postgres"),
	)
	viper.SetDefault(
		postgresSslKey,
		getStringEnvOrDefault("POSTGRES_SSL", "disable"),
	)
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
	viper.SetDefault(sqliteDbPathKey, fmt.Sprintf("%s/leviathan.db", configDir))
	viper.SetDefault(logDirKey, fmt.Sprintf("%s/logs/leviathan.log", configDir))
	viper.SetDefault(serverPortKey, "9221")
	viper.SetDefault(enableLocalDockerKey, true)
	viper.SetDefault(concurrentJobsKey, 50)
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
