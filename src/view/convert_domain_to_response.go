package view

import (
	"github.com/nathanfabio/CRUD-Go/src/controller/model/response"
	"github.com/nathanfabio/CRUD-Go/src/model"
)

func ConvertDomainToResponse(userDomain model.UserDomainInterface) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}

}
