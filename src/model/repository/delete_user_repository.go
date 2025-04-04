package repository


import (
	"context"
	"os"

	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

func (ur *userRepository) DeleteUser(userId string) *errs.Errs {
	logger.Info("Init DeleteUser repository", zap.String("jouney", "deleteUser"))

	collection_name := os.Getenv(MONGODB_USER_COLLECTION)

	collection := ur.databaseConnection.Collection(collection_name)

	userIdHex, _ := primitive.ObjectIDFromHex(userId)

	filter := bson.D{{Key: "_id", Value: userIdHex}}

	_, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Error("Error to delete user", err, zap.String("journey", "deleteUser"))
		return errs.NewInternalServerErrs(err.Error())
	}

	

	logger.Info("DeleteUser repository executed successfully",
		zap.String("userId", userId),
		zap.String("journey", "deleteUser"),
	)

	return nil
}