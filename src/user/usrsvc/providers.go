package usrsvc

import (
	"sync"
	"tkkt-auth/src/domain"
)

var (
	service     *Service
	serviceOnce sync.Once
)

func ProvideService(repo domain.Repository) *Service {
	serviceOnce.Do(func() {
		service = New(repo)
	})
	return service
}
