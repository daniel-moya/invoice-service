package entity

import "time"

// Invoice model
type Invoice struct {
	Id            int       `json:"id"`
	Position      int       `json:"position"`
	Name          string    `json:"name"`
	Archived      bool      `json:"archived"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
	Total         float64   `json:"total"`
	SubTotal      float64   `json:"subTotal"`
	VatPercentage float64   `json:"vatPercentage"`
	Description   string    `json:"description"`
}
