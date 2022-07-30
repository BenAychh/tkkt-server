package db

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func newDB(user, password, host, db string) (*sqlx.DB, error) {
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", user, password, host, db)
	return sqlx.Connect("postgres", connectionString)
}
