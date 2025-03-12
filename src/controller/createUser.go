package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"github.com/nathanfabio/CRUD-Go/src/configuration/validation"
	"github.com/nathanfabio/CRUD-Go/src/controller/model/request"
	"github.com/nathanfabio/CRUD-Go/src/controller/model/response"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
	logger.Info("Init CreateUser controller",
		zap.String("journey", "Create a new user"),
	)
	var userRequest request.UserRequest

	//ShouldBindJSON binds the JSON passed in the request body to the struct.
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, 
			zap.String("journey", "Create a new user"),
	)
		//c.JSON returns a JSON response with the status code and the JSON passed in the response body.
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return

	}

	
	response := response.UserResponse {
		ID: "test",
		Email: userRequest.Email,
		Name: userRequest.Name,
		Age: userRequest.Age,
	}

	logger.Info("User created successfully", 
		zap.String("journey", "Create a new user"),
)

	c.JSON(http.StatusOK, response)
}