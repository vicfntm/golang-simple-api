package services

import (
	"github.com/vicfntm/splitshit2/src/models"
	"github.com/vicfntm/splitshit2/src/repositories"
)

type Authorization interface {
	CreateCustomer(customer models.Customer) (int, error)
	LoginUser(user models.Customer) (string, error)
	ParseToken(token string) (int, error)
	FindCustomer(id interface{}) (models.Customer, error)
	AssignRole(id int, role string) (models.RoleCustomer, error)
}

type PublicApi interface {
}

type Services struct {
	Authorization
	PublicApi
}

func NewService(repo *repositories.Repositories) *Services {
	return &Services{
		Authorization: NewAuthService(repo),
	}
}
