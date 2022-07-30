package repo

import (
	"github.com/jmoiron/sqlx"
	"sync"
)

var (
	repo     *Repo
	repoOnce sync.Once
)

func ProvideRepo(db *sqlx.DB) *Repo {
	repoOnce.Do(func() {
		repo = New(db)
	})
	return repo
}
