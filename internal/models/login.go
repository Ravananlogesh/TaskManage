package models

type Login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}
