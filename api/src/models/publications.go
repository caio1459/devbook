package models

import (
	"errors"
	"strings"
	"time"
)

type Publications struct {
	PubID     uint64    `json:"pub_id,omitempty"`
	Title     string    `json:"title,omitempty"`
	Content   string    `json:"content,omitempty"`
	UserID    uint64    `json:"user_id,omitempty"`
	ActorNick string    `json:"actor_nick,omitempty"`
	ImageUrl  string    `json:"image_url,omitempty"`
	Likes     uint64    `json:"likes"`
	Register  time.Time `json:"register,omitempty"`
}

func (p *Publications) validate() error {
	if p.Title == "" {
		return errors.New("é necessário preencher o titulo da publicação")
	}
	if p.Content == "" {
		return errors.New("é necessário preencher o corpo da publicação")
	}
	return nil
}

func (p *Publications) formatStrings() error {
	p.Title = strings.TrimSpace(p.Title)
	p.Content = strings.TrimSpace(p.Content)
	p.ImageUrl = strings.TrimSpace(p.ImageUrl)
	return nil
}

func (p *Publications) Prepare() error {
	if err := p.validate(); err != nil {
		return err
	}
	if err := p.formatStrings(); err != nil {
		return err
	}
	return nil
}
