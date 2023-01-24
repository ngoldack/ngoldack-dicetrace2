package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"net/http"
	"schneider.vip/problem"
)

func HandleProblem(c *gin.Context, problem *problem.Problem) {
	_, err := problem.WriteTo(c.Writer)
	if err != nil {
		errs := make([]error, 0, 2)
		errs = append(errs, err, problem)
		log.Error().Errs("errors", errs).Str("problem", problem.JSONString()).Msg("error while returning problem")
		c.String(http.StatusInternalServerError, "error while returning problem")
	}
}
