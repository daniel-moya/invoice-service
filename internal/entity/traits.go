package entity

import "time"

type EmbeddedId struct {
	Id int `json:"id"`
}

type EmbeddedPosition struct {
	Position int `json:"position"`
}

type EmbeddedName struct {
	Name string `json:"name"`
}

type EmbeddedArchived struct {
	Archived bool `json:"archived"`
}

type EmbeddedCreated struct {
	CreatedAt *time.Time `json:"created_at"`
}

type EmbeddedUpdated struct {
	UpdatedAt *time.Time `json:"updated_at"`
}
