package models

import (
	"github.com/google/uuid"
)

type Table struct {
	ID        uuid.UUID `json:"id"`
	SalonID   uuid.UUID `json:"salonID"`
	Number    int       `json:"number"`
	Products  []Product `json:"products"`
	TotalBill float64   `json:"totalBill"`
}
