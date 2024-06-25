package models

import (
	"github.com/google/uuid"
)

type Salon struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Tables []Table   `json:"tables"`
}
