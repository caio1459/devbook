package models

type Response struct {
	User  User   `json:"user"`
	Token string `json:"token"`
}
