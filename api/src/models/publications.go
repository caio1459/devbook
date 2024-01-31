package models

import "time"

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
