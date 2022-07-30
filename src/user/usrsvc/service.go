package usrsvc

import (
	"context"
	"tkkt-auth/src/domain"
)

type Service struct {
	repo domain.Repository
}

func New(repo domain.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetOne(ctx context.Context, id int) (*domain.User, error) {
	return s.repo.User().GetOne(ctx, id)
}
