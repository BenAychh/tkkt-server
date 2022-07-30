package repo

import (
	"context"
	"tkkt-auth/src/domain"
)

type userRepo struct {
	db sqlxer
}

func newUserRepo(db sqlxer) *userRepo {
	return &userRepo{db: db}
}

func (r *userRepo) GetOne(ctx context.Context, id int) (*domain.User, error) {
	// language=SQL
	query := `
SELECT id,
       forename AS name
FROM users
WHERE id = $1
`
	args := []interface{}{id}
	var user domain.User
	err := r.db.GetContext(ctx, &user, query, args...)
	return &user, err
}
