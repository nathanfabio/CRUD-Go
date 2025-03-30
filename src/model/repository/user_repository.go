package repository

import (
	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func NewUserRepository(database *mongo.Database) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct {
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *errs.Errs)
	FindUserByEmail(email string) (model.UserDomainInterface, *errs.Errs)
	FindUserByID(id string) (model.UserDomainInterface, *errs.Errs)
	UpdateUser(userId string, userDomain model.UserDomainInterface) *errs.Errs
}