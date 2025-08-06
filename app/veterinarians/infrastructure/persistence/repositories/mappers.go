package sqlcVetRepo

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
	vetDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/veterinarians/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/sqlc"
)

func SqlcVetToDomain(sql sqlc.Veterinarian) *vetDomain.Veterinarian {
	name, _ := shared.NewPersonName(sql.FirstName, sql.LastName)

	var isActive bool
	if sql.IsActive.Valid {
		isActive = sql.IsActive.Bool
	}

	var userID *int
	if sql.UserID.Valid {
		uid := int(sql.UserID.Int32)
		userID = &uid
	}

	var createdAt, updatedAt time.Time
	if sql.CreatedAt.Valid {
		createdAt = sql.CreatedAt.Time
	}
	if sql.UpdatedAt.Valid {
		updatedAt = sql.UpdatedAt.Time
	}

	scheduleJSON, err := UnmarshalVetSchedule(sql.ScheduleJson)
	if err != nil {
		fmt.Println(err.Error())
		scheduleJSON = &vetDomain.Schedule{}
	}

	return &vetDomain.Veterinarian{
		ID:              int(sql.ID),
		Name:            name,
		Photo:           sql.Photo,
		LicenseNumber:   sql.LicenseNumber,
		Specialty:       vetDomain.VetSpecialtyFromString(shared.AssertString(sql.Speciality)),
		YearsExperience: int(sql.YearsOfExperience),
		ConsultationFee: nil,
		IsActive:        isActive,
		UserID:          userID,
		CreatedAt:       createdAt,
		UpdatedAt:       updatedAt,
		Schedule:        scheduleJSON,
		ScheduleJSON:    "{}",
	}
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
