package query

import (
	"clinic-vet-api/app/core/domain/entity/employee"
	"time"
)

type EmployeeResult struct {
	ID              uint
	FirstName       string
	LastName        string
	Photo           string
	LicenseNumber   string
	YearsExperience int
	Specialty       string
	IsActive        bool
	LaboralSchedule *[]ScheduleData
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type ScheduleData struct {
	Day           string
	EntryTime     int
	DepartureTime int
	StartBreak    int
	EndBreak      int
}

func ToResult(employee *employee.Employee) EmployeeResult {
	var scheduleResults *[]ScheduleData
	if employee.Schedule() != nil {
		days := employee.Schedule().WorkDays
		scheduleResultsSlice := make([]ScheduleData, len(days))
		for i, day := range days {
			schedule := ScheduleData{
				Day:           day.Day.String(),
				EntryTime:     day.StartHour,
				DepartureTime: day.EndHour,
				StartBreak:    day.Breaks.StartHour,
				EndBreak:      day.Breaks.EndHour,
			}
			scheduleResultsSlice[i] = schedule
		}
		scheduleResults = &scheduleResultsSlice
	}

	return EmployeeResult{
		ID:              employee.ID().Value(),
		FirstName:       employee.Name().FirstName,
		LastName:        employee.Name().LastName,
		Photo:           employee.Photo(),
		LicenseNumber:   employee.LicenseNumber(),
		Specialty:       employee.Specialty().String(),
		YearsExperience: employee.YearsExperience(),
		LaboralSchedule: scheduleResults,
		IsActive:        employee.IsActive(),
		CreatedAt:       employee.CreatedAt(),
		UpdatedAt:       employee.UpdatedAt(),
	}
}
