package dto

type CreateApptRequest struct {
	CustomerID uint    `json:"customer_id" binding:"required"`
	PetID      uint    `json:"pet_id" binding:"required"`
	VetID      *uint   `json:"vet_id"`
	Date       string  `json:"date" binding:"required,datetime=2006-01-02T15:04:05Z07:00"`
	Reason     string  `json:"reason" binding:"required"`
	Service    string  `json:"service" binding:"required"`
	Notes      *string `json:"notes"`
	Status     string  `json:"status" binding:"required,oneof=scheduled completed cancelled no_show"`
}
