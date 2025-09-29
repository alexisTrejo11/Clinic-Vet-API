package service

import (
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"errors"
	"time"
)

type VaccineCatalog struct {
	vaccines map[string]valueobject.VaccineDefinition
}

func NewVaccineCatalog() *VaccineCatalog {
	catalog := &VaccineCatalog{
		vaccines: make(map[string]valueobject.VaccineDefinition),
	}
	catalog.initializeDefaultVaccines()
	return catalog
}

func (vc *VaccineCatalog) initializeDefaultVaccines() {
	// Dog vaccines
	dhpp := vc.createDHPPVaccine()
	rabies := vc.createRabiesVaccine()
	leptospirosis := vc.createLeptospirosisVaccine()
	bordetella := vc.createBordetellaVaccine()

	// Cat vaccines
	fvrcp := vc.createFVRCPVaccine()
	felv := vc.createFeLVVaccine()
	rabiesCat := vc.createRabiesVaccineCat()

	vc.vaccines[dhpp.Name().Value()] = dhpp
	vc.vaccines[rabies.Name().Value()] = rabies
	vc.vaccines[leptospirosis.Name().Value()] = leptospirosis
	vc.vaccines[bordetella.Name().Value()] = bordetella
	vc.vaccines[fvrcp.Name().Value()] = fvrcp
	vc.vaccines[felv.Name().Value()] = felv
	vc.vaccines[rabiesCat.Name().Value()] = rabiesCat
}

// Dog vaccines
func (vc *VaccineCatalog) createDHPPVaccine() valueobject.VaccineDefinition {
	name, _ := valueobject.ParseVaccineName("DHPP")
	schedule := valueobject.NewVaccineSchedule(
		3,                // 3 inital doses
		21*24*time.Hour,  // every 3 weeks
		365*24*time.Hour, // annual booster
		42*24*time.Hour,  // starting at 6 weeks
	)

	return valueobject.NewVaccineDefinition(
		name,
		valueobject.VaccineTypeCore,
		[]string{"dog"},
		[]string{"Distemper", "Hepatitis", "Parvovirus", "Parainfluenza"},
		schedule,
	).WithDescription("Essential multivalent vaccine for dogs.").
		WithSideEffects([]string{"Mild lethargy", "Injection site inflammation", "Mild fever"}).
		WithContraindications([]string{"Active disease", "Severe immunosuppression"})
}

func (vc *VaccineCatalog) createRabiesVaccine() valueobject.VaccineDefinition {
	name, _ := valueobject.ParseVaccineName("Rabies")
	schedule := valueobject.NewVaccineSchedule(
		1,                // 1 initial dose
		0,                // no interval between initial doses
		365*24*time.Hour, // annual booster
		84*24*time.Hour,  // starting at 12 weeks
	).WithLifelong(false)

	return valueobject.NewVaccineDefinition(
		name,
		valueobject.VaccineTypeLegallyReq,
		[]string{"dog"},
		[]string{"Rabies"},
		schedule,
	).WithDescription("Vacuna antirrábica legalmente requerida").
		WithContraindications([]string{"Hipersensibilidad previa"})
}

func (vc *VaccineCatalog) createLeptospirosisVaccine() valueobject.VaccineDefinition {
	name, _ := valueobject.ParseVaccineName("Leptospirosis")
	schedule := valueobject.NewVaccineSchedule(
		2,                // 2 initial doses
		21*24*time.Hour,  // every 3 weeks
		365*24*time.Hour, // annual booster
		84*24*time.Hour,  // starting at 12 weeks
	)

	return valueobject.NewVaccineDefinition(
		name,
		valueobject.VaccineTypeCore,
		[]string{"dog"},
		[]string{"Leptospirosis"},
		schedule,
	).WithDescription("Prevents leptospirosis in dogs.").
		WithSideEffects([]string{"Possible allergic reactions", "Vomiting", "Diarrhea"})
}

func (vc *VaccineCatalog) createBordetellaVaccine() valueobject.VaccineDefinition {
	name, _ := valueobject.ParseVaccineName("Bordetella")
	schedule := valueobject.NewVaccineSchedule(
		1, // 1 initial dose
		0,
		180*24*time.Hour, // booster every 6 months
		56*24*time.Hour,  // starting at 8 weeks
	)

	return valueobject.NewVaccineDefinition(
		name,
		valueobject.VaccineTypeNonCore,
		[]string{"dog"},
		[]string{"Kennel Cough"},
		schedule,
	).WithDescription("Prevents kennel cough in dogs.")
}

// Cat vaccines
func (vc *VaccineCatalog) createFVRCPVaccine() valueobject.VaccineDefinition {
	name, _ := valueobject.ParseVaccineName("FVRCP")
	schedule := valueobject.NewVaccineSchedule(
		3,                // 3 initial doses
		21*24*time.Hour,  // every 3 weeks
		365*24*time.Hour, // annual booster
		42*24*time.Hour,  // starting at 6 weeks
	)

	return valueobject.NewVaccineDefinition(
		name,
		valueobject.VaccineTypeCore,
		[]string{"cat"},
		[]string{"Feline Viral Rhinotracheitis", "Calicivirus", "Panleukopenia"},
		schedule,
	).WithDescription("Essential feline trivalent vaccine.")
}

func (vc *VaccineCatalog) createFeLVVaccine() valueobject.VaccineDefinition {
	name, _ := valueobject.ParseVaccineName("FeLV")
	schedule := valueobject.NewVaccineSchedule(
		2,                // 2 initial doses
		21*24*time.Hour,  // cada 3 semanas
		365*24*time.Hour, // refuerzo anual
		56*24*time.Hour,  // desde las 8 semanas
	)

	return valueobject.NewVaccineDefinition(
		name,
		valueobject.VaccineTypeNonCore,
		[]string{"cat"},
		[]string{"Feline Leukemia Virus"},
		schedule,
	).WithDescription("Previene leucemia felina")
}

func (vc *VaccineCatalog) createRabiesVaccineCat() valueobject.VaccineDefinition {
	name, _ := valueobject.ParseVaccineName("Rabies Cat")
	schedule := valueobject.NewVaccineSchedule(
		1,
		0,
		365*24*time.Hour,
		84*24*time.Hour,
	)

	return valueobject.NewVaccineDefinition(
		name,
		valueobject.VaccineTypeLegallyReq,
		[]string{"cat"},
		[]string{"Rabies"},
		schedule,
	).WithDescription("Rabies vaccine for cats.")
}

// GetVaccineByName obtiene una vacuna del catálogo
func (vc *VaccineCatalog) GetVaccineByName(name string) (valueobject.VaccineDefinition, error) {
	vaccine, exists := vc.vaccines[name]
	if !exists {
		return valueobject.VaccineDefinition{}, errors.New("vaccine not found in catalog")
	}
	return vaccine, nil
}

// GetVaccinesForSpecies obtains all vaccines applicable for a given species
func (vc *VaccineCatalog) GetVaccinesForSpecies(species string) []valueobject.VaccineDefinition {
	var result []valueobject.VaccineDefinition
	for _, vaccine := range vc.vaccines {
		if vaccine.IsApplicableForSpecies(species) {
			result = append(result, vaccine)
		}
	}
	return result
}

// GetCoreVaccinesForSpecies obtains core vaccines for a given species
func (vc *VaccineCatalog) GetCoreVaccinesForSpecies(species string) []valueobject.VaccineDefinition {
	var result []valueobject.VaccineDefinition
	for _, vaccine := range vc.vaccines {
		if vaccine.IsApplicableForSpecies(species) &&
			(vaccine.VaccineType() == valueobject.VaccineTypeCore ||
				vaccine.VaccineType() == valueobject.VaccineTypeLegallyReq) {
			result = append(result, vaccine)
		}
	}
	return result
}

// AddVaccine allows adding a new vaccine to the catalog
func (vc *VaccineCatalog) AddVaccine(vaccine valueobject.VaccineDefinition) error {
	if _, exists := vc.vaccines[vaccine.Name().Value()]; exists {
		return errors.New("vaccine already exists in catalog")
	}
	vc.vaccines[vaccine.Name().Value()] = vaccine
	return nil
}
