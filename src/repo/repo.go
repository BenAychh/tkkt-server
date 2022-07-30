package repo

import (
	"context"
	"github.com/jmoiron/sqlx"
	"tkkt-auth/src/domain"
)

type sqlxer interface {
	sqlx.QueryerContext
	sqlx.ExecerContext
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type Repo struct {
	_db *sqlx.DB
	db  sqlxer
}

func New(db *sqlx.DB) *Repo {
	return &Repo{db, db}
}

func (r *Repo) User() domain.UserRepository {
	return &userRepo{r.db}
}
