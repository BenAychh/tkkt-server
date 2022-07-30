package usrhndlr

import (
	"sync"
	"tkkt-auth/src/domain"
)

var (
	handler     *Handler
	handlerOnce sync.Once
)

func ProvideHandler(service domain.UserService) *Handler {
	handlerOnce.Do(func() {
		handler = New(service)
	})
	return handler
}
