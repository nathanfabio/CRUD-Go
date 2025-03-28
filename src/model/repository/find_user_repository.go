package repository

import (
	"context"
	"fmt"
	"os"

	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"github.com/nathanfabio/CRUD-Go/src/model"
	"github.com/nathanfabio/CRUD-Go/src/model/repository/entity"
	"github.com/nathanfabio/CRUD-Go/src/model/repository/entity/converter"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/zap"
)

func (ur *userRepository) FindUserByEmail(email string) (model.UserDomainInterface, *errs.Errs) {
	logger.Info("Init findUserByEmail repository", zap.String("journey", "findUserByEmail"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "email", Value: email}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this email: %s", email)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
			return nil, errs.NewNotFoundErrs(errorMessage)
		}

		errorMessage := "Error trying to find user by email"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByEmail"))
		return nil, errs.NewInternalServerErrs(errorMessage)
	}

	logger.Info("FindUserByEmail repository executed successfully",
	zap.String("journey", "findUserByEmail"),
	zap.String("email", email),
	zap.String("userID", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}


func (ur *userRepository) FindUserByID(id string) (model.UserDomainInterface, *errs.Errs) {
	logger.Info("Init findUserByID repository", zap.String("journey", "findUserByID"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userEntity := &entity.UserEntity{}

	filter := bson.D{{Key: "_id", Value: id}}
	err := collection.FindOne(context.Background(), filter).Decode(userEntity)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			errorMessage := fmt.Sprintf("User not found with this ID: %s", id)
			logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
			return nil, errs.NewNotFoundErrs(errorMessage)
		}

		errorMessage := "Error trying to find user by ID"
		logger.Error(errorMessage, err, zap.String("journey", "findUserByID"))
		return nil, errs.NewInternalServerErrs(errorMessage)
	}

	logger.Info("FindUserByID repository executed successfully",
	zap.String("journey", "findUserByID"),
	zap.String("userID", userEntity.ID.Hex()))

	return converter.ConvertEntityToDomain(*userEntity), nil
}