package main

import (
	"context"
	"github.com/ngoldack/dicetrace/apps/api/internal/api"
	"github.com/ngoldack/dicetrace/apps/api/internal/app"
	"github.com/ngoldack/dicetrace/apps/api/internal/database"
	"github.com/ngoldack/dicetrace/apps/api/internal/users"
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

	// Neo4J client
	neo4jc, err := database.NewNeo4JClient(cfg.Neo4JURI, cfg.Neo4JUsername, cfg.Neo4JPassword)

	userRepository := &users.UserNeo4JRepository{Driver: neo4jc.Driver}

	// REST api
	router := api.NewAPI("8080")

	router.GetRouterGroupV1().POST("/user", users.RegisterUserHandler(userRepository))
	router.GetRouterGroupV1().GET("/user/:user-uuid", users.GetUserWithUsernameHandler(userRepository))
	router.GetRouterGroupV1().DELETE("/user/:user-uuid", users.DeleteUserWithUsernameHandler(userRepository))

	router.GetRouterGroupV1().GET("/find/user", users.FindUserHandler(userRepository))

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

	// Rest API
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = router.Stop(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to stop api server")
	}

	// Neo4jJ client
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err = neo4jc.Stop(ctx); err != nil {
		log.Error().Err(err).Msg("Failed to stop api server")
	}
}
