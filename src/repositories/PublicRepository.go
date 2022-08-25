package repositories

import "github.com/jmoiron/sqlx"

type Public struct {
	dbConn *sqlx.DB
}

func NewPublicRepo(conn *sqlx.DB) *Public {
	return &Public{
		dbConn: conn,
	}
}
