package repository

import (
	"context"
	"os"

	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"github.com/nathanfabio/CRUD-Go/src/model"
)

const (
	MONGODB_USER_COLLECTION = "MONGODB_USER_COLLECTION"
)

func (ur *userRepository) CreateUser(userDomain model.UserDomainInterface) (model.UserDomainInterface, *errs.Errs) {
	logger.Info("Init createUser repository")

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	value, err := userDomain.GetJSONValue()
	if err != nil {
		return nil, errs.NewInternalServerErrs(err.Error())
	}

	result, err := collection.InsertOne(context.Background(), value)
	if err != nil {
		return nil, errs.NewInternalServerErrs(err.Error())
	}

	userDomain.SetID(result.InsertedID.(string))

	return userDomain, nil
}