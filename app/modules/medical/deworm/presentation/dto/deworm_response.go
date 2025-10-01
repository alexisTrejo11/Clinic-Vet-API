package dto

import "clinic-vet-api/app/modules/medical/deworm/application/query"

// DewormResponse represents the deworming treatment record response
// @Description Detailed information about a deworming treatment administered to a pet, including medication details, administration dates, and next due date calculation based on veterinary guidelines
type DewormResponse struct {
	ID               int32   `json:"id" example:"1" description:"Unique identifier for the deworming record"`
	PetID            int32   `json:"petId" example:"5" description:"ID of the pet that received the deworming treatment"`
	AdministeredBy   int32   `json:"administeredBy" example:"3" description:"ID of the veterinarian or staff member who administered the treatment"`
	MedicationName   string  `json:"medicationName" example:"Drontal Plus" description:"Name of the deworming medication used (e.g., Drontal, Revolution, Stronghold)"`
	AdministeredDate string  `json:"administeredDate" example:"2024-01-15" description:"Date when the deworming treatment was administered (format: YYYY-MM-DD)"`
	NextDueDate      *string `json:"nextDueDate,omitempty" example:"2024-04-15" description:"Calculated next due date for deworming based on medication type and pet's age. For puppies: every 2-3 weeks until 3 months, then monthly until 6 months. For adults: every 3 months for indoor pets, every 1-2 months for outdoor pets"`
	Notes            *string `json:"notes,omitempty" example:"Administered with food. No adverse reactions observed. Pet weighed 5.2 kg." description:"Additional observations, instructions, or comments about the deworming treatment"`
	CreatedAt        string  `json:"createdAt" example:"2024-01-15 14:30:00" description:"Timestamp when this deworming record was created in the system (format: YYYY-MM-DD HH:MM:SS)"`
}

type DewormResponseMapper struct{}

func (m *DewormResponseMapper) FromResult(result query.DewormResult) DewormResponse {
	var nextDueDate *string
	if result.NextDueDate != nil {
		formattedDate := result.NextDueDate.Format("2006-01-02")
		nextDueDate = &formattedDate
	}

	return DewormResponse{
		ID:               result.ID.Int32(),
		PetID:            result.PetID.Int32(),
		AdministeredBy:   result.AdministeredBy.Int32(),
		MedicationName:   result.MedicationName,
		AdministeredDate: result.AdministeredDate.Format("2006-01-02"),
		NextDueDate:      nextDueDate,
		Notes:            result.Notes,
		CreatedAt:        result.CreatedAt.Format("2006-01-02 15:04:05"),
	}
}

func (m *DewormResponseMapper) FromResultsToResponses(results []query.DewormResult) []DewormResponse {
	responses := make([]DewormResponse, len(results))
	for i, result := range results {
		responses[i] = m.FromResult(result)
	}
	return responses
}
