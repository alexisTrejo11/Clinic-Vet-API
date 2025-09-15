package repository

import (
	"clinic-vet-api/app/core/domain/entity/employee"
	"clinic-vet-api/app/core/domain/enum"
	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/sqlc"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

func SqlcEmployeeToDomain(sql sqlc.Employee) (*employee.Employee, error) {
	if sql.FirstName == "" || sql.LastName == "" {
		return nil, errors.New("first name and last name are required")
	}

	if sql.LicenseNumber == "" {
		return nil, errors.New("license number is required")
	}

	if sql.Speciality == "" {
		return nil, errors.New("specialty is required")
	}

	employeeID := valueobject.NewEmployeeID(uint(sql.ID))

	name, err := valueobject.NewPersonName(sql.FirstName, sql.LastName)
	if err != nil {
		return nil, fmt.Errorf("error creating person name: %w", err)
	}

	// Mapeo de la especialidad
	specialty, err := enum.ParseVetSpecialty(string(sql.Speciality))
	if err != nil {
		return nil, fmt.Errorf("invalid specialty '%s': %w", sql.Speciality, err)
	}

	// Mapeo del schedule
	var schedule *valueobject.Schedule
	if sql.ScheduleJson != nil {
		schedule, err = UnmarshalEmployeeSchedule(sql.ScheduleJson)
		if err != nil {
			// Log the error but continue with empty schedule
			log.Printf("Warning: Failed to unmarshal schedule: %v", err)
			schedule = &valueobject.Schedule{}
		}
	} else {
		schedule = &valueobject.Schedule{}
	}

	// Mapeo del userID
	var userID *valueobject.UserID
	if sql.UserID.Valid {
		userIDVal := valueobject.NewUserID(uint(sql.UserID.Int32))
		userID = &userIDVal
	}

	// Crear options
	opts := []employee.EmployeeOption{
		employee.WithName(name),
		employee.WithPhoto(sql.Photo),
		employee.WithLicenseNumber(sql.LicenseNumber),
		employee.WithSpecialty(specialty),
		employee.WithYearsExperience(int(sql.YearsOfExperience)),
		employee.WithSchedule(schedule),
		employee.WithIsActive(sql.IsActive),
		employee.WithUserID(userID),
		employee.WithTimestamps(sql.CreatedAt.Time, sql.UpdatedAt.Time),
	}

	employee, err := employee.NewEmployee(employeeID, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create employee from database: %w", err)
	}

	return employee, nil
}

// Estructura temporal para parsear el JSON de PostgreSQL
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

func parseScheduleFromPostgres(jsonData []byte) (*valueobject.Schedule, error) {
	var pgSchedule postgresSchedule
	if err := json.Unmarshal(jsonData, &pgSchedule); err != nil {
		return nil, fmt.Errorf("error al parsear JSON de PostgreSQL: %v", err)
	}

	schedule := &valueobject.Schedule{WorkDays: make([]valueobject.WorkDaySchedule, 0)}

	// Mapear cada día
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

func getBoolFromNullBool(nullBool pgtype.Bool, defaultValue bool) bool {
	if nullBool.Valid {
		return nullBool.Bool
	}
	return defaultValue
}

func getTimeFromNullTime(nullTime pgtype.Timestamp, defaultValue time.Time) time.Time {
	if nullTime.Valid {
		return nullTime.Time
	}
	return defaultValue
}

func getIntFromNullInt32(nullInt sql.NullInt32, defaultValue int) int {
	if nullInt.Valid {
		return int(nullInt.Int32)
	}
	return defaultValue
}

func getStringFromNullString(nullString sql.NullString, defaultValue string) string {
	if nullString.Valid {
		return nullString.String
	}
	return defaultValue
}

func (r *SqlcEmployeeRepository) scanEmployeeFromRow(rows *sql.Rows, employee *employee.Employee) error {
	// Esta implementación depende de tu schema exacto
	// Aquí un ejemplo genérico:
	var (
		id              int32
		firstName       string
		lastName        string
		licenseNumber   string
		photo           sql.NullString
		specialty       string
		yearsExperience int32
		consultationFee sql.NullFloat64
		isActive        bool
		userID          sql.NullInt32
		scheduleJSON    sql.NullString
		createdAt       sql.NullTime
		updatedAt       sql.NullTime
	)

	err := rows.Scan(
		&id, &firstName, &lastName, &licenseNumber, &photo,
		&specialty, &yearsExperience, &consultationFee, &isActive,
		&userID, &scheduleJSON, &createdAt, &updatedAt,
	)
	if err != nil {
		return err
	}

	// TODO:
	// Aquí construirías la entidad Employee con los valores escaneados
	// Esto es solo un ejemplo - debes adaptarlo a tu implementación real

	return nil
}

func EmployeeToUpdateParams(employee *employee.Employee) *sqlc.UpdateEmployeeParams {
	return &sqlc.UpdateEmployeeParams{
		ID:                int32(employee.ID().Value()),
		FirstName:         employee.Name().FirstName,
		LastName:          employee.Name().LastName,
		LicenseNumber:     employee.LicenseNumber(),
		Photo:             employee.Photo(),
		Speciality:        enum.VetSpecialty(employee.Specialty().DisplayName()),
		YearsOfExperience: int32(employee.YearsExperience()),
		IsActive:          employee.IsActive(),
	}
}

func EmployeeToCreateParams(employee *employee.Employee) *sqlc.CreateEmployeeParams {
	return &sqlc.CreateEmployeeParams{
		FirstName:         employee.Name().FirstName,
		LastName:          employee.Name().LastName,
		LicenseNumber:     employee.LicenseNumber(),
		Photo:             employee.Photo(),
		Speciality:        enum.EmployeeSpecialty(employee.Specialty().String()),
		YearsOfExperience: int32(employee.YearsExperience()),
		IsActive:          employee.IsActive(),
	}
}
