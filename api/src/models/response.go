package models

type ResponseLogin struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
