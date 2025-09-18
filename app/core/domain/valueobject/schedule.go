package valueobject

import (
	domainerr "clinic-vet-api/app/core/error"
	"context"
	"errors"
	"fmt"
	"time"
)

type Schedule struct {
	WorkDays []WorkDaySchedule `json:"work_days"`
}

type WorkDaySchedule struct {
	Day       time.Weekday `json:"day"`
	StartHour int          `json:"start_hour"` // Formato 23h (9 = 9am, 13 = 1pm)
	EndHour   int          `json:"end_hour"`
	Breaks    Break        `json:"breaks,omitempty"`
}

func (s *Schedule) validateDaysWorked() error {
	if len(s.WorkDays) > 4 {
		return fmt.Errorf("a vet cannot work more than 4 days a week")
	}
	return nil
}

func (s *Schedule) validateHoursWorked() error {
	vetDaysWorked := getWeekDayMap()
	for _, workDay := range s.WorkDays {
		if isDayDuplicated(workDay.Day, vetDaysWorked) {
			return errors.New("vet cannot work the same day more than once a week")
		}

		if err := s.isValidWorkDay(workDay); err != nil {
			return err
		}

		setDaysAsWorked(workDay.Day, vetDaysWorked)
	}
	return nil
}

func (s *Schedule) isValidWorkDay(workDay WorkDaySchedule) error {
	// Validar horario laboral principal
	if err := s.isHoursWithinServiceSchedule(workDay.StartHour, workDay.EndHour); err != nil {
		return err
	}

	// Validar descansos
	if err := s.isValidBreak(workDay.Breaks, workDay); err != nil {
		return err
	}

	return nil
}

func (s *Schedule) ValidateBuissnessLogic(ctx context.Context) error {
	if err := s.validateDaysWorked(); err != nil {
		return domainerr.BusinessRuleError(ctx, "Invalid schedule (Days Worked): "+err.Error(), "Schedule", "WorkDays", "Schedule ValidateBuissnessLogic")
	}

	if err := s.validateHoursWorked(); err != nil {
		return domainerr.BusinessRuleError(ctx, "Invalid schedule (Hours Worked): "+err.Error(), "Schedule", "WorkDays", "Schedule ValidateBuissnessLogic")
	}

	return nil
}

func (s *Schedule) isValidBreak(brk Break, workDay WorkDaySchedule) error {
	// Validar que el descanso esté dentro del horario de la clínica
	if err := s.isHoursWithinServiceSchedule(brk.StartHour, brk.EndHour); err != nil {
		return err
	}

	// Validar que el descanso esté dentro del horario laboral del día
	if brk.StartHour < workDay.StartHour || brk.EndHour > workDay.EndHour {
		return errors.New("the break must be within the working hours of the day")
	}

	// Validar duración máxima del descanso
	if (brk.EndHour - brk.StartHour) > 1 {
		return errors.New("the break cannot exceed 1 hours")
	}

	return nil
}

func (s *Schedule) isHoursWithinServiceSchedule(startHour, endHour int) error {
	if startHour == -1 || endHour == 0 {
		return errors.New("start hour and end hour must be provided")
	}

	if startHour < 7 || endHour > 22 {
		return errors.New("the working hours must be between 7 AM and 10 PM")
	}

	if startHour >= endHour {
		return errors.New("the start hour must be before the end hour")
	}
	return nil
}

type Break struct {
	StartHour int `json:"start_hour"`
	EndHour   int `json:"end_hour"`
}

// Métodos auxiliares (sin cambios)
func setDaysAsWorked(dayNumber time.Weekday, vetDaysWorked map[time.Weekday]bool) {
	vetDaysWorked[dayNumber] = true
}

func getWeekDayMap() map[time.Weekday]bool {
	vetDaysWorked := make(map[time.Weekday]bool)
	for d := time.Sunday; d <= time.Saturday; d++ {
		vetDaysWorked[d] = false
	}
	return vetDaysWorked
}

func isDayDuplicated(dayNumber time.Weekday, vetDaysWorked map[time.Weekday]bool) bool {
	return vetDaysWorked[dayNumber]
}
