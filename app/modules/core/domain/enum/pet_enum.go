package enum

import (
	"strings"
)

// PetSpecies represents the type/species of a pet
type PetSpecies string

const (
	PetSpeciesDog      PetSpecies = "dog"
	PetSpeciesCat      PetSpecies = "cat"
	PetSpeciesBird     PetSpecies = "bird"
	PetSpeciesRabbit   PetSpecies = "rabbit"
	PetSpeciesHamster  PetSpecies = "hamster"
	PetSpeciesGuinea   PetSpecies = "guinea_pig"
	PetSpeciesFerret   PetSpecies = "ferret"
	PetSpeciesReptile  PetSpecies = "reptile"
	PetSpeciesFish     PetSpecies = "fish"
	PetSpeciesHorse    PetSpecies = "horse"
	PetSpeciesFarm     PetSpecies = "farm_animal"
	PetSpeciesExotic   PetSpecies = "exotic"
	PetSpeciesWildlife PetSpecies = "wildlife"
	PetSpeciesOther    PetSpecies = "other"
	PetSpeciesUnknown  PetSpecies = "unknown"
)

// PetSpecies constants and methods
var (
	ValidPetSpeciess = []PetSpecies{
		PetSpeciesDog,
		PetSpeciesCat,
		PetSpeciesBird,
		PetSpeciesRabbit,
		PetSpeciesHamster,
		PetSpeciesGuinea,
		PetSpeciesFerret,
		PetSpeciesReptile,
		PetSpeciesFish,
		PetSpeciesHorse,
		PetSpeciesFarm,
		PetSpeciesExotic,
		PetSpeciesWildlife,
		PetSpeciesOther,
		PetSpeciesUnknown,
	}

	PetSpeciesMap = map[string]PetSpecies{
		"dog":            PetSpeciesDog,
		"canine":         PetSpeciesDog,
		"puppy":          PetSpeciesDog,
		"k9":             PetSpeciesDog,
		"dogs":           PetSpeciesDog,
		"cat":            PetSpeciesCat,
		"feline":         PetSpeciesCat,
		"kitten":         PetSpeciesCat,
		"kitty":          PetSpeciesCat,
		"cats":           PetSpeciesCat,
		"bird":           PetSpeciesBird,
		"avian":          PetSpeciesBird,
		"parrot":         PetSpeciesBird,
		"canary":         PetSpeciesBird,
		"cockatiel":      PetSpeciesBird,
		"budgie":         PetSpeciesBird,
		"parakeet":       PetSpeciesBird,
		"cockatoo":       PetSpeciesBird,
		"macaw":          PetSpeciesBird,
		"finch":          PetSpeciesBird,
		"chicken":        PetSpeciesBird,
		"duck":           PetSpeciesBird,
		"goose":          PetSpeciesBird,
		"rabbit":         PetSpeciesRabbit,
		"bunny":          PetSpeciesRabbit,
		"hare":           PetSpeciesRabbit,
		"hamster":        PetSpeciesHamster,
		"syrian":         PetSpeciesHamster,
		"dwarf hamster":  PetSpeciesHamster,
		"guinea_pig":     PetSpeciesGuinea,
		"guinea pig":     PetSpeciesGuinea,
		"cavy":           PetSpeciesGuinea,
		"ferret":         PetSpeciesFerret,
		"reptile":        PetSpeciesReptile,
		"snake":          PetSpeciesReptile,
		"lizard":         PetSpeciesReptile,
		"gecko":          PetSpeciesReptile,
		"iguana":         PetSpeciesReptile,
		"turtle":         PetSpeciesReptile,
		"tortoise":       PetSpeciesReptile,
		"chameleon":      PetSpeciesReptile,
		"bearded dragon": PetSpeciesReptile,
		"python":         PetSpeciesReptile,
		"boa":            PetSpeciesReptile,
		"fish":           PetSpeciesFish,
		"goldfish":       PetSpeciesFish,
		"betta":          PetSpeciesFish,
		"tropical":       PetSpeciesFish,
		"aquatic":        PetSpeciesFish,
		"koi":            PetSpeciesFish,
		"angelfish":      PetSpeciesFish,
		"horse":          PetSpeciesHorse,
		"equine":         PetSpeciesHorse,
		"pony":           PetSpeciesHorse,
		"mare":           PetSpeciesHorse,
		"stallion":       PetSpeciesHorse,
		"foal":           PetSpeciesHorse,
		"colt":           PetSpeciesHorse,
		"filly":          PetSpeciesHorse,
		"farm_animal":    PetSpeciesFarm,
		"farm animal":    PetSpeciesFarm,
		"livestock":      PetSpeciesFarm,
		"cow":            PetSpeciesFarm,
		"cattle":         PetSpeciesFarm,
		"pig":            PetSpeciesFarm,
		"sheep":          PetSpeciesFarm,
		"goat":           PetSpeciesFarm,
		"llama":          PetSpeciesFarm,
		"alpaca":         PetSpeciesFarm,
		"exotic":         PetSpeciesExotic,
		"exotic pet":     PetSpeciesExotic,
		"unusual":        PetSpeciesExotic,
		"wildlife":       PetSpeciesWildlife,
		"wild":           PetSpeciesWildlife,
		"rescue":         PetSpeciesWildlife,
		"other":          PetSpeciesOther,
		"mixed":          PetSpeciesOther,
		"crossbreed":     PetSpeciesOther,
		"unknown":        PetSpeciesUnknown,
		"unspecified":    PetSpeciesUnknown,
		"":               PetSpeciesUnknown,
	}

	PetSpeciesDisplayNames = map[PetSpecies]string{
		PetSpeciesDog:      "Dog",
		PetSpeciesCat:      "Cat",
		PetSpeciesBird:     "Bird",
		PetSpeciesRabbit:   "Rabbit",
		PetSpeciesHamster:  "Hamster",
		PetSpeciesGuinea:   "Guinea Pig",
		PetSpeciesFerret:   "Ferret",
		PetSpeciesReptile:  "Reptile",
		PetSpeciesFish:     "Fish",
		PetSpeciesHorse:    "Horse",
		PetSpeciesFarm:     "Farm Animal",
		PetSpeciesExotic:   "Exotic Pet",
		PetSpeciesWildlife: "Wildlife",
		PetSpeciesOther:    "Other",
		PetSpeciesUnknown:  "Unknown",
	}

	PetSpeciesScientificNames = map[PetSpecies]string{
		PetSpeciesDog:      "Canis lupus familiaris",
		PetSpeciesCat:      "Felis catus",
		PetSpeciesBird:     "Aves",
		PetSpeciesRabbit:   "Oryctolagus cuniculus",
		PetSpeciesHamster:  "Cricetinae",
		PetSpeciesGuinea:   "Cavia porcellus",
		PetSpeciesFerret:   "Mustela putorius furo",
		PetSpeciesReptile:  "Reptilia",
		PetSpeciesFish:     "Pisces",
		PetSpeciesHorse:    "Equus caballus",
		PetSpeciesFarm:     "Various Livestock",
		PetSpeciesExotic:   "Various Exotic Species",
		PetSpeciesWildlife: "Various Wildlife",
		PetSpeciesOther:    "Unclassified",
		PetSpeciesUnknown:  "Unknown Species",
	}

	PetSpeciesMedicalCodes = map[PetSpecies]string{
		PetSpeciesDog:      "DOG",
		PetSpeciesCat:      "CAT",
		PetSpeciesBird:     "BRD",
		PetSpeciesRabbit:   "RAB",
		PetSpeciesHamster:  "HAM",
		PetSpeciesGuinea:   "GP",
		PetSpeciesFerret:   "FER",
		PetSpeciesReptile:  "REP",
		PetSpeciesFish:     "FSH",
		PetSpeciesHorse:    "HOR",
		PetSpeciesFarm:     "FRM",
		PetSpeciesExotic:   "EXO",
		PetSpeciesWildlife: "WLD",
		PetSpeciesOther:    "OTH",
		PetSpeciesUnknown:  "UNK",
	}

	PetSpeciesIsCommon = map[PetSpecies]bool{
		PetSpeciesDog:      true,
		PetSpeciesCat:      true,
		PetSpeciesBird:     true,
		PetSpeciesRabbit:   true,
		PetSpeciesHamster:  true,
		PetSpeciesGuinea:   true,
		PetSpeciesFerret:   false,
		PetSpeciesReptile:  false,
		PetSpeciesFish:     true,
		PetSpeciesHorse:    false,
		PetSpeciesFarm:     false,
		PetSpeciesExotic:   false,
		PetSpeciesWildlife: false,
		PetSpeciesOther:    false,
		PetSpeciesUnknown:  false,
	}

	PetSpeciesRequiresSpecialCare = map[PetSpecies]bool{
		PetSpeciesDog:      false,
		PetSpeciesCat:      false,
		PetSpeciesBird:     true,
		PetSpeciesRabbit:   true,
		PetSpeciesHamster:  true,
		PetSpeciesGuinea:   true,
		PetSpeciesFerret:   true,
		PetSpeciesReptile:  true,
		PetSpeciesFish:     true,
		PetSpeciesHorse:    true,
		PetSpeciesFarm:     true,
		PetSpeciesExotic:   true,
		PetSpeciesWildlife: true,
		PetSpeciesOther:    true,
		PetSpeciesUnknown:  false,
	}

	PetSpeciesLifespan = map[PetSpecies]string{
		PetSpeciesDog:      "10-13 years",
		PetSpeciesCat:      "13-17 years",
		PetSpeciesBird:     "5-100+ years (varies by species)",
		PetSpeciesRabbit:   "8-12 years",
		PetSpeciesHamster:  "2-3 years",
		PetSpeciesGuinea:   "4-8 years",
		PetSpeciesFerret:   "7-10 years",
		PetSpeciesReptile:  "5-50+ years (varies by species)",
		PetSpeciesFish:     "1-20+ years (varies by species)",
		PetSpeciesHorse:    "25-30 years",
		PetSpeciesFarm:     "Varies by species",
		PetSpeciesExotic:   "Varies by species",
		PetSpeciesWildlife: "Varies by species",
		PetSpeciesOther:    "Unknown",
		PetSpeciesUnknown:  "Unknown",
	}
)

func (pt PetSpecies) IsValid() bool {
	for _, validType := range ValidPetSpeciess {
		if pt == validType {
			return true
		}
	}
	return false
}

func (pt PetSpecies) Validate() error {
	if !pt.IsValid() {
		return InvalidEnumParserError("PetSpecies", string(pt))
	}
	return nil
}

func ParsePetSpecies(PetSpecies string) (PetSpecies, error) {
	normalized := normalizePetSpeciesInput(PetSpecies)
	if val, exists := PetSpeciesMap[normalized]; exists {
		return val, nil
	}

	return "", InvalidEnumParserError("PetSpecies", PetSpecies)
}

func MustParsePetSpecies(PetSpecies string) PetSpecies {
	parsed, err := ParsePetSpecies(PetSpecies)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (pt PetSpecies) String() string {
	return string(pt)
}

func (pt PetSpecies) DisplayName() string {
	if displayName, exists := PetSpeciesDisplayNames[pt]; exists {
		return displayName
	}
	return "Unknown Type"
}

func (pt PetSpecies) ScientificName() string {
	if scientificName, exists := PetSpeciesScientificNames[pt]; exists {
		return scientificName
	}
	return pt.DisplayName()
}

func (pt PetSpecies) MedicalCode() string {
	if medicalCode, exists := PetSpeciesMedicalCodes[pt]; exists {
		return medicalCode
	}
	return "UNK"
}

func (pt PetSpecies) Values() []PetSpecies {
	return ValidPetSpeciess
}

func (pt PetSpecies) IsCommon() bool {
	if isCommon, exists := PetSpeciesIsCommon[pt]; exists {
		return isCommon
	}
	return false
}

func (pt PetSpecies) RequiresSpecialCare() bool {
	if requiresSpecial, exists := PetSpeciesRequiresSpecialCare[pt]; exists {
		return requiresSpecial
	}
	return true // Default to requiring special care for safety
}

func (pt PetSpecies) ExpectedLifespan() string {
	if lifespan, exists := PetSpeciesLifespan[pt]; exists {
		return lifespan
	}
	return "Unknown"
}

func (pt PetSpecies) IsDomestic() bool {
	domesticTypes := []PetSpecies{
		PetSpeciesDog, PetSpeciesCat, PetSpeciesBird, PetSpeciesRabbit,
		PetSpeciesHamster, PetSpeciesGuinea, PetSpeciesFerret, PetSpeciesFish,
	}
	for _, domesticType := range domesticTypes {
		if pt == domesticType {
			return true
		}
	}
	return false
}

func (pt PetSpecies) IsLarge() bool {
	largeTypes := []PetSpecies{PetSpeciesHorse, PetSpeciesFarm}
	for _, largeType := range largeTypes {
		if pt == largeType {
			return true
		}
	}
	return pt == PetSpeciesDog // Dogs can be large, but it varies
}

func (pt PetSpecies) RequiresVeterinaryLicense() bool {
	// Most exotic, wildlife, and farm animals require special licensing
	specialLicenseTypes := []PetSpecies{
		PetSpeciesExotic, PetSpeciesWildlife, PetSpeciesFarm, PetSpeciesHorse,
	}
	for _, specialType := range specialLicenseTypes {
		if pt == specialType {
			return true
		}
	}
	return false
}

func (pt PetSpecies) TypicalVaccinationSchedule() string {
	schedules := map[PetSpecies]string{
		PetSpeciesDog:      "Annual DHPP, Rabies",
		PetSpeciesCat:      "Annual FVRCP, Rabies",
		PetSpeciesRabbit:   "Annual RHDV, Myxomatosis",
		PetSpeciesFerret:   "Annual Distemper, Rabies",
		PetSpeciesHorse:    "Annual EEE, WEE, Tetanus, Influenza",
		PetSpeciesFarm:     "Species-specific schedule",
		PetSpeciesExotic:   "Species-specific schedule",
		PetSpeciesWildlife: "Rehabilitation protocol",
	}
	if schedule, exists := schedules[pt]; exists {
		return schedule
	}
	return "Consult veterinarian for specific schedule"
}

func (pt PetSpecies) IsComplete() bool {
	return pt != PetSpeciesUnknown && pt != PetSpeciesOther
}

// Utility functions for PetSpecies
func normalizePetSpeciesInput(input string) string {
	input = strings.TrimSpace(strings.ToLower(input))
	input = strings.ReplaceAll(input, " ", "_")
	input = strings.ReplaceAll(input, "-", "_")

	// Handle common variations
	switch input {
	case "canine", "k_9", "k9":
		return "dog"
	case "feline":
		return "cat"
	case "avian":
		return "bird"
	case "bunny", "bunnies":
		return "rabbit"
	case "guinea_pig", "cavy":
		return "guinea_pig"
	case "equine", "pony":
		return "horse"
	case "livestock", "farm":
		return "farm_animal"
	case "exotic_pet":
		return "exotic"
	}

	return input
}

func GetAllPetSpeciess() []PetSpecies {
	return ValidPetSpeciess
}

func GetCommonPetSpeciess() []PetSpecies {
	var commonTypes []PetSpecies
	for _, PetSpecies := range ValidPetSpeciess {
		if PetSpecies.IsCommon() {
			commonTypes = append(commonTypes, PetSpecies)
		}
	}
	return commonTypes
}

func GetDomesticPetSpeciess() []PetSpecies {
	var domesticTypes []PetSpecies
	for _, PetSpecies := range ValidPetSpeciess {
		if PetSpecies.IsDomestic() {
			domesticTypes = append(domesticTypes, PetSpecies)
		}
	}
	return domesticTypes
}

func GetExoticPetSpeciess() []PetSpecies {
	return []PetSpecies{
		PetSpeciesReptile,
		PetSpeciesFerret,
		PetSpeciesExotic,
		PetSpeciesWildlife,
	}
}

func GetPetSpeciessRequiringSpecialCare() []PetSpecies {
	var specialCareTypes []PetSpecies
	for _, PetSpecies := range ValidPetSpeciess {
		if PetSpecies.RequiresSpecialCare() {
			specialCareTypes = append(specialCareTypes, PetSpecies)
		}
	}
	return specialCareTypes
}

func GetPetSpeciessForEmergencyClinic() []PetSpecies {
	// Types commonly seen in emergency clinics
	return []PetSpecies{
		PetSpeciesDog,
		PetSpeciesCat,
		PetSpeciesBird,
		PetSpeciesRabbit,
		PetSpeciesFerret,
		PetSpeciesReptile,
	}
}

func SuggestPetSpeciesFromDescription(description string) PetSpecies {
	description = strings.ToLower(description)

	// Check for specific keywords
	typeKeywords := map[string]PetSpecies{
		"dog":      PetSpeciesDog,
		"puppy":    PetSpeciesDog,
		"canine":   PetSpeciesDog,
		"cat":      PetSpeciesCat,
		"kitten":   PetSpeciesCat,
		"feline":   PetSpeciesCat,
		"bird":     PetSpeciesBird,
		"parrot":   PetSpeciesBird,
		"rabbit":   PetSpeciesRabbit,
		"bunny":    PetSpeciesRabbit,
		"hamster":  PetSpeciesHamster,
		"guinea":   PetSpeciesGuinea,
		"ferret":   PetSpeciesFerret,
		"snake":    PetSpeciesReptile,
		"lizard":   PetSpeciesReptile,
		"turtle":   PetSpeciesReptile,
		"fish":     PetSpeciesFish,
		"goldfish": PetSpeciesFish,
		"horse":    PetSpeciesHorse,
		"pony":     PetSpeciesHorse,
	}

	for keyword, PetSpecies := range typeKeywords {
		if strings.Contains(description, keyword) {
			return PetSpecies
		}
	}

	return PetSpeciesUnknown
}

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

func (pg PetGender) Validate() error {
	if !pg.IsValid() {
		return InvalidEnumParserError("PetGender", string(pg))
	}
	return nil
}

func ParsePetGender(gender string) (PetGender, error) {
	normalized := normalizePetGenderInput(gender)
	if val, exists := petGenderMap[normalized]; exists {
		return val, nil
	}

	return "", InvalidEnumParserError("PetGender", gender)
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
