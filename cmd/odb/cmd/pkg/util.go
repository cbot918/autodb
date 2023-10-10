package pkg

import (
	"database/sql"
	"os"

	"github.com/cbot918/autodb/internal"
)

func IsFileExists(path string) bool {
	_, err := os.Stat(path)

	return err == nil
}

func Init() (*internal.Config, *sql.DB, error, error) {

	cfg, cfgErr := internal.NewConfig()
	db, dbErr := internal.NewDB(cfg)

	return cfg, db, cfgErr, dbErr
}
