package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nathanfabio/CRUD-Go/src/controller"
)

// InitRoutes init all routes
func InitRoutes(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/getUserById/:userId", userController.FindUserByID) //I can pass more than one handler. One of them could be middleware, for example.
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deleteUser/:userId", userController.DeleteUser)

	r.POST("/login", userController.LoginUser)
}