package main

import (
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
)

func main() {
	logger := log.With().Str("event", "migration").Logger()

	migrationDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	mysqlDSN := "mysql://root:admin@tcp(127.0.0.1:3306)/hero_dev"
	logger.Debug().Msgf("connecting mysql with dsn %s", mysqlDSN)

	m, err := migrate.New(fmt.Sprintf("file://%s/sql/", migrationDir), mysqlDSN)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to initialize migration")
		panic(err)
	}

	if err = m.Up(); err != nil {
		switch err {
		case migrate.ErrNoChange:
			logger.Error().Err(err).Msg("no change to database")
		case migrate.ErrNilVersion:
			logger.Error().Err(err).Msg("no migration")
		default:
			logger.Error().Err(err).Msg("other error")
			panic(err)
		}
	}
}