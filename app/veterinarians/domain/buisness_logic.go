package vetDomain

import (
	"encoding/json"
	"errors"
	"fmt"
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

func (v *Veterinarian) ValidateBusinessLogic() error {
	var errs []error

	if err := v.validateLicenseNumber(); err != nil {
		errs = append(errs, err)
	}
	if err := v.validateYearsOfExperience(); err != nil {
		errs = append(errs, err)
	}
	if err := v.validateSchedule(); err != nil {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		return errors.Join(errs...)
	}
	return nil
}

func (v *Veterinarian) validateLicenseNumber() error {
	if len(v.licenseNumber) < MIN_LICENSE_LENGTH || len(v.licenseNumber) > MAX_LICENSE_LENGTH {
		return fmt.Errorf("veterinarian license number invalid length")
	}
	return nil
}

func (v *Veterinarian) validateYearsOfExperience() error {
	if v.yearsExperience > MAX_EXPERIENCE_YEARS {
		return fmt.Errorf("years of experience seems unrealistic for a human career span")
	}
	return nil
}

func (v *Veterinarian) validateSchedule() error {
	if v.scheduleJSON == "" {
		return nil
	}

	if err := json.Unmarshal([]byte(v.scheduleJSON), &v.schedule); err != nil {
		return fmt.Errorf("invalid schedule format: %v", err)
	}

	return v.schedule.Validate()
}

func (v *Veterinarian) BeforeSave() error {
	scheduleBytes, err := json.Marshal(v.schedule)
	if err != nil {
		return err
	}
	v.scheduleJSON = string(scheduleBytes)
	return nil
}

func (v *Veterinarian) AfterFind() error {
	if v.scheduleJSON != "" {
		return json.Unmarshal([]byte(v.scheduleJSON), &v.schedule)
	}
	return nil
}
