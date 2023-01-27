package utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"schneider.vip/problem"
)

func CheckApplicationJson(c *gin.Context) *problem.Problem {
	if c.ContentType() != ApplicationJson {
		return problem.Of(http.StatusBadRequest).Append(
			problem.Title("wrong content type"),
			problem.Detailf("content type should be %s", ApplicationJson),
			problem.Instance(ErrorRequestInvalidContentType),
		)
	}
	return nil
}
