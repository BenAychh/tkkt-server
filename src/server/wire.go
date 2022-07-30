//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"tkkt-auth/src/config"
	"tkkt-auth/src/db"
	"tkkt-auth/src/domain"
	"tkkt-auth/src/repo"
	"tkkt-auth/src/user/usrhndlr"
	"tkkt-auth/src/user/usrsvc"
)

func Wire() (*Server, error) {
	panic(wire.Build(
		// config
		config.ProvideConfig,

		// db
		db.ProvideDB,

		// repo
		repo.ProvideRepo,
		wire.Bind(new(domain.Repository), new(*repo.Repo)),

		// services
		usrsvc.ProvideService,
		wire.Bind(new(domain.UserService), new(*usrsvc.Service)),

		//handlers
		usrhndlr.ProvideHandler,
		wire.Bind(new(domain.UserHandler), new(*usrhndlr.Handler)),

		// server itself
		ProvideServer,
	))
}
