package config

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
	submissionDirKey = "folder.tmp_submission_dir"
	outputDirKey     = "folder.job_output_dir"
	sshDirKey        = "folder.ssh_config"
	labDirKey        = "folder.labs"
	tmpUploadFolder  = "folder.tmp_uploads"

	// docker config
	enableLocalDockerKey = "clients.enable_local_docker"
	clientSSHKey         = "clients.ssh"

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
	LogLevel = Config{loglevelKey}
	LogDir   = Config{logDirKey}

	// general
	ApiKey         = Config{apiKeyKey}
	ServerPort     = Config{serverPortKey}
	ConcurrentJobs = Config{concurrentJobsKey}

	// folders
	SSHConfigFolder  = Config{sshDirKey}
	SubmissionFolder = Config{submissionDirKey}
	OutputFolder     = Config{outputDirKey}
	LabsFolder       = Config{labDirKey}
	TmpUploadFolder  = Config{tmpUploadFolder}

	// docker config
	EnableLocalDocker = Config{enableLocalDockerKey}
	ClientsSSH        = Config{clientSSHKey}

	// postgres
	EnablePostgres = Config{enablePostgresKey}
	PostgresHost   = Config{postgresHostKey}
	PostgresPort   = Config{postgresPortKey}
	PostgresUser   = Config{postgresUserKey}
	PostgresPass   = Config{postgresPassKey}
	PostgresDB     = Config{postgresDBKey}
	PostgresSsl    = Config{postgresSslKey}

	// sqlite
	SqliteDbPath = Config{sqliteDbPathKey}
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

func (c Config) GetAny() any {
	return viper.Get(c.ConfigKey)
}
