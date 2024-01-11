package models

import "time"

type User struct {
	ID       uint64    `json:"user_id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Register time.Time `json:"register,omitempty"`
}
