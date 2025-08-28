package valueobject

import (
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

type Schedule struct {
	WorkDays []WorkDaySchedule `json:"work_days"`
}

type WorkDaySchedule struct {
	Day       time.Weekday `json:"day"`
	StartHour int          `json:"start_hour"` // Formato 24h (9 = 9am, 13 = 1pm)
	EndHour   int          `json:"end_hour"`
	Breaks    Break        `json:"breaks,omitempty"`
}

func (s *Schedule) validateDaysWorked() error {
	if len(s.WorkDays) > 5 {
		return fmt.Errorf("el veterinario puede trabajar máximo %d días por semana, asegurando al menos un día libre", 5)
	}
	return nil
}

func (s *Schedule) validateHoursWorked() error {
	vetDaysWorked := getWeekDayMap()
	for _, workDay := range s.WorkDays {
		if isDayDuplicated(workDay.Day, vetDaysWorked) {
			return errors.New("el veterinario no puede trabajar dos veces el mismo día de la semana")
		}

		if !s.isValidWorkDay(workDay) {
			return fmt.Errorf("el veterinario solo puede trabajar en el horario de la clínica (9:00 AM - 8:00 PM)")
		}

		setDaysAsWorked(workDay.Day, vetDaysWorked)
	}
	return nil
}

func (s *Schedule) isValidWorkDay(workDay WorkDaySchedule) bool {
	// Validar horario laboral principal
	if !s.isHoursWithinServiceSchedule(workDay.StartHour, workDay.EndHour) {
		return false
	}

	// Validar descansos
	if !s.isValidBreak(workDay.Breaks, workDay) {
		return false
	}

	return true
}

func (s *Schedule) Validate() error {
	if err := s.validateDaysWorked(); err != nil {
		return err
	}

	if err := s.validateHoursWorked(); err != nil {
		return err
	}

	return nil
}

func (s *Schedule) isValidBreak(brk Break, workDay WorkDaySchedule) bool {
	// Validar que el descanso esté dentro del horario de la clínica
	if !s.isHoursWithinServiceSchedule(brk.StartHour, brk.EndHour) {
		return false
	}

	// Validar que el descanso esté dentro del horario laboral del día
	if brk.StartHour < workDay.StartHour || brk.EndHour > workDay.EndHour {
		return false
	}

	// Validar duración máxima del descanso
	if (brk.EndHour - brk.StartHour) > 2 {
		return false
	}

	return true
}

func (s *Schedule) isHoursWithinServiceSchedule(startHour, endHour int) bool {
	if startHour < 9 || endHour > 19 {
		return false
	}

	if startHour >= endHour {
		return false
	}
	return true
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

type VetId struct {
	intId shared.IntegerId
}

func NewVeterinarianId(value any) (VetId, error) {
	id, err := shared.NewIntegerId(value)
	if err != nil {
		return VetId{}, fmt.Errorf("invalid VetId: %w", err)
	}

	return VetId{intId: id}, nil
}

func (v VetId) GetValue() int {
	return v.intId.GetValue()
}

func (v VetId) String() string {
	return v.intId.String()
}

func (v VetId) Equals(other VetId) bool {
	return v.intId.Equals(other.intId)
}

func (v VetId) IsValid() bool {
	return v.intId.GetValue() > 0
}

func (v VetId) IsZero() bool {
	return v.intId.GetValue() == 0
}

func (v VetId) Validate() error {
	if !v.IsValid() {
		return fmt.Errorf("invalid VetId: %d", v.GetValue())
	}
	return nil
}

type VetID struct {
	id shared.IntegerId
}

func NewVetID(value any) (VetID, error) {
	id, err := shared.NewIntegerId(value)
	if err != nil {
		return VetID{}, err
	}

	return VetID{id: id}, nil
}

func (p VetID) GetValue() int {
	return p.id.GetValue()
}

func (p VetID) String() string {
	return p.id.String()
}

func (p VetID) Equals(other VetID) bool {
	return p.id.GetValue() == other.id.GetValue()
}
