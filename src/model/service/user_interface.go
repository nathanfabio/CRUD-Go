package service

import (
	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/model"
)

func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *errs.Errs
	UpdateUser(string, model.UserDomainInterface) *errs.Errs
	FindUser(string) (*model.UserDomainInterface, *errs.Errs)
	DeleteUser(string) *errs.Errs
}