package service

import (
	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"github.com/nathanfabio/CRUD-Go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) UpdateUserServices(userId string, userDomain model.UserDomainInterface) *errs.Errs {
	logger.Info("Init UpdateUser model", zap.String("journey", "UudateUser"))


	err := ud.userRepository.UpdateUser(userId, userDomain)
	if err != nil {
		logger.Error("Init UpdateUser model", err, zap.String("journey", "updateUser"))
		return err
	}

	logger.Info("UpdateUser service executed successfuly",
		zap.String("userId", userDomain.GetEmail()),
		zap.String("journey", "updateUser"),
)

	return nil
}