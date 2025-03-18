package service

import (
	"fmt"

	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"github.com/nathanfabio/CRUD-Go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(userDomain model.UserDomainInterface) *errs.Errs {
	logger.Info("Init CreateUser model", zap.String("journey", "Create a new user"))

	userDomain.EncryptPassword()

	fmt.Println(userDomain.GetPassword())

	return nil
}