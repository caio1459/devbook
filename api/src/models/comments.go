package models

import (
	"errors"
	"strings"
	"time"
)

type Comment struct {
	ComID     uint64    `json:"com_id,omitempty"`
	PubID     uint64    `json:"pub_id,omitempty"`
	UserID    uint64    `json:"user_id,omitempty"`
	ActorNick string    `json:"actor_nick,omitempty"`
	UserImage string    `json:"user_image,omitempty"`
	Text      string    `json:"text,omitempty"`
	Register  time.Time `json:"register,omitempty"`
}

func (c *Comment) validate() error {
	if c.Text == "" {
		return errors.New("é necessário preencher o corpo do Comentario")
	}
	return nil
}

func (c *Comment) formatStrings() error {
	c.Text = strings.TrimSpace(c.Text)
	return nil
}

func (c *Comment) Prepare() error {
	if err := c.validate(); err != nil {
		return err
	}
	if err := c.formatStrings(); err != nil {
		return err
	}
	return nil
}
