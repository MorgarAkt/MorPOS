package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `json:"id"`
	FullName string    `json:"fullName"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Role     string    `json:"role"`
}
