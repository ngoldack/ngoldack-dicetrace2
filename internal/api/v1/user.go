package v1

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/ngoldack/dicetrace/internal/controller"
	"github.com/ngoldack/dicetrace/internal/models"
	"github.com/ngoldack/dicetrace/internal/utils"
	"io"
	"net/http"
	"schneider.vip/problem"
	"strconv"
	"time"
)

type GetUserResponse struct {
	Status int           `json:"status,omitempty"`
	Size   int           `json:"size,omitempty"`
	Users  []models.User `json:"users,omitempty"`
}

type PostUserRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Name     string `json:"name,omitempty"`
}

type PostUserResponse struct {
	Status int          `json:"status,omitempty"`
	User   *models.User `json:"user,omitempty"`
}

type GetUserWithIdResponse struct {
	Status int          `json:"status,omitempty"`
	User   *models.User `json:"user,omitempty"`
}

// HandleGetUser returns multiple users from the database
func HandleGetUser(userController *controller.UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Query mapping: -1 represents not present
		pageQuery, sizeQuery := c.Query("page"), c.Query("size")
		page, size := -1, -1
		var err error

		// Convert page to int
		if pageQuery != "" {
			page, err = strconv.Atoi(pageQuery)
			if err != nil {
				utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
					problem.Title("page is not a number"),
					problem.Detailf("provided page '%s' is not a number", pageQuery),
					problem.Wrap(err)),
				)
				return
			}
		}

		// Convert size to int
		if sizeQuery != "" {
			size, err = strconv.Atoi(sizeQuery)
			if err != nil {
				utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
					problem.Title("size is not a number"),
					problem.Detailf("provided size '%s' is not a number", pageQuery),
					problem.Wrap(err)),
				)
				return
			}
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
		defer cancel()

		// Get users from controller
		users, err := userController.GetUser(ctx, page, size)
		if err != nil {
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("could not get users from database"),
				problem.Wrap(err)),
			)
			return
		}

		// Create response
		response := &GetUserResponse{
			Status: http.StatusOK,
			Size:   len(users),
			Users:  users,
		}

		c.JSON(response.Status, response)
	}
}

func HandlePostUser(userController *controller.UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check for content type
		if c.ContentType() != utils.APPLICATION_JSON {
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("request body is not application/json"),
				problem.Detail("request body should be application/json"),
			))
			return
		}

		// Read the response body
		rawBody, err := io.ReadAll(c.Request.Body)
		if err != nil {
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("could not read request body"),
				problem.Detail("request body should be a application/json"),
				problem.Wrap(err),
			))
			return
		}
		var requestBody PostUserRequest
		err = json.Unmarshal(rawBody, &requestBody)
		if err != nil {
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("could not unmarshal request body"),
				problem.Detail("request body should be a valid request body"),
				problem.Wrap(err),
			))
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
		defer cancel()

		user, err := userController.PostUser(ctx, requestBody.Username, requestBody.Email, requestBody.Name)
		if err != nil {
			utils.HandleProblem(c, problem.Of(http.StatusInternalServerError).Append(
				problem.Title("could not add user"),
				problem.Wrap(err),
			))
			return
		}

		response := &PostUserResponse{
			Status: http.StatusOK,
			User:   user,
		}

		c.JSON(response.Status, response)
	}
}

func HandleGetUserWithUserID(userController *controller.UserController) gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.Param("user_id")

		// Check if user id is empty
		if userId == "" {
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("user_id is empty"),
				problem.Detail("user_id should contain a valid user_id"),
			))
			return
		}

		ctx, cancel := context.WithTimeout(c.Request.Context(), time.Second*5)
		defer cancel()

		user, err := userController.GetUserWithUserId(ctx, userId)
		if err != nil {
			utils.HandleProblem(c, problem.Of(http.StatusBadRequest).Append(
				problem.Title("user_id is empty"),
				problem.Detail("user_id should contain a valid user_id"),
				problem.Wrap(err),
			))
			return
		}

		response := &GetUserWithIdResponse{
			Status: http.StatusOK,
			User:   user,
		}

		c.JSON(response.Status, response)
	}
}
