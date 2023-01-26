package users

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
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

type GetUserWithUsernameResponse struct {
	Status int   `json:"status"`
	User   *User `json:"user"`
}

type FindUserRequest struct {
	Username string `json:"username" validate:"required"`
}

type FindUserResponse struct {
	Status int    `json:"status"`
	Size   int    `json:"size"`
	Users  []User `json:"users"`
}

func FindUserHandler(repository UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check content type
		if prob := utils.CheckApplicationJson(c); prob != nil {
			log.Debug().Str("system", "api/user").Err(prob).Msg("wrong content type")
			utils.HandleProblem(c, prob)
			return
		}

		// read body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Debug().Str("system", "api/user").Err(err).Interface("body", c.Request.Body).Msg("cannot read body")
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("cannot read body"),
				problem.Detailf("body should be %s", utils.ApplicationJson),
				problem.Instance(utils.ErrorRequestMalformedRequestBody),
			))
			return
		}

		// unmarshal body
		request := &FindUserRequest{}
		err = json.Unmarshal(body, &request)
		if err != nil {
			log.Debug().Str("system", "api/user").Err(err).Msg("cannot unmarshal body")
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("cannot unmarshal request body"),
				problem.Detailf("provide a valid %s request body", utils.ApplicationJson),
				problem.Instance(utils.ErrorRequestMalformedRequestBody),
			))
			return
		}

		// validate body
		err = valid.Struct(request)
		if err != nil {
			log.Debug().Str("system", "api/user").Err(err).Msg("invalid request body")
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("invalid request body"),
				problem.Detailf("provide a valid %s request body", utils.ApplicationJson),
				problem.Instance(utils.ErrorRequestInvalidRequestBody),
			))
			return
		}

		users, err := repository.FindUserWithUsername(request.Username)
		if err != nil {
			log.Debug().Str("system", "api/user").Err(err).Msg("no user found")
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("no user found"),
				problem.Instance(utils.ErrorDatabaseNodeNotFound),
			))
			return
		}

		response := &FindUserResponse{
			Status: http.StatusOK,
			Size:   len(users),
			Users:  users,
		}

		c.JSON(response.Status, response)
	}
}

func RegisterUserHandler(userRepository UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check content type
		if prob := utils.CheckApplicationJson(c); prob != nil {
			log.Debug().Str("system", "api/user").Err(prob).Msg("wrong content type")
			utils.HandleProblem(c, prob)
			return
		}

		// read body
		body, err := io.ReadAll(c.Request.Body)
		if err != nil {
			log.Debug().Str("system", "api/user").Err(err).Interface("body", c.Request.Body).Msg("cannot read body")
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("cannot read body"),
				problem.Detailf("body should be %s", utils.ApplicationJson),
			))
			return
		}

		// unmarshal body
		request := &PostUserRequest{}
		err = json.Unmarshal(body, &request)
		if err != nil {
			log.Debug().Str("system", "api/user").Err(err).Msg("cannot unmarshal body")
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("cannot unmarshal request body"),
				problem.Detailf("provide a valid %s request body", utils.ApplicationJson),
			))
			return
		}

		// validate body
		err = valid.Struct(request)
		if err != nil {
			log.Debug().Str("system", "api/user").Err(err).Msg("invalid request body")
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("invalid request body"),
				problem.Detailf("provide a valid %s request body", utils.ApplicationJson),
			))
			return
		}

		user, err := userRepository.RegisterUser(&User{
			Username: request.Username,
			Email:    request.Email,
			Name:     request.Name,
		})
		if err != nil {
			log.Debug().Str("system", "api/user").Err(err).Msg("cannot create user in database")
			utils.HandleProblem(c, problem.Of(http.StatusInternalServerError).Append(
				problem.Title("cannot create user in database"),
				problem.Instance(utils.ErrorDatabaseNodeNotCreated),
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

func GetUserWithUsernameHandler(userRepository UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("user-uuid"))
		if err != nil {
			log.Debug().Str("system", "api/user").Msg("invalid uuid provided")
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("invalid uuid provided"),
				problem.Detail("provide a valid uuid"),
				problem.Instance(utils.ErrorRequestInvalidUuidProvided),
			))
			return
		}

		user, err := userRepository.GetUser(id)
		if err != nil {
			log.Debug().Str("system", "api/user").Err(err).Msg("cannot get user from database")
			utils.HandleProblem(c, problem.Of(http.StatusInternalServerError).Append(
				problem.Title("cannot get user from database"),
				problem.Instance(utils.ErrorDatabaseNodeNotFound),
			))
			return
		}

		response := GetUserWithUsernameResponse{
			Status: http.StatusOK,
			User:   user,
		}

		c.JSON(response.Status, response)
	}
}

func DeleteUserWithUsernameHandler(userRepository UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.Parse(c.Param("user-uuid"))
		if err != nil {
			log.Debug().Str("system", "api/user").Err(err).Msg("invalid uuid provided")
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("invalid uuid provided"),
				problem.Detail("provide a valid uuid"),
				problem.Instance(utils.ErrorRequestInvalidUuidProvided),
			))
			return
		}

		err = userRepository.DeleteUser(id)
		if err != nil {
			log.Debug().Str("system", "api/user").Err(err).Msg("user not deleted")
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("user not deleted"),
				problem.Instance(utils.ErrorDatabaseNodeNotDeleted),
			))
			return
		}

		c.Status(http.StatusOK)
	}
}
