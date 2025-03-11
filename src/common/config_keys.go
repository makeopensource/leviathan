package common

import (
	"github.com/spf13/viper"
	"time"
)

const (
	concurrentJobsKey = "jobs.concurrent_jobs"

	apiKeyKey     = "server.apikey"
	serverPortKey = "server.port"
	loglevelKey   = "server.log_level"

	// folders
	logDirKey        = "folder.log_dir"
	submissionDirKey = "folder.submission_dir"
	outputDirKey     = "folder.log_output_dir"
	// docker config
	enableLocalDockerKey = "clients.enable_local_docker"
	ClientSSHKey         = "clients.ssh"

	sqliteDbPathKey = "db.sqlite.db_path"
	// postgres
	enablePostgresKey = "db.postgres.enable_postgres"
	postgresHostKey   = "db.postgres.postgres_host"
	postgresPortKey   = "db.postgres.postgres_port"
	postgresUserKey   = "db.postgres.postgres_user"
	postgresPassKey   = "db.postgres.postgres_pass"
	postgresDBKey     = "db.postgres.postgres_db"
	postgresSslKey    = "db.postgres.postgres_ssl"
)

var (
	// internal use
	LogLevel     = Config{loglevelKey}
	LogDir       = Config{logDirKey}
	SqliteDbPath = Config{sqliteDbPathKey}

	// general

	ApiKey         = Config{apiKeyKey}
	ServerPort     = Config{serverPortKey}
	ConcurrentJobs = Config{concurrentJobsKey}

	// folderstuff

	SubmissionFolder  = Config{submissionDirKey}
	OutputFolder      = Config{outputDirKey}
	EnableLocalDocker = Config{enableLocalDockerKey}

	// postgres

	EnablePostgres = Config{enablePostgresKey}
	postgresHost   = Config{postgresHostKey}
	postgresPort   = Config{postgresPortKey}
	postgresUser   = Config{postgresUserKey}
	postgresPass   = Config{postgresPassKey}
	postgresDB     = Config{postgresDBKey}
	postgresSsl    = Config{postgresSslKey}
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
