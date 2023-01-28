package main

import (
	"context"
	"github.com/ngoldack/dicetrace/apps/api/internal"
	"github.com/rs/zerolog/log"
	"gopkg.in/errgo.v2/errors"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	log.Info().Msg("starting api...")
	ctx := context.Background()

	// config
	cfg, err := internal.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to get config")
	}

	// Kafka connection
	kafka, err := internal.NewKafkaClient(cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to create kafka client")
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err = kafka.Start(ctx); err != nil {
		log.Fatal().Err(err).Msg("failed to start kafka client")
	}
	// REST api
	router := internal.NewAPI(cfg)

	go func() {
		if err = router.Start(ctx); err != nil {
			if ok := errors.Is(http.ErrServerClosed); !ok(err) {
				log.Fatal().Err(err).Msg("Failed to start http server")
			}
		}
	}()

	log.Info().Msg("api started!")

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info().Msg("gracefully stopping api...")

	// Rest API
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = router.Stop(ctx); err != nil {
		log.Error().Err(err).Msg("failed to stop api server")
	}

	// Kafka client
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = kafka.Stop(ctx); err != nil {
		log.Error().Err(err).Msg("failed to stop kafka client")
	}

	log.Info().Msg("api stopped!")
}
