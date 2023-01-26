package api

import (
	"context"
	"fmt"
	"github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	"github.com/ngoldack/dicetrace/apps/api/internal/app"
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

func (api *API) GetRouterGroupV1() *gin.RouterGroup {
	return api.engine.Group("/api/v1")
}

func (api *API) configureRoutes() {
	api.engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "pong")
	})
}

func (api *API) Start(_ context.Context) error {
	return api.srv.ListenAndServe()
}

func (api *API) Stop(ctx context.Context) error {
	err := api.srv.Shutdown(ctx)
	if err != nil {
		return fmt.Errorf("error while stopping api server: %w", err)
	}
	return nil
}
