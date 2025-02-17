package common

import (
	"github.com/spf13/viper"
	"time"
)

const (
	dbPathKey         = "db_path"
	logDirKey         = "log_dir"
	apiKeyKey         = "apikey"
	serverPortKey     = "server.port"
	concurrentJobsKey = "concurrent_jobs"

	// folders
	submissionFolderKey  = "folder.submission_folder"
	dockerFilesFolderKey = "folder.docker_files_folder"
	outputFolderKey      = "folder.output_folder"
	// docker config
	enableLocalDockerKey = "clients.enable_local_docker"
)

var (
	// internal use
	LogDir = Config{logDirKey}
	DbPath = Config{dbPathKey}
	// general
	ApiKey         = Config{apiKeyKey}
	ServerPort     = Config{serverPortKey}
	ConcurrentJobs = Config{concurrentJobsKey}
	// folderstuff
	SubmissionTarFolder = Config{submissionFolderKey}
	DockerFilesFolder   = Config{dockerFilesFolderKey}
	OutputFolder        = Config{outputFolderKey}
	EnableLocalDocker   = Config{enableLocalDockerKey}
)

type Config struct {
	ConfigKey string
}

func (c Config) Set(value any) {
	viper.Set(c.ConfigKey, value)
}

func (c Config) GetStr() string {
	return viper.GetString(c.ConfigKey)
}

func (c Config) GetInt() int {
	return viper.GetInt(c.ConfigKey)
}

func (c Config) GetDuration() time.Duration {
	return viper.GetDuration(c.ConfigKey)
}

func (c Config) GetBool() bool {
	return viper.GetBool(c.ConfigKey)
}

func (c Config) GetUint64() uint64 {
	return viper.GetUint64(c.ConfigKey)
}
