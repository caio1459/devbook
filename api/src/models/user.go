package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/caio1459/devbook/src/security"
)

type User struct {
	ID       uint64    `json:"user_id,omitempty"`
	Name     string    `json:"name,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Password string    `json:"password,omitempty"`
	ImageUrl string    `json:"image_url,omitempty"`
	Register time.Time `json:"register,omitempty"`
}

func (user *User) validate(stage string) error {
	if user.Password == "" && stage == "cadastro" {
		return errors.New("é necessário preencher a senha")
	}
	if user.Email == "" {
		return errors.New("é necessário preencher o nome")
	}
	if user.Nick == "" {
		return errors.New("é necessário preencher o Nick de usuário")
	}
	if user.Email == "" {
		return errors.New("é necessário preencher o email")
	}
	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("e-mail inválido")
	}
	return nil
}

// Fução responsavel por formatar strings e cryptografar senhas
func (user *User) formatStrings(stage string) error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
	user.ImageUrl = strings.TrimSpace(user.ImageUrl)

	if stage == "cadastro" {
		hashedPassword, err := security.Hash(user.Password)
		if err != nil {
			return err
		}
		user.Password = string(hashedPassword)
	}

	return nil
}

// Chama os métodos de valitação e formatação
func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}
	if err := user.formatStrings(stage); err != nil {
		return err
	}
	return nil
}
