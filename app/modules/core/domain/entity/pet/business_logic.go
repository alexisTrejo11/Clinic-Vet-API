package pet

import "context"

func (p *Pet) Activate() error {
	if p.isActive {
		return nil // Already active
	}
	p.isActive = true
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

func (p *Pet) ValidatePersistence(ctx context.Context) error {
	operation := "ValidatePet"

	if p.customerID.IsZero() {
		return CustomerIDRequiredError(ctx, operation)
	}

	if p.name == "" {
		return NameRequiredError(ctx, operation)
	}
	if len(p.name) > 100 {
		return NameTooLongError(ctx, len(p.name), operation)
	}

	if p.species == "" {
		return SpeciesRequiredError(ctx, operation)
	}

	if !p.species.IsValid() {
		return InvalidSpeciesError(ctx, p.species.String())
	}

	if p.age != nil {
		if *p.age < 0 {
			return AgeInvalidError(ctx, *p.age, operation)
		}
		if *p.age > 50 {
			return AgeUnrealisticError(ctx, *p.age, operation)
		}
	}

	if !p.gender.IsValid() {
		return GenderInvalidError(ctx, p.gender, operation)
	}

	if p.breed != nil && len(*p.breed) > 50 {
		return BreedTooLongError(ctx, len(*p.breed), operation)
	}

	if p.microchip != nil && len(*p.microchip) > 50 {
		return MicrochipTooLongError(ctx, len(*p.microchip), operation)
	}

	if p.color != nil && len(*p.color) > 30 {
		return ColorTooLongError(ctx, len(*p.color), operation)
	}

	if p.photo != nil && len(*p.photo) > 500 {
		return PhotoURLTooLongError(ctx, len(*p.photo), operation)
	}

	return nil
}
