package user

import (
	"context"
	"github.com/ngoldack/dicetrace/apps/user/internal"
	"github.com/rs/zerolog/log"
	"os"
	"os/signal"
	"time"
)

func main() {

	log.Info().Msg("starting user...")
	ctx := context.Background()

	// config
	cfg, err := internal.GetConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to get config")
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

	log.Info().Msg("user started!")

	request := make(chan []byte)
	errs := make(chan error)
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	userNewConsumer := kafka.Consumer("user/new")

	//Listen for messages
	go func() {
		for {
			message, err := userNewConsumer.ReadMessage(context.TODO())
			if err != nil {
				errs <- err
				return
			}

			log.Info().Str("message", string(message.Value)).Msg("received message")
		}
	}()

	run(request, errs, quit)
	log.Info().Msg("gracefully stopping user...")

	// Kafka client
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = kafka.Stop(ctx); err != nil {
		log.Error().Err(err).Msg("failed to stop kafka client")
	}

	log.Info().Msg("user stopped!")
}

func run(request chan []byte, errs chan error, quit chan os.Signal) {

	for {
		select {
		case err := <-errs:
			log.Error().Err(err).Msg("error received")
		case data := <-request:
			log.Debug().Interface("request", data).Msg("new request received")
		case <-quit:
			return
		}
	}

}
