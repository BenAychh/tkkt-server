package db

import (
	"github.com/jmoiron/sqlx"
	"log"
	"sync"
	"tkkt-auth/src/config"
)

var (
	db     *sqlx.DB
	dbOnce sync.Once
)

func ProvideDB(config *config.Config) *sqlx.DB {
	dbOnce.Do(func() {
		var err error
		dbConfig := config.DB()
		db, err = newDB(dbConfig.User(), dbConfig.Password(), dbConfig.Host(), dbConfig.DB())
		if err != nil {
			log.Fatalf("failed to create db %v", err)
		}
	})
	return db
}
