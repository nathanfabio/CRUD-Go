package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/nathanfabio/CRUD-Go/src/controller/routes"
)

func main() {

	err := godotenv.Load()
	if err != nil {
    log.Fatal("Error loading .env file")
}

	router:= gin.Default() //Create a new router. Could be gin.New(), but gin.New() initializes a router without any middleware. While
	//gin.Default() initializes the router with logger and recovery middleware.

	routes.InitRoutes(&router.RouterGroup)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}