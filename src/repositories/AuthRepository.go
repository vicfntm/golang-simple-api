package repositories

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/vicfntm/splitshit2/src/models"
)

const customerTable = "customers"

type AuthRepo struct {
	dbConnect sqlx.DB
}

func NewAuthRepository(dbConn *sqlx.DB) *AuthRepo {
	return &AuthRepo{
		dbConnect: *dbConn,
	}
}

func (ar *AuthRepo) CreateUser(user models.Customer) (int, error) {
	{
		var id int
		query := fmt.Sprintf("INSERT INTO %s (login, password, created_on) values($1, $2, $3) RETURNING id", customerTable)
		row := ar.dbConnect.QueryRow(query, user.Login, user.Password, time.Now())
		if err := row.Scan(&id); err != nil {
			return 0, err
		}
		return id, nil
	}
}

func (ar *AuthRepo) GetUserFromDb(login string, hash string) (models.Customer, error) {
	var customer models.Customer
	query := fmt.Sprintf("SELECT id, login, password FROM %s WHERE login = $1 AND password = $2", customerTable)
	err := ar.dbConnect.Get(&customer, query, login, hash)

	return customer, err
}

func (ar *AuthRepo) GetCustomerById(id interface{}) (models.Customer, error) {
	var customer models.Customer
	query := fmt.Sprintf("SELECT id, login, role FROM %s WHERE id = $1 LIMIT 1", customerTable)
	err := ar.dbConnect.Get(&customer, query, id)

	return customer, err
}

func (ar *AuthRepo) AssignRole(id int, rolei string) (models.RoleCustomer, error) {
	var c models.RoleCustomer
	query := fmt.Sprintf("UPDATE %s SET role = $1 WHERE id = $2 RETURNING id, role", customerTable)
	err := ar.dbConnect.QueryRow(query, rolei, id).Scan(&c.Id, &c.Role)
	return c, err
}
