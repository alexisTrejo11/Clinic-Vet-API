package service

import (
	med "clinic-vet-api/app/modules/core/domain/entity/medical"
	"clinic-vet-api/app/modules/core/domain/entity/pet"

	"errors"
	"time"
)

type VaccinationScheduleService struct {
	catalog *VaccineCatalog
}

func NewVaccinationScheduleService(catalog *VaccineCatalog) *VaccinationScheduleService {
	return &VaccinationScheduleService{
		catalog: catalog,
	}
}

func (s *VaccinationScheduleService) ValidateVaccination(
	pet *pet.Pet,
	vaccineName string,
	administeredDate time.Time,
) error {
	vaccine, err := s.catalog.GetVaccineByName(vaccineName)
	if err != nil {
		return err
	}

	if !vaccine.IsApplicableForSpecies(string(pet.Species())) {
		return errors.New("vaccine not applicable for this pet species")
	}

	if pet.Age() != nil {
		petAgeDays := time.Duration(*pet.Age()) * 24 * time.Hour
		if petAgeDays < vaccine.Schedule().MinAgeForFirst() {
			return errors.New("pet is too young for this vaccine")
		}
	}

	return nil
}

func (s *VaccinationScheduleService) CalculateNextVaccination(
	vaccineName string,
	lastVaccinationDate time.Time,
	previousVaccinations int,
) (*time.Time, error) {
	vaccine, err := s.catalog.GetVaccineByName(vaccineName)
	if err != nil {
		return nil, err
	}

	schedule := vaccine.Schedule()

	// Si ya completó las dosis iniciales, usar intervalo de refuerzo
	var nextDate time.Time
	if previousVaccinations < schedule.InitialDoses() {
		nextDate = lastVaccinationDate.Add(schedule.IntervalBetween())
	} else {
		if schedule.IsLifelong() {
			return nil, nil
		}
		nextDate = lastVaccinationDate.Add(schedule.BoosterInterval())
	}

	return &nextDate, nil
}

func (s *VaccinationScheduleService) GetVaccinationStatus(
	pet *pet.Pet,
	vaccinations []med.PetVaccination,
) med.VaccinationStatus {
	coreVaccines := s.catalog.GetCoreVaccinesForSpecies(pet.Species().String())

	status := med.VaccinationStatus{
		PetID:            pet.ID(),
		IsUpToDate:       true,
		MissingVaccines:  []string{},
		OverdueVaccines:  []med.VaccineStatus{},
		UpcomingVaccines: []med.VaccineStatus{},
	}

	now := time.Now()
	vaccinationMap := s.groupVaccinationsByName(vaccinations)

	// Verificar cada vacuna core
	for _, vaccine := range coreVaccines {
		vaccineName := vaccine.Name().Value()
		petVaccinations := vaccinationMap[vaccineName]

		if len(petVaccinations) == 0 {
			// Falta la vacuna completamente
			status.MissingVaccines = append(status.MissingVaccines, vaccineName)
			status.IsUpToDate = false
			continue
		}

		// Obtener la última vacunación
		lastVaccination := s.getLastVaccination(petVaccinations)
		nextDueDate, err := s.CalculateNextVaccination(
			vaccineName,
			lastVaccination.AdministeredDate(),
			len(petVaccinations),
		)

		if err != nil || nextDueDate == nil {
			continue
		}

		vaccineStatus := med.VaccineStatus{
			VaccineName:      vaccineName,
			LastAdministered: lastVaccination.AdministeredDate(),
			NextDueDate:      *nextDueDate,
			DosesReceived:    len(petVaccinations),
		}

		// Verificar si está vencida
		if nextDueDate.Before(now) {
			daysOverdue := int(now.Sub(*nextDueDate).Hours() / 24)
			vaccineStatus.DaysOverdue = daysOverdue
			status.OverdueVaccines = append(status.OverdueVaccines, vaccineStatus)
			status.IsUpToDate = false
		} else if nextDueDate.Sub(now) <= 30*24*time.Hour {
			// Próxima en los siguientes 30 días
			daysUntilDue := int(nextDueDate.Sub(now).Hours() / 24)
			vaccineStatus.DaysUntilDue = daysUntilDue
			status.UpcomingVaccines = append(status.UpcomingVaccines, vaccineStatus)
		}
	}

	return status
}

func (s *VaccinationScheduleService) GenerateVaccinationPlan(
	pet *pet.Pet,
	startDate time.Time,
	existingVaccinations []med.PetVaccination,
) med.VaccinationPlan {
	plan := med.VaccinationPlan{
		PetID:           pet.ID(),
		GeneratedAt:     time.Now(),
		PlannedVaccines: []med.PlannedVaccine{},
	}

	coreVaccines := s.catalog.GetCoreVaccinesForSpecies(string(pet.Species()))
	vaccinationMap := s.groupVaccinationsByName(existingVaccinations)

	for _, vaccine := range coreVaccines {
		vaccineName := vaccine.Name().Value()
		existingDoses := vaccinationMap[vaccineName]

		// Determinar cuántas dosis faltan
		totalDosesNeeded := vaccine.Schedule().InitialDoses()
		dosesReceived := len(existingDoses)

		currentDate := startDate
		if dosesReceived > 0 {
			lastDose := s.getLastVaccination(existingDoses)
			currentDate = lastDose.AdministeredDate()
		}

		// Planificar dosis faltantes
		for i := dosesReceived; i < totalDosesNeeded; i++ {
			if i > dosesReceived {
				currentDate = currentDate.Add(vaccine.Schedule().IntervalBetween())
			}

			plannedVaccine := med.PlannedVaccine{
				VaccineName:   vaccineName,
				ScheduledDate: currentDate,
				DoseNumber:    i + 1,
				VaccineType:   vaccine.VaccineType(),
				ProtectsFrom:  vaccine.ProtectsFrom(),
				IsBooster:     false,
			}
			plan.PlannedVaccines = append(plan.PlannedVaccines, plannedVaccine)
		}

		// Planificar primer refuerzo si ya completó las dosis iniciales
		if dosesReceived >= totalDosesNeeded && !vaccine.Schedule().IsLifelong() {
			boosterDate := currentDate.Add(vaccine.Schedule().BoosterInterval())
			plannedVaccine := med.PlannedVaccine{
				VaccineName:   vaccineName,
				ScheduledDate: boosterDate,
				DoseNumber:    dosesReceived + 1,
				VaccineType:   vaccine.VaccineType(),
				ProtectsFrom:  vaccine.ProtectsFrom(),
				IsBooster:     true,
			}
			plan.PlannedVaccines = append(plan.PlannedVaccines, plannedVaccine)
		}
	}

	return plan
}

func (s *VaccinationScheduleService) CheckVaccinationConflicts(
	vaccineName string,
	administeredDate time.Time,
	recentVaccinations []med.PetVaccination,
) error {
	// No aplicar múltiples vacunas el mismo día (regla general)
	for _, vacc := range recentVaccinations {
		if vacc.AdministeredDate().Format("2006-01-02") == administeredDate.Format("2006-01-02") {
			if vacc.VaccineName() != vaccineName {
				return errors.New("another vaccine was administered on the same day")
			}
		}
	}

	// No aplicar la misma vacuna en menos de 14 días
	for _, vacc := range recentVaccinations {
		if vacc.VaccineName() == vaccineName {
			daysSince := administeredDate.Sub(vacc.AdministeredDate()).Hours() / 24
			if daysSince < 14 {
				return errors.New("minimum interval of 14 days between same vaccine doses not met")
			}
		}
	}

	return nil
}

// Helper methods

func (s *VaccinationScheduleService) groupVaccinationsByName(
	vaccinations []med.PetVaccination,
) map[string][]*med.PetVaccination {
	result := make(map[string][]*med.PetVaccination)
	for _, v := range vaccinations {
		result[v.VaccineName()] = append(result[v.VaccineName()], &v)
	}
	return result
}

func (s *VaccinationScheduleService) getLastVaccination(
	vaccinations []*med.PetVaccination,
) *med.PetVaccination {
	if len(vaccinations) == 0 {
		return nil
	}

	last := vaccinations[0]
	for _, v := range vaccinations {
		if v.AdministeredDate().After(last.AdministeredDate()) {
			last = v
		}
	}
	return last
}
