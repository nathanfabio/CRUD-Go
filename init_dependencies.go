package main

import (
	"github.com/nathanfabio/CRUD-Go/src/controller"
	"github.com/nathanfabio/CRUD-Go/src/model/repository"
	"github.com/nathanfabio/CRUD-Go/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)

	return controller.NewUserControllerInterface(service)
}