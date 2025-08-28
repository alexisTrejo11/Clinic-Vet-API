package sqlcVetRepo

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

func SqlcVetToDomain(sql sqlc.Veterinarian) (*vetDomain.Veterinarian, error) {
	name, err := valueObjects.NewPersonName(sql.FirstName, sql.LastName)
	if err != nil {
		return nil, fmt.Errorf("error al crear el nombre de la persona: %w", err)
	}

	schedule, err := UnmarshalVetSchedule(sql.ScheduleJson)
	if err != nil {
		schedule = &vetDomain.Schedule{}
	}

	// Utiliza el builder para construir el objeto del dominio
	builder := vetDomain.NewVeterinarianBuilder().
		WithID(int(sql.ID)).
		WithName(name).
		WithPhoto(sql.Photo).
		WithLicenseNumber(sql.LicenseNumber).
		WithYearsExperience(int(sql.YearsOfExperience)).
		WithSpecialty(vetDomain.VetSpecialtyFromString(shared.AssertString(sql.Speciality))).
		WithSchedule(schedule).
		WithScheduleJSON(string(sql.ScheduleJson))

	// Manejar campos opcionales/nulos
	if sql.IsActive.Valid {
		builder.WithIsActive(sql.IsActive.Bool)
	} else {
		// Asume un valor por defecto si no es válido
		builder.WithIsActive(false)
	}

	if sql.UserID.Valid {
		uid := int(sql.UserID.Int32)
		builder.WithUserID(&uid)
	}

	if sql.CreatedAt.Valid {
		builder.WithCreatedAt(sql.CreatedAt.Time)
	} else {
		builder.WithCreatedAt(time.Time{})
	}

	if sql.UpdatedAt.Valid {
		builder.WithUpdatedAt(sql.UpdatedAt.Time)
	} else {
		builder.WithUpdatedAt(time.Time{})
	}

	return builder.Build(), nil
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

func parseScheduleFromPostgres(jsonData []byte) (*vetDomain.Schedule, error) {
	var pgSchedule postgresSchedule
	if err := json.Unmarshal(jsonData, &pgSchedule); err != nil {
		return nil, fmt.Errorf("error al parsear JSON de PostgreSQL: %v", err)
	}

	schedule := &vetDomain.Schedule{WorkDays: make([]vetDomain.WorkDaySchedule, 0)}

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

func parseDaySchedule(day time.Weekday, pgDay *postgresDaySchedule) vetDomain.WorkDaySchedule {
	startHour := parseHourToInt(pgDay.Start)
	endHour := parseHourToInt(pgDay.End)

	workDay := vetDomain.WorkDaySchedule{
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

func parseBreak(breakStr string) *vetDomain.Break {
	var startH, startM, endH, endM int
	_, err := fmt.Sscanf(breakStr, "%d:%d-%d:%d", &startH, &startM, &endH, &endM)
	if err != nil {
		return nil
	}
	return &vetDomain.Break{
		StartHour: startH,
		EndHour:   endH,
	}
}

func UnmarshalVetSchedule(sqlJSON []byte) (*vetDomain.Schedule, error) {
	if sqlJSON == nil {
		return &vetDomain.Schedule{}, nil
	}

	return parseScheduleFromPostgres(sqlJSON)
}
