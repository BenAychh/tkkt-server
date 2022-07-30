package server

import (
	"sync"
	"tkkt-auth/src/config"
	"tkkt-auth/src/domain"
)

var (
	svr     *Server
	svrOnce sync.Once
)

func ProvideServer(config *config.Config, userHandler domain.UserHandler) *Server {
	svrOnce.Do(func() {
		svr = &Server{
			config,
			userHandler,
			nil,
			nil,
		}
	})
	return svr
}
