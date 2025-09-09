package persistence

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/veterinarian"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
	"github.com/jackc/pgx/v5/pgtype"
)

func SqlcVetToDomain(sql sqlc.Veterinarian) (*veterinarian.Veterinarian, error) {
	if sql.FirstName == "" || sql.LastName == "" {
		return nil, errors.New("first name and last name are required")
	}

	if sql.LicenseNumber == "" {
		return nil, errors.New("license number is required")
	}

	if sql.Speciality == "" {
		return nil, errors.New("specialty is required")
	}

	vetID, err := valueobject.NewVetID(int(sql.ID))
	if err != nil {
		return nil, fmt.Errorf("invalid vet ID: %w", err)
	}

	// Mapeo del nombre
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
		schedule, err = UnmarshalVetSchedule(sql.ScheduleJson)
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
		userIDVal, err := valueobject.NewUserID(int(sql.UserID.Int32))
		if err != nil {
			return nil, fmt.Errorf("invalid user ID: %w", err)
		}
		userID = &userIDVal
	}

	// Crear options
	opts := []veterinarian.VeterinarianOption{
		veterinarian.WithName(name),
		veterinarian.WithPhoto(sql.Photo),
		veterinarian.WithLicenseNumber(sql.LicenseNumber),
		veterinarian.WithSpecialty(specialty),
		veterinarian.WithYearsExperience(int(sql.YearsOfExperience)),
		veterinarian.WithSchedule(schedule),
		veterinarian.WithIsActive(getBoolFromNullBool(sql.IsActive, true)),
		veterinarian.WithUserID(userID),
		veterinarian.WithTimestamps(sql.CreatedAt.Time, sql.UpdatedAt.Time),
	}

	// Crear la entidad
	vet, err := veterinarian.NewVeterinarian(vetID, opts...)
	if err != nil {
		return nil, fmt.Errorf("failed to create veterinarian from database: %w", err)
	}

	return vet, nil
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

func UnmarshalVetSchedule(sqlJSON []byte) (*valueobject.Schedule, error) {
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
