package internal

import (
	"context"
	"fmt"
	"net/http"

	core "github.com/ngoldack/dicetrace/libs/util-core"

	ginzerolog "github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
)

type API struct {
	engine *gin.Engine
	srv    *http.Server
}

// Force interface implementation
var _ core.Controllable = &API{}

func NewAPI(cfg *Config) *API {
	router := &API{}
	router.engine = gin.New()
	router.engine.Use(gin.Recovery())
	router.engine.Use(ginzerolog.Logger("gin"))

	router.srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ApiPort),
		Handler: router.engine,
	}

	router.configureRoutes()

	router.GetRouterGroupV1(cfg).GET("/test")

	return router
}

func (api *API) GetRouterGroupV1(cfg *Config) *gin.RouterGroup {
	return api.engine.Group("/api/v1", checkJWT(cfg))
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
