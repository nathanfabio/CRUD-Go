package controller

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/validation"
	"github.com/nathanfabio/CRUD-Go/src/controller/model/request"
)

func CreateUser(c *gin.Context) {

	var userRequest request.UserRequest

	//ShouldBindJSON binds the JSON passed in the request body to the struct.
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := errs.NewBadRequestErrs(
			fmt.Sprintf("Incorrect fields, error=%s", err.Error()))
			//c.JSON returns a JSON response with the status code and the JSON passed in the response body.
			errRest := validation.ValidateUserError(err)
			c.JSON(errRest.Code, errRest)
			return

	}

	fmt.Println(userRequest)
}