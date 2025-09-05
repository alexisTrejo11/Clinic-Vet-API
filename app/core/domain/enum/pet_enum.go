package enum

import (
	"fmt"
	"strings"
)

// PetGender represents the gender of a pet
type PetGender string

const (
	PetGenderMale         PetGender = "male"
	PetGenderFemale       PetGender = "female"
	PetGenderNeutered     PetGender = "neutered"
	PetGenderSpayed       PetGender = "spayed"
	PetGenderUnknown      PetGender = "unknown"
	PetGenderNotSpecified PetGender = "not_specified"
)

// PetGender constants and methods
var (
	ValidPetGenders = []PetGender{
		PetGenderMale,
		PetGenderFemale,
		PetGenderNeutered,
		PetGenderSpayed,
		PetGenderUnknown,
		PetGenderNotSpecified,
	}

	petGenderMap = map[string]PetGender{
		"male":          PetGenderMale,
		"m":             PetGenderMale,
		"boy":           PetGenderMale,
		"female":        PetGenderFemale,
		"f":             PetGenderFemale,
		"girl":          PetGenderFemale,
		"neutered":      PetGenderNeutered,
		"neutered_male": PetGenderNeutered,
		"neutered male": PetGenderNeutered,
		"castrated":     PetGenderNeutered,
		"spayed":        PetGenderSpayed,
		"spayed_female": PetGenderSpayed,
		"spayed female": PetGenderSpayed,
		"unknown":       PetGenderUnknown,
		"unk":           PetGenderUnknown,
		"not_specified": PetGenderNotSpecified,
		"not specified": PetGenderNotSpecified,
		"unspecified":   PetGenderNotSpecified,
		"":              PetGenderNotSpecified,
	}

	petGenderDisplayNames = map[PetGender]string{
		PetGenderMale:         "Male",
		PetGenderFemale:       "Female",
		PetGenderNeutered:     "Neutered Male",
		PetGenderSpayed:       "Spayed Female",
		PetGenderUnknown:      "Unknown",
		PetGenderNotSpecified: "Not Specified",
	}

	petGenderScientificNames = map[PetGender]string{
		PetGenderMale:     "Male Intact",
		PetGenderFemale:   "Female Intact",
		PetGenderNeutered: "Male Neutered",
		PetGenderSpayed:   "Female Spayed",
	}

	petGenderMedicalCodes = map[PetGender]string{
		PetGenderMale:         "M",
		PetGenderFemale:       "F",
		PetGenderNeutered:     "MN",
		PetGenderSpayed:       "FS",
		PetGenderUnknown:      "U",
		PetGenderNotSpecified: "NS",
	}

	petGenderIsAltered = map[PetGender]bool{
		PetGenderMale:         false,
		PetGenderFemale:       false,
		PetGenderNeutered:     true,
		PetGenderSpayed:       true,
		PetGenderUnknown:      false,
		PetGenderNotSpecified: false,
	}

	petGenderBaseGender = map[PetGender]PetGender{
		PetGenderMale:         PetGenderMale,
		PetGenderFemale:       PetGenderFemale,
		PetGenderNeutered:     PetGenderMale,
		PetGenderSpayed:       PetGenderFemale,
		PetGenderUnknown:      PetGenderUnknown,
		PetGenderNotSpecified: PetGenderNotSpecified,
	}
)

func (pg PetGender) IsValid() bool {
	_, exists := petGenderMap[string(pg)]
	return exists
}

func ParsePetGender(gender string) (PetGender, error) {
	normalized := normalizePetGenderInput(gender)
	if val, exists := petGenderMap[normalized]; exists {
		return val, nil
	}
	return PetGenderUnknown, fmt.Errorf("invalid pet gender: %s", gender)
}

func MustParsePetGender(gender string) PetGender {
	parsed, err := ParsePetGender(gender)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (pg PetGender) String() string {
	return string(pg)
}

func (pg PetGender) DisplayName() string {
	if displayName, exists := petGenderDisplayNames[pg]; exists {
		return displayName
	}
	return "Unknown Gender"
}

func (pg PetGender) ScientificName() string {
	if scientificName, exists := petGenderScientificNames[pg]; exists {
		return scientificName
	}
	return pg.DisplayName()
}

func (pg PetGender) MedicalCode() string {
	if medicalCode, exists := petGenderMedicalCodes[pg]; exists {
		return medicalCode
	}
	return "U"
}

func (pg PetGender) Values() []PetGender {
	return ValidPetGenders
}

func (pg PetGender) IsAltered() bool {
	if isAltered, exists := petGenderIsAltered[pg]; exists {
		return isAltered
	}
	return false
}

func (pg PetGender) BaseGender() PetGender {
	if baseGender, exists := petGenderBaseGender[pg]; exists {
		return baseGender
	}
	return PetGenderUnknown
}

func (pg PetGender) IsMale() bool {
	return pg.BaseGender() == PetGenderMale
}

func (pg PetGender) IsFemale() bool {
	return pg.BaseGender() == PetGenderFemale
}

func (pg PetGender) IsIntact() bool {
	return !pg.IsAltered() && (pg.IsMale() || pg.IsFemale())
}

func (pg PetGender) RequiresSpayNeuterAlert() bool {
	// Alert for intact pets that are not specifically marked as unknown/unspecified
	return pg.IsIntact() && pg != PetGenderUnknown && pg != PetGenderNotSpecified
}

func (pg PetGender) CanBeBred() bool {
	// Only intact males and females can be bred
	return pg.IsIntact() && (pg.IsMale() || pg.IsFemale())
}

func (pg PetGender) RecommendedSurgery() PetGender {
	switch pg {
	case PetGenderMale:
		return PetGenderNeutered
	case PetGenderFemale:
		return PetGenderSpayed
	default:
		return pg
	}
}

func (pg PetGender) IsComplete() bool {
	// Returns true if gender is specified and not unknown/unspecified
	return pg != PetGenderUnknown && pg != PetGenderNotSpecified
}

// Utility functions
func normalizePetGenderInput(input string) string {
	input = strings.TrimSpace(strings.ToLower(input))
	input = strings.ReplaceAll(input, " ", "_")

	// Handle common variations
	switch input {
	case "intact_male", "intact male":
		return "male"
	case "intact_female", "intact female":
		return "female"
	case "fixed", "altered":
		// Can't determine specific altered type from these terms
		return "unknown"
	}

	return input
}

func GetAllPetGenders() []PetGender {
	return ValidPetGenders
}

func GetIntactPetGenders() []PetGender {
	return []PetGender{
		PetGenderMale,
		PetGenderFemale,
	}
}

func GetAlteredPetGenders() []PetGender {
	return []PetGender{
		PetGenderNeutered,
		PetGenderSpayed,
	}
}

func GetBasePetGenders() []PetGender {
	return []PetGender{
		PetGenderMale,
		PetGenderFemale,
		PetGenderUnknown,
		PetGenderNotSpecified,
	}
}

func GetPetGendersForMedicalRecords() []PetGender {
	return []PetGender{
		PetGenderMale,
		PetGenderFemale,
		PetGenderNeutered,
		PetGenderSpayed,
	}
}

func IsValidForBreeding(gender PetGender) bool {
	return gender.IsIntact() && gender.IsComplete()
}

func SuggestGenderFromMedicalTerm(term string) PetGender {
	term = strings.ToLower(term)
	switch {
	case strings.Contains(term, "neutered") || strings.Contains(term, "castrated"):
		return PetGenderNeutered
	case strings.Contains(term, "spayed"):
		return PetGenderSpayed
	case strings.Contains(term, "male") || strings.Contains(term, "boy"):
		return PetGenderMale
	case strings.Contains(term, "female") || strings.Contains(term, "girl"):
		return PetGenderFemale
	default:
		return PetGenderUnknown
	}
}
