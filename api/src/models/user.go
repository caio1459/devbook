package models

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID       uint64    `json:"user_id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	Register time.Time `json:"register,omitempty"`
}

func (user *User) validate(stage string) error {
	if user.Password == "" && stage == "cadastro" {
		return errors.New("é necessário preencher a senha")
	}

	switch "" {
	case user.Name:
		return errors.New("é necessário preencher o nome")
	case user.Nick:
		return errors.New("é necessário preencher o Nick de usuário")
	case user.Email:
		return errors.New("é necessário preencher o email")
	default:
		return nil
	}
}

func (user *User) formatStrings() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}

// Chama os métodos de valitação e formatação
func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}
	user.formatStrings()
	return nil
}
