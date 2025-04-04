package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) DeleteUser(c *gin.Context) {
	logger.Info("Init DeleteUser controller",
		zap.String("journey", "deleteUser"),
	)

	userId := c.Param("userId")
	if _, err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Invalid userID, must be hex value", err, zap.String("journey", "deleteUser"))
		errRest := errs.NewBadRequestErrs("Invalid userID, must be hex value")
		c.JSON(errRest.Code, errRest)
	}

	err := uc.service.DeleteUserServices(userId)
	if err != nil {
		logger.Error("Error to call DeleteUser service", err, zap.String("journey", "deleteUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User Deleted successfully", 
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
)

	c.Status(http.StatusOK)
}