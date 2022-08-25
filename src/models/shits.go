package models

type Shit struct {
	Id      int `json:"-"`
	GroupId int `json:"group_id"`
	UserId  int `json:"user_performed"`
}
