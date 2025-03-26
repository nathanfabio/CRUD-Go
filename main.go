package main

import (
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nathanfabio/CRUD-Go/src/configuration/database/mongodb"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"github.com/nathanfabio/CRUD-Go/src/controller/routes"
)

func main() {
	logger.Info("Starting the application...")

	err := godotenv.Load()
	if err != nil {
    log.Fatal("Error loading .env file")
}

	database, err := mongodb.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf("Error to connect to database, error=%s\n", err.Error())
		return
	}

	userController := initDependencies(database)
	
	

	router:= gin.Default() //Create a new router. Could be gin.New(), but gin.New() initializes a router without any middleware. While
	//gin.Default() initializes the router with logger and recovery middleware.
	
	routes.InitRoutes(&router.RouterGroup, userController)
	
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}