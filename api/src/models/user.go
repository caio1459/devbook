package models

import "time"

type User struct {
	ID       uint32    `json:"user_id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Image    string    `json:"image,omitempty"`
	Register time.Time `json:"register,omitempty"`
}
