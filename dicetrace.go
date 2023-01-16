package main

import (
	"context"
	"github.com/ngoldack/dicetrace/internal/api"
	"github.com/ngoldack/dicetrace/internal/app"
	"github.com/ngoldack/dicetrace/internal/database"
	"github.com/rs/zerolog/log"
	"gopkg.in/errgo.v2/errors"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	log.Info().Msg("Starting dicetrace backend")
	ctx := context.Background()

	// Config
	cfg, err := app.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get config")
	}

	// Database
	db, err := database.NewDatabase(cfg.DatabaseURI())
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create database client")
	}
	if err = db.Start(ctx); err != nil {
		log.Fatal().Err(err).Msg("Failed to connect to database")
	}

	// REST api
	router := api.NewAPI("8080")
	go func() {
		if err = router.Start(ctx); err != nil {
			if ok := errors.Is(http.ErrServerClosed); !ok(err) {
				log.Fatal().Err(err).Msg("Failed to start http server")
			}
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info().Msg("Gracefully stopping dicetrace backend")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Rest API
	if err = router.Stop(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to stop api server")
	}

	// Database
	if err = db.Stop(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to close database connection")
	}
}
