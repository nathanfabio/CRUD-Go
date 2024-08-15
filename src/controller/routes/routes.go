package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/CRUD-Go/src/controller"
)

// InitRoutes init all routes
func InitRoutes(r *gin.RouterGroup) {
	r.GET("/getUserById/:userId", controller.FindUserByID) //I can pass more than one handler. One of them could be middleware, for example.
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deleteUser/:userId", controller.DeleteUser)
}