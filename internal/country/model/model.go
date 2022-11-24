package model

import "github.com/gofrs/uuid"

type Country struct {
	ID   uuid.UUID
	Name string `json:"name"`
}
