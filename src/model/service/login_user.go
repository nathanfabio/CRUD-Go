package service


import (

	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/configuration/logger"
	"github.com/nathanfabio/CRUD-Go/src/model"
	"go.uber.org/zap"
)

func (ud *userDomainService) LoginUserServices(userDomain model.UserDomainInterface) (model.UserDomainInterface, string, *errs.Errs) {
	logger.Info("Init LoginUser model", zap.String("journey", "loginUser"))
	
	userDomain.EncryptPassword()

	user, err := ud.findUserByEmailAndPasswordServices(userDomain.GetEmail(), userDomain.GetPassword())
	if err != nil {
		return nil, "", err
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, "", err
	}
	
	logger.Info("LoginUSer service executed successfuly",
		zap.String("userId", user.GetID()),
		zap.String("journey", "loginUser"),
)

	return user, token, nil
}