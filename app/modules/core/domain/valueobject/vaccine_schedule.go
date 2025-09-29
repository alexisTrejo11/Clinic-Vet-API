package valueobject

import "time"

type VaccineSchedule struct {
	initialDoses    int
	intervalBetween time.Duration
	boosterInterval time.Duration
	minAgeForFirst  time.Duration
	maxAgeForFirst  *time.Duration
	isLifelong      bool
}

func NewVaccineSchedule(
	initialDoses int,
	intervalBetween time.Duration,
	boosterInterval time.Duration,
	minAgeForFirst time.Duration,
) VaccineSchedule {
	return VaccineSchedule{
		initialDoses:    initialDoses,
		intervalBetween: intervalBetween,
		boosterInterval: boosterInterval,
		minAgeForFirst:  minAgeForFirst,
		isLifelong:      false,
	}
}

func (vs VaccineSchedule) InitialDoses() int              { return vs.initialDoses }
func (vs VaccineSchedule) IntervalBetween() time.Duration { return vs.intervalBetween }
func (vs VaccineSchedule) BoosterInterval() time.Duration { return vs.boosterInterval }
func (vs VaccineSchedule) MinAgeForFirst() time.Duration  { return vs.minAgeForFirst }
func (vs VaccineSchedule) IsLifelong() bool               { return vs.isLifelong }

func (vs VaccineSchedule) WithMaxAge(maxAge time.Duration) VaccineSchedule {
	vs.maxAgeForFirst = &maxAge
	return vs
}

func (vs VaccineSchedule) WithLifelong(isLifelong bool) VaccineSchedule {
	vs.isLifelong = isLifelong
	return vs
}

type VaccineDefinition struct {
	name         VaccineName
	vaccineType  VaccineType
	species      []string
	protectsFrom []string
	schedule     VaccineSchedule
	description  string
	sideEffects  []string
	contraindic  []string
}

func NewVaccineDefinition(
	name VaccineName,
	vaccineType VaccineType,
	species []string,
	protectsFrom []string,
	schedule VaccineSchedule,
) VaccineDefinition {
	return VaccineDefinition{
		name:         name,
		vaccineType:  vaccineType,
		species:      species,
		protectsFrom: protectsFrom,
		schedule:     schedule,
	}
}

func (vd VaccineDefinition) Name() VaccineName           { return vd.name }
func (vd VaccineDefinition) VaccineType() VaccineType    { return vd.vaccineType }
func (vd VaccineDefinition) Species() []string           { return vd.species }
func (vd VaccineDefinition) ProtectsFrom() []string      { return vd.protectsFrom }
func (vd VaccineDefinition) Schedule() VaccineSchedule   { return vd.schedule }
func (vd VaccineDefinition) Description() string         { return vd.description }
func (vd VaccineDefinition) SideEffects() []string       { return vd.sideEffects }
func (vd VaccineDefinition) Contraindications() []string { return vd.contraindic }

func (vd VaccineDefinition) WithDescription(desc string) VaccineDefinition {
	vd.description = desc
	return vd
}

func (vd VaccineDefinition) WithSideEffects(effects []string) VaccineDefinition {
	vd.sideEffects = effects
	return vd
}

func (vd VaccineDefinition) WithContraindications(contraindic []string) VaccineDefinition {
	vd.contraindic = contraindic
	return vd
}

func (vd VaccineDefinition) IsApplicableForSpecies(species string) bool {
	for _, s := range vd.species {
		if s == species {
			return true
		}
	}
	return false
}

func (vd VaccineDefinition) CalculateNextDueDate(lastVaccinationDate time.Time, doseNumber int) time.Time {
	if doseNumber < vd.schedule.initialDoses {
		return lastVaccinationDate.Add(vd.schedule.intervalBetween)
	}
	return lastVaccinationDate.Add(vd.schedule.boosterInterval)
}
