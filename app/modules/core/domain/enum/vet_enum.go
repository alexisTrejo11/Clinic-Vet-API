package enum

import (
	"slices"
)

// VetSpecialty represents the specialty of a veterinarian
type VetSpecialty string

const (
	VetSpecialtyUnknown               VetSpecialty = "unknown"
	VetSpecialtyGeneralPractice       VetSpecialty = "general_practice"
	VetSpecialtySurgery               VetSpecialty = "surgery"
	VetSpecialtyInternalMedicine      VetSpecialty = "internal_medicine"
	VetSpecialtyDentistry             VetSpecialty = "dentistry"
	VetSpecialtyDermatology           VetSpecialty = "dermatology"
	VetSpecialtyOncology              VetSpecialty = "oncology"
	VetSpecialtyCardiology            VetSpecialty = "cardiology"
	VetSpecialtyNeurology             VetSpecialty = "neurology"
	VetSpecialtyOphthalmology         VetSpecialty = "ophthalmology"
	VetSpecialtyRadiology             VetSpecialty = "radiology"
	VetSpecialtyEmergencyCriticalCare VetSpecialty = "emergency_critical_care"
	VetSpecialtyAnesthesiology        VetSpecialty = "anesthesiology"
	VetSpecialtyPathology             VetSpecialty = "pathology"
	VetSpecialtyPreventiveMedicine    VetSpecialty = "preventive_medicine"
	VetSpecialtyExoticAnimalMedicine  VetSpecialty = "exotic_animal_medicine"
	VetSpecialtyEquineMedicine        VetSpecialty = "equine_medicine"
	VetSpecialtyAvianMedicine         VetSpecialty = "avian_medicine"
	VetSpecialtyZooAnimalMedicine     VetSpecialty = "zoo_animal_medicine"
	VetSpecialtyFoodAnimalMedicine    VetSpecialty = "food_animal_medicine"
	VetSpecialtyPublicHealth          VetSpecialty = "public_health"
	VetSpecialtyBehavior              VetSpecialty = "behavior"
	VetSpecialtyRehabilitation        VetSpecialty = "rehabilitation"
	VetSpecialtyNutrition             VetSpecialty = "nutrition"
)

// VetSpecialty constants and methods
var (
	ValidVetSpecialties = []VetSpecialty{
		VetSpecialtyUnknown,
		VetSpecialtyGeneralPractice,
		VetSpecialtySurgery,
		VetSpecialtyInternalMedicine,
		VetSpecialtyDentistry,
		VetSpecialtyDermatology,
		VetSpecialtyOncology,
		VetSpecialtyCardiology,
		VetSpecialtyNeurology,
		VetSpecialtyOphthalmology,
		VetSpecialtyRadiology,
		VetSpecialtyEmergencyCriticalCare,
		VetSpecialtyAnesthesiology,
		VetSpecialtyPathology,
		VetSpecialtyPreventiveMedicine,
		VetSpecialtyExoticAnimalMedicine,
		VetSpecialtyEquineMedicine,
		VetSpecialtyAvianMedicine,
		VetSpecialtyZooAnimalMedicine,
		VetSpecialtyFoodAnimalMedicine,
		VetSpecialtyPublicHealth,
		VetSpecialtyBehavior,
		VetSpecialtyRehabilitation,
		VetSpecialtyNutrition,
	}

	vetSpecialtyMap = map[string]VetSpecialty{
		"unknown":                 VetSpecialtyUnknown,
		"general_practice":        VetSpecialtyGeneralPractice,
		"general practice":        VetSpecialtyGeneralPractice,
		"general":                 VetSpecialtyGeneralPractice,
		"surgery":                 VetSpecialtySurgery,
		"surgical":                VetSpecialtySurgery,
		"internal_medicine":       VetSpecialtyInternalMedicine,
		"internal medicine":       VetSpecialtyInternalMedicine,
		"internal":                VetSpecialtyInternalMedicine,
		"dentistry":               VetSpecialtyDentistry,
		"dental":                  VetSpecialtyDentistry,
		"dermatology":             VetSpecialtyDermatology,
		"dermatologic":            VetSpecialtyDermatology,
		"oncology":                VetSpecialtyOncology,
		"oncologic":               VetSpecialtyOncology,
		"cardiology":              VetSpecialtyCardiology,
		"cardiac":                 VetSpecialtyCardiology,
		"neurology":               VetSpecialtyNeurology,
		"neurologic":              VetSpecialtyNeurology,
		"ophthalmology":           VetSpecialtyOphthalmology,
		"eye":                     VetSpecialtyOphthalmology,
		"radiology":               VetSpecialtyRadiology,
		"radiologic":              VetSpecialtyRadiology,
		"emergency_critical_care": VetSpecialtyEmergencyCriticalCare,
		"emergency critical care": VetSpecialtyEmergencyCriticalCare,
		"emergency":               VetSpecialtyEmergencyCriticalCare,
		"critical":                VetSpecialtyEmergencyCriticalCare,
		"anesthesiology":          VetSpecialtyAnesthesiology,
		"anesthesia":              VetSpecialtyAnesthesiology,
		"pathology":               VetSpecialtyPathology,
		"pathologic":              VetSpecialtyPathology,
		"preventive_medicine":     VetSpecialtyPreventiveMedicine,
		"preventive medicine":     VetSpecialtyPreventiveMedicine,
		"preventive":              VetSpecialtyPreventiveMedicine,
		"exotic_animal_medicine":  VetSpecialtyExoticAnimalMedicine,
		"exotic animal medicine":  VetSpecialtyExoticAnimalMedicine,
		"exotic":                  VetSpecialtyExoticAnimalMedicine,
		"equine_medicine":         VetSpecialtyEquineMedicine,
		"equine medicine":         VetSpecialtyEquineMedicine,
		"equine":                  VetSpecialtyEquineMedicine,
		"horse":                   VetSpecialtyEquineMedicine,
		"avian_medicine":          VetSpecialtyAvianMedicine,
		"avian medicine":          VetSpecialtyAvianMedicine,
		"avian":                   VetSpecialtyAvianMedicine,
		"bird":                    VetSpecialtyAvianMedicine,
		"zoo_animal_medicine":     VetSpecialtyZooAnimalMedicine,
		"zoo animal medicine":     VetSpecialtyZooAnimalMedicine,
		"zoo":                     VetSpecialtyZooAnimalMedicine,
		"food_animal_medicine":    VetSpecialtyFoodAnimalMedicine,
		"food animal medicine":    VetSpecialtyFoodAnimalMedicine,
		"food":                    VetSpecialtyFoodAnimalMedicine,
		"public_health":           VetSpecialtyPublicHealth,
		"public health":           VetSpecialtyPublicHealth,
		"behavior":                VetSpecialtyBehavior,
		"behavioral":              VetSpecialtyBehavior,
		"rehabilitation":          VetSpecialtyRehabilitation,
		"rehab":                   VetSpecialtyRehabilitation,
		"physical":                VetSpecialtyRehabilitation,
		"nutrition":               VetSpecialtyNutrition,
		"nutritional":             VetSpecialtyNutrition,
		"diet":                    VetSpecialtyNutrition,
	}

	vetSpecialtyDisplayNames = map[VetSpecialty]string{
		VetSpecialtyUnknown:               "Unknown Specialty",
		VetSpecialtyGeneralPractice:       "General Practice",
		VetSpecialtySurgery:               "Surgery",
		VetSpecialtyInternalMedicine:      "Internal Medicine",
		VetSpecialtyDentistry:             "Dentistry",
		VetSpecialtyDermatology:           "Dermatology",
		VetSpecialtyOncology:              "Oncology",
		VetSpecialtyCardiology:            "Cardiology",
		VetSpecialtyNeurology:             "Neurology",
		VetSpecialtyOphthalmology:         "Ophthalmology",
		VetSpecialtyRadiology:             "Radiology",
		VetSpecialtyEmergencyCriticalCare: "Emergency & Critical Care",
		VetSpecialtyAnesthesiology:        "Anesthesiology",
		VetSpecialtyPathology:             "Pathology",
		VetSpecialtyPreventiveMedicine:    "Preventive Medicine",
		VetSpecialtyExoticAnimalMedicine:  "Exotic Animal Medicine",
		VetSpecialtyEquineMedicine:        "Equine Medicine",
		VetSpecialtyAvianMedicine:         "Avian Medicine",
		VetSpecialtyZooAnimalMedicine:     "Zoo Animal Medicine",
		VetSpecialtyFoodAnimalMedicine:    "Food Animal Medicine",
		VetSpecialtyPublicHealth:          "Public Health",
		VetSpecialtyBehavior:              "Behavioral Medicine",
		VetSpecialtyRehabilitation:        "Rehabilitation",
		VetSpecialtyNutrition:             "Nutrition",
	}

	vetSpecialtyCategories = map[VetSpecialty]string{
		VetSpecialtyGeneralPractice:       "general",
		VetSpecialtySurgery:               "surgical",
		VetSpecialtyInternalMedicine:      "medical",
		VetSpecialtyDentistry:             "surgical",
		VetSpecialtyDermatology:           "medical",
		VetSpecialtyOncology:              "medical",
		VetSpecialtyCardiology:            "medical",
		VetSpecialtyNeurology:             "medical",
		VetSpecialtyOphthalmology:         "medical",
		VetSpecialtyRadiology:             "diagnostic",
		VetSpecialtyEmergencyCriticalCare: "emergency",
		VetSpecialtyAnesthesiology:        "surgical",
		VetSpecialtyPathology:             "diagnostic",
		VetSpecialtyPreventiveMedicine:    "preventive",
		VetSpecialtyExoticAnimalMedicine:  "specialty",
		VetSpecialtyEquineMedicine:        "large_animal",
		VetSpecialtyAvianMedicine:         "specialty",
		VetSpecialtyZooAnimalMedicine:     "specialty",
		VetSpecialtyFoodAnimalMedicine:    "large_animal",
		VetSpecialtyPublicHealth:          "public_health",
		VetSpecialtyBehavior:              "behavioral",
		VetSpecialtyRehabilitation:        "therapeutic",
		VetSpecialtyNutrition:             "nutritional",
	}
)

func (vs VetSpecialty) IsValid() bool {
	_, exists := vetSpecialtyMap[string(vs)]
	return exists
}

func OptVetSpecialty(vs *string) *VetSpecialty {
	if vs == nil {
		return nil
	}
	v := VetSpecialty(*vs)
	return &v
}

func ParseVetSpecialty(specialty string) (VetSpecialty, error) {
	normalized := normalizeInput(specialty)
	if val, exists := vetSpecialtyMap[normalized]; exists {
		return val, nil
	}
	return VetSpecialtyUnknown, InvalidEnumParserError("VetSpecialty", specialty)
}

func MustParseVetSpecialty(specialty string) VetSpecialty {
	parsed, err := ParseVetSpecialty(specialty)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (vs VetSpecialty) String() string {
	return string(vs)
}

func (vs VetSpecialty) DisplayName() string {
	if displayName, exists := vetSpecialtyDisplayNames[vs]; exists {
		return displayName
	}
	return "Unknown Specialty"
}

func (vs VetSpecialty) Values() []VetSpecialty {
	return ValidVetSpecialties
}

func (vs VetSpecialty) Category() string {
	if category, exists := vetSpecialtyCategories[vs]; exists {
		return category
	}
	return "general"
}

func (vs VetSpecialty) IsSurgical() bool {
	surgicalSpecialties := []VetSpecialty{
		VetSpecialtySurgery,
		VetSpecialtyDentistry,
		VetSpecialtyAnesthesiology,
	}
	return slices.Contains(surgicalSpecialties, vs)
}

func (vs VetSpecialty) IsMedical() bool {
	medicalSpecialties := []VetSpecialty{
		VetSpecialtyInternalMedicine,
		VetSpecialtyDermatology,
		VetSpecialtyOncology,
		VetSpecialtyCardiology,
		VetSpecialtyNeurology,
		VetSpecialtyOphthalmology,
	}
	return slices.Contains(medicalSpecialties, vs)
}

func (vs VetSpecialty) IsDiagnostic() bool {
	diagnosticSpecialties := []VetSpecialty{
		VetSpecialtyRadiology,
		VetSpecialtyPathology,
	}
	return slices.Contains(diagnosticSpecialties, vs)
}

func (vs VetSpecialty) IsLargeAnimal() bool {
	largeAnimalSpecialties := []VetSpecialty{
		VetSpecialtyEquineMedicine,
		VetSpecialtyFoodAnimalMedicine,
	}
	return slices.Contains(largeAnimalSpecialties, vs)
}

func (vs VetSpecialty) IsExoticAnimal() bool {
	exoticSpecialties := []VetSpecialty{
		VetSpecialtyExoticAnimalMedicine,
		VetSpecialtyAvianMedicine,
		VetSpecialtyZooAnimalMedicine,
	}
	return slices.Contains(exoticSpecialties, vs)
}

func (vs VetSpecialty) RequiresSpecialEquipment() bool {
	return vs.IsSurgical() || vs.IsDiagnostic() || vs == VetSpecialtyDentistry
}

func GetAllVetSpecialties() []VetSpecialty {
	return ValidVetSpecialties
}

func GetSurgicalSpecialties() []VetSpecialty {
	return []VetSpecialty{
		VetSpecialtySurgery,
		VetSpecialtyDentistry,
		VetSpecialtyAnesthesiology,
	}
}

func GetMedicalSpecialties() []VetSpecialty {
	return []VetSpecialty{
		VetSpecialtyInternalMedicine,
		VetSpecialtyDermatology,
		VetSpecialtyOncology,
		VetSpecialtyCardiology,
		VetSpecialtyNeurology,
		VetSpecialtyOphthalmology,
	}
}

func GetDiagnosticSpecialties() []VetSpecialty {
	return []VetSpecialty{
		VetSpecialtyRadiology,
		VetSpecialtyPathology,
	}
}

func GetLargeAnimalSpecialties() []VetSpecialty {
	return []VetSpecialty{
		VetSpecialtyEquineMedicine,
		VetSpecialtyFoodAnimalMedicine,
	}
}

func GetExoticAnimalSpecialties() []VetSpecialty {
	return []VetSpecialty{
		VetSpecialtyExoticAnimalMedicine,
		VetSpecialtyAvianMedicine,
		VetSpecialtyZooAnimalMedicine,
	}
}
