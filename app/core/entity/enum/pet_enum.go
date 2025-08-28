package enum

type PetGender string

const (
	GenderMale    PetGender = "Male"
	GenderFemale  PetGender = "Female"
	GenderUnknown PetGender = "Unknown"
)

func (g PetGender) IsValid() bool {
	switch g {
	case GenderMale, GenderFemale, GenderUnknown:
		return true
	}
	return false
}
