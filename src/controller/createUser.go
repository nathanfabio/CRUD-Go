package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"github.com/nathanfabio/CRUD-Go/src/configuration/validation"
	"github.com/nathanfabio/CRUD-Go/src/controller/model/request"
	"github.com/nathanfabio/CRUD-Go/src/model"
	"github.com/nathanfabio/CRUD-Go/src/view"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
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

	domain := model.NewUserDomain(userRequest.Email, userRequest.Password, userRequest.Name, userRequest.Age)
	
	

	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully", 
		zap.String("journey", "Create a new user"),
)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domain,
	))
}