package service

import (
	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userDomainService) DeleteUserServices(userId string) *errs.Errs {
	logger.Info("Init DeleteUser model", zap.String("journey", "deleteUser"))


	err := ud.userRepository.DeleteUser(userId)
	if err != nil {
		logger.Error("Init DeleteUser model", err, zap.String("journey", "deleteUser"))
		return err
	}

	logger.Info("DeleteUser service executed successfuly",
		zap.String("userId", userId),
		zap.String("journey",  "deleteUser"),
)

	return nil
}