package models

type Customer struct {
	Id        int    `json:"-" db:"id"`
	Login     string `json:"login" db:"login" binding:"required"`
	Password  string `json:"password" db:"password" binding:"required"`
	CreatedOn string `json:"-" db:"created_on"`
	Role      string `json:"-" db:"role"`
}

type RoleCustomer struct {
	Id   int    `json:"id" db:"id"`
	Role string `json:"role" db:"role" binding:"required"`
}
