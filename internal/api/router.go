package api

import (
	"context"
	"fmt"
	"github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/ngoldack/dicetrace/internal/app"
	"net/http"
)

type API struct {
	engine *gin.Engine
	srv    *http.Server
}

// Force interface implementation
var _ app.Controllable = &API{}

func NewAPI(port string) *API {
	router := &API{}
	router.engine = gin.New()
	router.engine.Use(gin.Recovery())
	router.engine.Use(ginzerolog.Logger("gin"))

	router.srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.engine,
	}

	router.configureRoutes()

	return router
}

func (router *API) configureRoutes() {
	router.engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "pong")
	})
}

func (router *API) Start(_ context.Context) error {
	return router.srv.ListenAndServe()
}

func (router *API) Stop(ctx context.Context) error {
	err := router.srv.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("error while stopping api server: %w", err)
	}
	return nil
}
