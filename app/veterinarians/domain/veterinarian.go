package vetDomain

import (
	"errors"
	"fmt"
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared"
)

const (
	CLINIC_OPENING_HOUR  = 9
	CLINIC_CLOSING_HOUR  = 20
	TOTAL_WEEK_DAYS      = 7
	MAX_WORK_DAYS        = TOTAL_WEEK_DAYS - 1
	MIN_LICENSE_LENGTH   = 8
	MAX_LICENSE_LENGTH   = 20
	MAX_EXPERIENCE_YEARS = 60
	MAX_BREAK_HOURS      = 2
)

type Veterinarian struct {
	ID               uint
	Name             shared.PersonName
	Photo            string
	LicenseNumber    string
	Specialty        VetSpecialty
	YearsExperience  uint
	ConsultationFee  *shared.Money
	IsActive         bool
	UserID           *uint
	WorkDaysSchedule []WorkDaySchedule
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

func (v *Veterinarian) ValidateBuissnessLogic() error {
	var errs []error

	if err := v.validateLicenseNumber(); err != nil {
		errs = append(errs, err)

	}
	if err := v.validateYearsOfExperience(); err != nil {
		errs = append(errs, err)

	}
	if err := v.validateWorkDaysSchedule(); err != nil {
		errs = append(errs, err)

	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}

	return nil
}

func (v *Veterinarian) validateLicenseNumber() error {
	if len(v.LicenseNumber) < MIN_LICENSE_LENGTH || len(v.LicenseNumber) > MAX_LICENSE_LENGTH {
		return fmt.Errorf("veterinarian license number invalid length")
	}
	return nil
}

func (v *Veterinarian) validateYearsOfExperience() error {
	if v.YearsExperience > MAX_EXPERIENCE_YEARS {
		return fmt.Errorf("years of experience seems unrealistic for a human career span")
	}
	return nil
}

// Schedule Move??
func (v *Veterinarian) validateWorkDaysSchedule() error {
	if err := v.validateDaysWorked(); err != nil {
		return err
	}

	if err := v.validateHoursWorked(); err != nil {
		return err
	}

	return nil
}

func (v *Veterinarian) validateDaysWorked() error {
	if len(v.WorkDaysSchedule) > MAX_WORK_DAYS {
		return fmt.Errorf("veterinarian can work a maximum of %d days per week, ensuring at least one day off", MAX_WORK_DAYS)
	}
	return nil
}

func (v *Veterinarian) validateHoursWorked() error {
	vetDaysWorked := getWeekDayMap()
	for _, workDay := range v.WorkDaysSchedule {
		if isDayDuplicated(workDay.Day, vetDaysWorked) {
			return fmt.Errorf("vet can't work twice in the same week day")
		} else {
			if isValid := v.validateWorkdayHour(workDay); !isValid {
				return fmt.Errorf("vet can only have work in clinic operating schedules. Clinic Schedule All Days 8:00 AM - 8:00 PM")
			}

			setDaysAsWorked(workDay.Day, vetDaysWorked)
		}
	}
	return nil
}

func (v *Veterinarian) validateWorkdayHour(workDay WorkDaySchedule) bool {
	ws := workDay.WorkDayHourRange
	if isValid := v.isHoursWithinServiceSchedule(ws.StartHour, ws.EndHour); !isValid {
		return false
	}

	if isValid := v.validateBreakHours(workDay); !isValid {
		return false
	}

	return true
}

func (v *Veterinarian) validateBreakHours(workDay WorkDaySchedule) bool {
	ws := workDay.WorkDayHourRange
	bs := workDay.BreakHourRange

	if bs.StartHour != 0 || bs.EndHour != 0 {
		if !v.isHoursWithinServiceSchedule(bs.StartHour, bs.EndHour) {
			return false
		}

		if bs.StartHour < ws.StartHour || bs.EndHour > ws.EndHour {
			return false
		}

		if (bs.EndHour - bs.StartHour) > MAX_BREAK_HOURS {
			return false
		}
	}
	return true
}

func (v *Veterinarian) isHoursWithinServiceSchedule(startHour, endHour int) bool {
	if startHour < CLINIC_OPENING_HOUR || endHour > CLINIC_CLOSING_HOUR {
		return false
	}

	if startHour >= endHour {
		return false
	}
	return true
}

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
