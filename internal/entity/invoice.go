package entity

// Invoice model
type Invoice struct {
	EmbeddedId
	EmbeddedName
	EmbeddedPosition
	EmbeddedArchived
	EmbeddedCreated
	EmbeddedUpdated
	Total         float64 `json:"total"`
	SubTotal      float64 `json:"subTotal"`
	VatPercentage float64 `json:"vatPercentage"`
	Description   string  `json:"description"`
}
