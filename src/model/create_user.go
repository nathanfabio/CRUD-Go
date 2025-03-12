package model

import (
	"fmt"

	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"go.uber.org/zap"
)

func (ud *userDomain) CreateUser() *errs.Errs {
	logger.Info("Init CreateUser model", zap.String("journey", "Create a new user"))

	ud.EncryptPassword()

	fmt.Println(ud)

	return nil
}