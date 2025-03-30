package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"github.com/nathanfabio/CRUD-Go/src/configuration/validation"
	"github.com/nathanfabio/CRUD-Go/src/controller/model/request"
	"github.com/nathanfabio/CRUD-Go/src/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) UpdateUser(c *gin.Context) {
	logger.Info("Init UpdateUser controller",
		zap.String("journey", "updateUser"),
	)
	var userRequest request.UserUpdateRequest

	
	//ShouldBindJSON binds the JSON passed in the request body to the struct.
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, 
			zap.String("journey", "updateUser"),
	)
		//c.JSON returns a JSON response with the status code and the JSON passed in the response body.
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return

	}

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Invalid userID, must be hex value", err, zap.String("journey", "updateUser"))
		errRest := errs.NewBadRequestErrs("Invalid userID, must be hex value")
		c.JSON(errRest.Code, errRest)
	}

	domain := model.NewUserUpdateDomain(userRequest.Name, userRequest.Age)

	err := uc.service.UpdateUserServices(userId, domain)
	if err != nil {
		logger.Error("Error to call UpdateUser service", err, zap.String("journey", "updateUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User updated successfully", 
		zap.String("userId", userId),
		zap.String("journey", "updateUser"),
)

	c.Status(http.StatusOK)
}