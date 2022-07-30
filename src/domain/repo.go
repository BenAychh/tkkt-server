package domain

import "context"

type (
	User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
	UserRepository interface {
		GetOne(ctx context.Context, id int) (*User, error)
	}
	Repository interface {
		User() UserRepository
	}
)
