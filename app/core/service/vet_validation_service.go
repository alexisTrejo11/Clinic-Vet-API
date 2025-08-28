package service

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
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

type VetValidatorService struct {
	vet *entity.Veterinarian
}

func (v *VetValidatorService) ValidateVetCreation() error {
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

func (v *VetValidatorService) validateLicenseNumber() error {
	if len(v.vet.GetLicenseNumber()) < MIN_LICENSE_LENGTH || len(v.vet.GetLicenseNumber()) > MAX_LICENSE_LENGTH {
		return fmt.Errorf("veterinarian license number invalid length")
	}
	return nil
}

func (v *VetValidatorService) validateYearsOfExperience() error {
	if v.vet.GetYearsExperience() > MAX_EXPERIENCE_YEARS {
		return fmt.Errorf("years of experience seems unrealistic for a human career span")
	}
	return nil
}

func (v *VetValidatorService) validateSchedule() error {
	if v.vet.GetScheduleJSON() == "" {
		return nil
	}

	if err := json.Unmarshal([]byte(v.vet.GetScheduleJSON()), v.vet.GetSchedule()); err != nil {
		return fmt.Errorf("invalid schedule format: %v", err)
	}

	return v.vet.GetSchedule().Validate()
}

func (v *VetValidatorService) BeforeSave() error {
	scheduleBytes, err := json.Marshal(v.vet.GetSchedule())
	if err != nil {
		return err
	}
	v.vet.SetScheduleJSON(string(scheduleBytes))
	return nil
}

func (v *VetValidatorService) AfterFind() error {
	if v.vet.GetScheduleJSON() != "" {
		return json.Unmarshal([]byte(v.vet.GetScheduleJSON()), v.vet.GetSchedule())
	}
	return nil
}
