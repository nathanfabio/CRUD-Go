package converter

import (
	"github.com/nathanfabio/CRUD-Go/src/model"
	"github.com/nathanfabio/CRUD-Go/src/model/repository/entity"
)

func ConvertDomainToEntity(domain model.UserDomainInterface) *entity.UserEntity {
	return &entity.UserEntity{
		Email: domain.GetEmail(),
		Password: domain.GetPassword(),
		Name: domain.GetName(),
		Age: domain.GetAge(),
	}
}