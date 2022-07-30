package server

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"tkkt-auth/src/config"
	"tkkt-auth/src/domain"
	"tkkt-auth/src/utilities"
)

type Server struct {
	config      *config.Config
	userHandler domain.UserHandler
	router      *gin.Engine
	httpServer  *http.Server
}

func (s *Server) Start(ctx context.Context) {
	s.router = gin.Default()
	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.config.HTTP().Port()),
		Handler: s.router,
	}
	api := s.router.Group("/api")
	v1 := api.Group("/v1")
	v1.GET("/ping", func(c *gin.Context) {
		utilities.RespondData(c, http.StatusOK, "pong", "a user message")
	})
	s.setupRoutes(v1)
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			log.Info().Err(err).Msg("listen")
		}
	}()
}

func (s *Server) setupRoutes(entryPoint *gin.RouterGroup) {
	s.userHandler.HandleUserGroup(s.router.Group("/user"))
}

func (s *Server) Stop(ctx context.Context) {
	log.Info().Msg("shutting down gracefully")
	if err := s.httpServer.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("server forced to shutdown")
	}
	log.Info().Msg("server exiting")
}
