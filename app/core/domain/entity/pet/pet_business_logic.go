package pet

import "time"

func (p *Pet) Activate() error {
	if p.isActive {
		return nil // Already active
	}
	p.isActive = true
	p.deletedAt = nil
	p.IncrementVersion()
	return nil
}

func (p *Pet) Deactivate() error {
	if !p.isActive {
		return nil // Already inactive
	}
	p.isActive = false
	p.IncrementVersion()
	return nil
}

func (p *Pet) SoftDelete() error {
	if p.deletedAt != nil {
		return nil // Already deleted
	}
	now := time.Now()
	p.isActive = false
	p.deletedAt = &now
	p.IncrementVersion()
	return nil
}

func (p *Pet) Restore() error {
	if p.deletedAt == nil {
		return nil // Not deleted
	}
	p.isActive = true
	p.deletedAt = nil
	p.IncrementVersion()
	return nil
}

func (p *Pet) RequiresVaccination() bool {
	// Logic to determine if pet needs vaccination based on age and species
	if p.age == nil {
		return false
	}

	// Puppies/kittens need more frequent vaccinations
	if *p.age < 1 {
		return true
	}

	// Adult pets need annual vaccinations
	return true
}

func (p *Pet) LifeStage() string {
	if p.age == nil {
		return "unknown"
	}

	age := *p.age
	switch {
	case age < 1:
		return "baby"
	case age < 3:
		return "young"
	case age < 8:
		return "adult"
	default:
		return "senior"
	}
}

func (p *Pet) HasMedicalConditions() bool {
	return (p.allergies != nil && *p.allergies != "") ||
		(p.currentMedications != nil && *p.currentMedications != "") ||
		(p.specialNeeds != nil && *p.specialNeeds != "")
}
