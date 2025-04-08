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


func (uc *userControllerInterface) LoginUser(c *gin.Context) {
	logger.Info("Init LoginUser controller",
		zap.String("journey", "loginUser"),
	)
	var userLogin request.UserLogin

	//ShouldBindJSON binds the JSON passed in the request body to the struct.
	if err := c.ShouldBindJSON(&userLogin); err != nil {
		logger.Error("Error trying to validate user info", err, 
			zap.String("journey", "loginUser"),
	)
		//c.JSON returns a JSON response with the status code and the JSON passed in the response body.
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Code, errRest)
		return

	}

	domain := model.NewUserLoginDomain(userLogin.Email, userLogin.Password)
	
	

	domainResult, err := uc.service.LoginUserServices(domain)
	if err != nil {
		logger.Error("Error to call LoginUser service", err, zap.String("journey", "loginUser"))
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User Logind successfully", 
		zap.String("userId", domainResult.GetID()),
		zap.String("journey", "loginUser"),
)

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(
		domainResult,
	))
}