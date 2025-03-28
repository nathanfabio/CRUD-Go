package service

import (

	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"github.com/nathanfabio/CRUD-Go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, *errs.Errs) {
	logger.Info("Init CreateUser model", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil {
		logger.Error("Init CreateUser model", err, zap.String("journey", "createUser"))
		return nil, err
	}

	logger.Info("CreateUSer service executed successfuly",
		zap.String("userId", userDomainRepository.GetID()),
		zap.String("journey", "createUser"),
)

	return userDomainRepository, nil
}