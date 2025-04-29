package database

import (
	"fmt"
	"github.com/makeopensource/leviathan/internal/config"
	"github.com/makeopensource/leviathan/internal/jobs"
	"github.com/makeopensource/leviathan/internal/labs"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"
	"time"
)

func InitDB() (*gorm.DB, *jobs.BroadcastChannel) {
	var connection gorm.Dialector
	var dbConfig *gorm.Config

	if config.EnablePostgres.GetBool() {
		connection, dbConfig = usePostgres()
	} else {
		connection, dbConfig = useSqlite()
	}

	db, err := gorm.Open(connection, dbConfig)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	if config.EnablePostgres.GetBool() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to database")
		}
		sqlDB.SetMaxIdleConns(10)           // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxOpenConns(100)          // SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	}

	err = db.AutoMigrate(&labs.Lab{}, &jobs.Job{})
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to migrate database")
	}

	bc, ctx := jobs.NewBroadcastChannel()
	db = db.WithContext(ctx) // inject broadcast channel to database

	return db, bc
}

func useSqlite() (gorm.Dialector, *gorm.Config) {
	dbPath := config.SqliteDbPath.GetStr()
	if dbPath == "" {
		log.Fatal().Msgf("db_path is empty")
	}

	connectionStr := sqlite.Open(dbPath + "?_journal_mode=WAL&_busy_timeout=5000")
	dbConfig := &gorm.Config{
		PrepareStmt: true,
	}

	abs, err := filepath.Abs(dbPath)
	if err != nil {
		log.Warn().Err(err).Msgf("failed to determine absolute path")
		log.Info().Msgf("using sqlite at: %s", dbPath)
	} else {
		log.Info().Msgf("using sqlite at: %s", abs)
	}

	return connectionStr, dbConfig
}

func usePostgres() (gorm.Dialector, *gorm.Config) {
	host := config.PostgresHost.GetStr()
	user := config.PostgresUser.GetStr()
	password := config.PostgresPass.GetStr()
	database := config.PostgresDB.GetStr()
	port := config.PostgresPort.GetStr()
	sslMode := config.PostgresSsl.GetStr()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslMode=%s",
		host,
		user,
		password,
		database,
		port,
		sslMode,
	)

	log.Info().Msgf("using postgres at: %s", dsn)

	return postgres.Open(dsn), &gorm.Config{}
}
