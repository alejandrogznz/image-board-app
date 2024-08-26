package model

type User struct {
	id int `json:"id"`
	LastName string `json:"LastName"`
	FirstName string `json:"FirstName"`
}
