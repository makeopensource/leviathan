package common

import (
	"fmt"
	"github.com/makeopensource/leviathan/models"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"
	"time"
)

func InitDB() (*gorm.DB, *models.BroadcastChannel) {
	var connection gorm.Dialector
	var config *gorm.Config

	if EnablePostgres.GetBool() {
		connection, config = usePostgres()
	} else {
		connection, config = useSqlite()
	}

	db, err := gorm.Open(connection, config)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
	}

	if EnablePostgres.GetBool() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal().Err(err).Msg("failed to connect to database")
		}
		sqlDB.SetMaxIdleConns(10)           // SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
		sqlDB.SetMaxOpenConns(100)          // SetMaxOpenConns sets the maximum number of open connections to the database.
		sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	}

	err = db.AutoMigrate(&models.Job{})
	if err != nil {
		log.Fatal().Err(err).Msgf("failed to migrate database")
	}

	bc, ctx := models.NewBroadcastChannel()
	db = db.WithContext(ctx) // inject broadcast channel to database

	return db, bc
}

func useSqlite() (gorm.Dialector, *gorm.Config) {
	dbPath := SqliteDbPath.GetStr()
	if dbPath == "" {
		log.Fatal().Msgf("db_path is empty")
	}

	connectionStr := sqlite.Open(dbPath + "?_journal_mode=WAL&_busy_timeout=5000")
	config := &gorm.Config{
		PrepareStmt: true,
	}

	abs, err := filepath.Abs(dbPath)
	if err != nil {
		log.Warn().Err(err).Msgf("failed to determine absolute path")
		log.Info().Msgf("using sqlite at: %s", dbPath)
	} else {
		log.Info().Msgf("using sqlite at: %s", abs)
	}

	return connectionStr, config
}

func usePostgres() (gorm.Dialector, *gorm.Config) {
	host := postgresHost.GetStr()
	user := postgresUser.GetStr()
	password := postgresPass.GetStr()
	database := postgresDB.GetStr()
	port := postgresPort.GetStr()
	sslmode := postgresSsl.GetStr()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		host,
		user,
		password,
		database,
		port,
		sslmode,
	)

	log.Info().Msgf("using postgres at: %s", dsn)

	return postgres.Open(dsn), &gorm.Config{}
}
