package model

import (
	"crypto/md5"
	"encoding/hex"

	"github.com/nathanfabio/CRUD-Go/src/configuration/errs"
)

func NewUserDomain(email, password, name string, age int8) UserDomainInterface {
	return &userDomain{email, password, name, age}
}

type userDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *errs.Errs
	UpdateUser(string, userDomain) *errs.Errs
	FindUser(string) (*userDomain, *errs.Errs)
	DeleteUser(string) *errs.Errs
}