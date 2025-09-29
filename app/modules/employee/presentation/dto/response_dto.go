package dto

import (
	"clinic-vet-api/app/modules/employee/application/handler"
	"strconv"
	"time"
)

type EmployeeResponse struct {
	ID              string `json:"id" example:"123e4567-e89b-12d3-a456-426614174000"`
	FirstName       string `json:"first_name" example:"John"`
	LastName        string `json:"last_name" example:"Doe"`
	Photo           string `json:"photo" example:"https://example.com/photo.jpg"`
	LicenseNumber   string `json:"license_number" example:"VET123456"`
	UserID          *uint  `json:"user_id,omitempty" example:"42"`
	YearsExperience int32  `json:"years_experience" example:"5"`
	IsActive        bool   `json:"is_active" example:"true"`
	Specialty       string `json:"specialty" example:"CARDIOLOGY"`
	// LaboralSchedule []ScheduleResponse `json:"laboral_schedule"`
	CreatedAt string `json:"created_at" example:"2023-10-01T12:00:00Z"`
	UpdatedAt string `json:"updated_at" example:"2023-10-10T15:30:00Z"`
}

func (r *EmployeeResponse) FromResult(result handler.EmployeeResult) {
	r.ID = strconv.FormatUint(uint64(result.ID), 10)
	r.FirstName = result.FirstName
	r.LastName = result.LastName
	r.Photo = result.Photo
	r.LicenseNumber = result.LicenseNumber
	r.YearsExperience = result.YearsExperience
	r.IsActive = result.IsActive
	r.Specialty = result.Specialty
	r.UserID = result.UserID
	// if result.LaboralSchedule != nil {
	// 	schedules := make([]ScheduleResponse, len(*result.LaboralSchedule))
	// 	for i, sched := range *result.LaboralSchedule {
	// 		var schedResp ScheduleResponse
	// 		schedResp.FromData(sched)
	// 		schedules[i] = schedResp
	// 	}
	// 	r.LaboralSchedule = schedules
	// }
	r.CreatedAt = result.CreatedAt.Format(time.RFC3339)
	r.UpdatedAt = result.UpdatedAt.Format(time.RFC3339)
}

func ToEmployeeResponse(result handler.EmployeeResult) EmployeeResponse {
	var resp EmployeeResponse
	resp.FromResult(result)
	return resp
}

func ToEmployeeResponseList(results []handler.EmployeeResult) []EmployeeResponse {
	responses := make([]EmployeeResponse, len(results))
	for i, result := range results {
		responses[i] = ToEmployeeResponse(result)
	}
	return responses
}
