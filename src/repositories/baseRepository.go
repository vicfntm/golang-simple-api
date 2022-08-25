package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/vicfntm/splitshit2/src/models"
)

type Authorization interface {
	CreateUser(user models.Customer) (int, error)
	GetUserFromDb(login string, hash string) (models.Customer, error)
	GetCustomerById(id interface{}) (models.Customer, error)
	AssignRole(id int, role string) (models.RoleCustomer, error)
}

type PublicApi interface {
}

type Repositories struct {
	Authorization
	PublicApi
}

func NewRepository(dbConn *sqlx.DB) *Repositories {
	return &Repositories{
		Authorization: NewAuthRepository(dbConn),
		PublicApi:     NewPublicRepo(dbConn),
	}
}
