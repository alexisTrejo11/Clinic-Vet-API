package repository

import (
	"clinic-vet-api/app/modules/core/domain/entity/employee"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/db/models"
	"clinic-vet-api/sqlc"
	"encoding/json"
	"fmt"
	"log"
	"time"
)

type postgresSchedule struct {
	Monday    *postgresDaySchedule `json:"monday,omitempty"`
	Tuesday   *postgresDaySchedule `json:"tuesday,omitempty"`
	Wednesday *postgresDaySchedule `json:"wednesday,omitempty"`
	Thursday  *postgresDaySchedule `json:"thursday,omitempty"`
	Friday    *postgresDaySchedule `json:"friday,omitempty"`
	Saturday  *postgresDaySchedule `json:"saturday,omitempty"`
	Sunday    *postgresDaySchedule `json:"sunday,omitempty"`
}

type postgresDaySchedule struct {
	Start string `json:"start"`
	End   string `json:"end"`
	Break string `json:"break,omitempty"`
}

func (r *SqlcEmployeeRepository) toEntity(sql sqlc.Employee) *employee.Employee {
	var schedule *valueobject.Schedule
	var err error
	if sql.ScheduleJson != nil {
		schedule, err = UnmarshalEmployeeSchedule(sql.ScheduleJson)
		if err != nil {
			log.Printf("Warning: Failed to unmarshal schedule: %v", err)
			schedule = &valueobject.Schedule{}
		}
	} else {
		schedule = &valueobject.Schedule{}
	}

	employeeID := valueobject.NewEmployeeID(uint(sql.ID))
	personName := valueobject.PersonName{FirstName: sql.FirstName, LastName: sql.LastName}
	userID := r.pgMap.PgInt4.ToUserIDPtr(sql.UserID)
	return employee.NewEmployeeBuilder().
		WithID(employeeID).
		WithName(personName).
		WithPhoto(sql.Photo).
		WithLicenseNumber(sql.LicenseNumber).
		WithSpecialty(enum.VetSpecialty(string(sql.Speciality))).
		WithYearsExperience(sql.YearsOfExperience).
		WithSchedule(schedule).
		WithIsActive(sql.IsActive).
		WithUserID(userID).
		WithTimestamps(sql.CreatedAt.Time, sql.UpdatedAt.Time).
		Build()
}

func (r *SqlcEmployeeRepository) toEntities(rows []sqlc.Employee) []employee.Employee {
	if len(rows) == 0 {
		return []employee.Employee{}
	}

	employees := make([]employee.Employee, len(rows))
	for _, row := range rows {
		emp := r.toEntity(row)
		employees = append(employees, *emp)
	}
	return employees
}

func parseScheduleFromPostgres(jsonData []byte) (*valueobject.Schedule, error) {
	var pgSchedule postgresSchedule
	if err := json.Unmarshal(jsonData, &pgSchedule); err != nil {
		return nil, fmt.Errorf("error al parsear JSON de PostgreSQL: %v", err)
	}

	schedule := &valueobject.Schedule{WorkDays: make([]valueobject.WorkDaySchedule, 0)}

	if pgSchedule.Monday != nil {
		schedule.WorkDays = append(schedule.WorkDays, parseDaySchedule(time.Monday, pgSchedule.Monday))
	}

	if pgSchedule.Tuesday != nil {
		schedule.WorkDays = append(schedule.WorkDays, parseDaySchedule(time.Tuesday, pgSchedule.Tuesday))
	}

	if pgSchedule.Wednesday != nil {
		schedule.WorkDays = append(schedule.WorkDays, parseDaySchedule(time.Wednesday, pgSchedule.Wednesday))
	}
	if pgSchedule.Thursday != nil {
		schedule.WorkDays = append(schedule.WorkDays, parseDaySchedule(time.Thursday, pgSchedule.Thursday))
	}
	if pgSchedule.Friday != nil {
		schedule.WorkDays = append(schedule.WorkDays, parseDaySchedule(time.Friday, pgSchedule.Friday))
	}
	if pgSchedule.Saturday != nil {
		schedule.WorkDays = append(schedule.WorkDays, parseDaySchedule(time.Saturday, pgSchedule.Saturday))
	}
	if pgSchedule.Sunday != nil {
		schedule.WorkDays = append(schedule.WorkDays, parseDaySchedule(time.Sunday, pgSchedule.Sunday))
	}

	return schedule, nil
}

func parseDaySchedule(day time.Weekday, pgDay *postgresDaySchedule) valueobject.WorkDaySchedule {
	startHour := parseHourToInt(pgDay.Start)
	endHour := parseHourToInt(pgDay.End)

	workDay := valueobject.WorkDaySchedule{
		Day:       day,
		StartHour: startHour,
		EndHour:   endHour,
	}

	if pgDay.Break != "" {
		breakParts := parseBreak(pgDay.Break)
		if breakParts != nil {
			workDay.Breaks = *breakParts
		}
	}

	return workDay
}

func parseHourToInt(timeStr string) int {
	var h, m int
	fmt.Sscanf(timeStr, "%d:%d", &h, &m)
	return h
}

func parseBreak(breakStr string) *valueobject.Break {
	var startH, startM, endH, endM int
	_, err := fmt.Sscanf(breakStr, "%d:%d-%d:%d", &startH, &startM, &endH, &endM)
	if err != nil {
		return nil
	}
	return &valueobject.Break{
		StartHour: startH,
		EndHour:   endH,
	}
}

func UnmarshalEmployeeSchedule(sqlJSON []byte) (*valueobject.Schedule, error) {
	if sqlJSON == nil {
		return &valueobject.Schedule{}, nil
	}

	return parseScheduleFromPostgres(sqlJSON)
}

func (r *SqlcEmployeeRepository) toUpdateParams(employee *employee.Employee) *sqlc.UpdateEmployeeParams {
	return &sqlc.UpdateEmployeeParams{
		ID:                employee.ID().Int32(),
		FirstName:         employee.Name().FirstName,
		LastName:          employee.Name().LastName,
		LicenseNumber:     employee.LicenseNumber(),
		Photo:             employee.Photo(),
		Speciality:        models.VeterinarianSpeciality(employee.Specialty().String()),
		YearsOfExperience: employee.YearsExperience(),
		IsActive:          employee.IsActive(),
		UserID:            r.pgMap.PgInt4.FromUserIDPtr(employee.UserID()),
	}
}

func (r *SqlcEmployeeRepository) toCreateParams(employee *employee.Employee) *sqlc.CreateEmployeeParams {
	return &sqlc.CreateEmployeeParams{
		FirstName:         employee.Name().FirstName,
		LastName:          employee.Name().LastName,
		LicenseNumber:     employee.LicenseNumber(),
		Photo:             employee.Photo(),
		Speciality:        models.VeterinarianSpeciality(employee.Specialty().String()),
		YearsOfExperience: employee.YearsExperience(),
		IsActive:          employee.IsActive(),
		UserID:            r.pgMap.PgInt4.FromUserIDPtr(employee.UserID()),
	}
}
