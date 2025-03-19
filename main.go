package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"github.com/nathanfabio/CRUD-Go/src/controller"
	"github.com/nathanfabio/CRUD-Go/src/controller/routes"
	"github.com/nathanfabio/CRUD-Go/src/model/service"
)

func main() {
	logger.Info("Starting the application...")

	err := godotenv.Load()
	if err != nil {
    log.Fatal("Error loading .env file")
}

	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router:= gin.Default() //Create a new router. Could be gin.New(), but gin.New() initializes a router without any middleware. While
	//gin.Default() initializes the router with logger and recovery middleware.

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}