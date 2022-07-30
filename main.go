package main

import (
	"context"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"tkkt-auth/src/server"
)

func main() {
	app, err := server.Wire()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to wire")
	}
	go app.Start(context.Background())

	quit := make(chan os.Signal)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need to add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	app.Stop(timeoutCtx)
}
