package service

import (
	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"github.com/nathanfabio/CRUD-Go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) FindUserByIDServices(id string) (model.UserDomainInterface, *errs.Errs) {
	logger.Info("Init findUserByID services", zap.String("journey", "findUserById"))

	return ud.userRepository.FindUserByID(id)
}


func (ud *userDomainService) FindUserByEmailServices(email string) (model.UserDomainInterface, *errs.Errs) {
	logger.Info("Init findUserByEmail services", zap.String("journey", "findUserByEmail"))

	return ud.userRepository.FindUserByEmail(email)

}


