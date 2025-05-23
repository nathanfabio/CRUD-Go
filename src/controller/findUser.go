package controller

import (
	"fmt"
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"github.com/nathanfabio/CRUD-Go/src/model"
	"github.com/nathanfabio/CRUD-Go/src/view"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) FindUserByID(c *gin.Context) {
	logger.Info("Init findUserByID controller", zap.String("journey", "findUserById"))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	userId := c.Param("userId")

	if _ , err := primitive.ObjectIDFromHex(userId); err != nil {
		logger.Error("Error to validate userId", err, zap.String("journey", "findUserById"))
		errorMessage := errs.NewBadRequestErrs("User id isn't valid is")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByIDServices(userId)
	if err != nil {
		logger.Error("Error to to call finduserById services", err, zap.String("journey", "findUserById"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByID controller executed successfully", zap.String("journey", "findUserById"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}


func (uc *userControllerInterface) FindUserByEmail(c *gin.Context) {
	logger.Info("Init findUserByEmail controller", zap.String("journey", "findUserByEmail"))

	user, err := model.VerifyToken(c.Request.Header.Get("Authorization"))
	if err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info(fmt.Sprintf("User authenticated: %#v", user))

	userEmail := c.Param("userEmail")

	if _ , err := mail.ParseAddress(userEmail); err != nil {
		logger.Error("Error to validate Email", err, zap.String("journey", "findUserByEmail"))
		errorMessage := errs.NewBadRequestErrs("User id isn't valid is")
		c.JSON(errorMessage.Code, errorMessage)
		return
	}

	userDomain, err := uc.service.FindUserByEmailServices(userEmail)
	if err != nil {
		logger.Error("Error to to call finduserByEmail services", err, zap.String("journey", "findUserByEmail"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("FindUserByEmail controller executed successfully", zap.String("journey", "findUserByEmail"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(userDomain))
}