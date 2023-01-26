package users

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/ngoldack/dicetrace/apps/api/internal/utils"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
	"schneider.vip/problem"
)

var valid = validator.New()

type PostUserRequest struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" validate:"required"`
}

type PostUserResponse struct {
	Status int   `json:"status"`
	User   *User `json:"user"`
}

func RegisterUserHandler(userRepository UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check content type
		prob := utils.CheckApplicationJson(c)
		if prob != nil {
			log.Debug().Err(prob).Msg("wrong content type")
			utils.HandleProblem(c, prob)
			return
		}

		// read body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Debug().Err(err).Interface("body", c.Request.Body).Msg("cannot read body")
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("cannot read body"),
				problem.Detailf("body should be %s", utils.APPLICATION_JSON),
			))
			return
		}

		// unmarshal body
		request := &PostUserRequest{}
		err = json.Unmarshal(body, &request)
		if err != nil {
			log.Debug().Err(err).Msg("cannot unmarshal body")
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("cannot unmarshal request body"),
				problem.Detailf("provide a valid %s request body", utils.APPLICATION_JSON),
			))
			return
		}

		// validate body
		err = valid.Struct(request)
		if err != nil {
			log.Debug().Err(err).Msg("invalid request body")
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("invalid request body"),
				problem.Detailf("provide a valid %s request body", utils.APPLICATION_JSON),
			))
			return
		}

		user, err := userRepository.RegisterUser(&User{
			Username: request.Username,
			Email:    request.Email,
			Name:     request.Name,
		})
		if err != nil {
			log.Debug().Err(err).Msg("cannot create user in database")
			utils.HandleProblem(c, problem.Of(http.StatusInternalServerError).Append(
				problem.Title("cannot create user in database"),
			))
			return
		}

		// return the response
		response := &PostUserResponse{
			Status: http.StatusCreated,
			User:   user,
		}
		c.JSON(response.Status, response)
	}
}
