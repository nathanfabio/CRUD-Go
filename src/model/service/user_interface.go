package service

import (
	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
	"github.com/nathanfabio/CRUD-Go/src/model"
	"github.com/nathanfabio/CRUD-Go/src/model/repository"
)

func NewUserDomainService(userRepository repository.UserRepository) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository
}

type UserDomainService interface {
	CreateUserServices(model.UserDomainInterface) (model.UserDomainInterface, *errs.Errs)
	UpdateUserServices(string, model.UserDomainInterface) *errs.Errs
	FindUserByIDServices(id string) (model.UserDomainInterface, *errs.Errs)
	FindUserByEmailServices(email string) (model.UserDomainInterface, *errs.Errs)
	DeleteUser(string) *errs.Errs
}