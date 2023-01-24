package api

import (
	"context"
	"fmt"
	"github.com/dn365/gin-zerolog"
	"github.com/gin-gonic/gin"
	v1 "github.com/ngoldack/dicetrace/internal/api/v1"
	"github.com/ngoldack/dicetrace/internal/app"
	"github.com/ngoldack/dicetrace/internal/controller"
	"net/http"
)

type API struct {
	engine *gin.Engine
	srv    *http.Server
}

// Force interface implementation
var _ app.Controllable = &API{}

func NewAPI(port string, userController *controller.UserController) *API {
	router := &API{}
	router.engine = gin.New()
	router.engine.Use(gin.Recovery())
	router.engine.Use(ginzerolog.Logger("gin"))

	router.srv = &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.engine,
	}

	router.configureRoutes(userController)

	return router
}

func (router *API) configureRoutes(userController *controller.UserController) {
	router.engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "%s", "pong")
	})

	// V1
	apiRoutes := router.engine.Group("/api/v1")
	apiRoutes.GET("/user", v1.HandleGetUser(userController))
	apiRoutes.POST("/user", v1.HandlePostUser(userController))
	apiRoutes.GET("/user/:user_id", v1.HandleGetUserWithUserID(userController))
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
