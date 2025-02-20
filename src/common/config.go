package common

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/makeopensource/leviathan/models"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"os"
)

//dbPathKey         = "db_path"
//logDirKey = "log_dir"
//apiKeyKey = "apikey"
//serverPortKey = "server.port"
//concurrentJobsKey = "concurrent_jobs"
//
//// folders
//submissionFolderKey  = "folder.submission_folder"
//dockerFilesFolderKey = "folder.docker_files_folder"
//outputFolderKey = "folder.output_folder"
//// docker config
//enableLocalDockerKey = "clients.enable_local_docker"
//type Prefs struct {
//	Server struct {
//		Host string `mapstructure:"host" default:"localhost"`
//		Port int    `mapstructure:"port" default:"8080"`
//		TLS  bool   `mapstructure:"tls"` // No default, will be false
//	} `mapstructure:"server"`
//	Machines []models.MachineOptions `mapstructure:"database"`
//	Database struct {
//		User     string `mapstructure:"user" default:"admin"`
//		Password string `mapstructure:"password"` // No default
//		Name     string `mapstructure:"name" default:"mydatabase"`
//	} `mapstructure:"database"`
//	EnablelocalDocker bool   `mapstructure:"enable_local_docker"`
//	LogLevel          string `mapstructure:"log_level" default:"info"`
//	// Example with nested defaults
//	Timeout time.Duration `mapstructure:"timeout" default:"5s"` // Example with time.Duration
//}

func InitConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Warn().Err(err).Msg("enable to load .env file")
	}

	defer func() {
		log.Logger = FileConsoleLogger()
	}()

	baseDir := getBaseDir()
	configDir := getConfigDir(baseDir)

	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath(configDir)

	setupDefaultOptions(configDir)
	loadPostgresOptions()

	submissionFolderPath := getStringEnvOrDefault("SUBMISSION_FOLDER", fmt.Sprintf("%s/%s", baseDir, "submissions"))
	viper.SetDefault(submissionFolderKey, submissionFolderPath)

	outputFolderPath := getStringEnvOrDefault("OUTPUT_FOLDER", fmt.Sprintf("%s/%s", baseDir, "output"))
	viper.SetDefault(outputFolderKey, submissionFolderPath)

	err = makeDirectories([]string{submissionFolderPath, outputFolderPath})

	if err := viper.WriteConfig(); err != nil {
		log.Fatal().Err(err).Msg("viper could not write to config file")
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
