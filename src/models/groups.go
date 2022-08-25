package models

type Groups struct {
	Id         int    `json:"-"`
	Name       string `json:"group_name"`
	Definition string `json:"definition"`
}
