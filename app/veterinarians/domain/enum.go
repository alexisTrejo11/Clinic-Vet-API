package vetDomain

import (
	"fmt"
	"strings"
)

type VetSpecialty int

const (
	UnknownSpecialty               VetSpecialty = iota // 0
	GeneralPracticeSpecialty                           // 1
	SurgerySpecialty                                   // 2
	InternalMedicineSpecialty                          // 3
	DentistrySpecialty                                 // 4
	DermatologySpecialty                               // 5
	OncologySpecialty                                  // 6
	CardiologySpecialty                                // 7
	NeurologySpecialty                                 // 8
	OphthalmologySpecialty                             // 9
	RadiologySpecialty                                 // 10
	EmergencyCriticalCareSpecialty                     // 11
	AnesthesiologySpecialty                            // 12
	PathologySpecialty                                 // 13
	PreventiveMedicineSpecialty                        // 14
	ExoticAnimalMedicineSpecialty                      // 15
	EquineMedicineSpecialty                            // 16
	AvianMedicineSpecialty                             // 17
	ZooAnimalMedicineSpecialty                         // 18
	FoodAnimalMedicineSpecialty                        // 19
	PublicHealthSpecialty                              // 20
)

func (s VetSpecialty) String() string {
	switch s {
	case UnknownSpecialty:
		return "unknown_specialty"
	case GeneralPracticeSpecialty:
		return "general_practice"
	case SurgerySpecialty:
		return "surgery"
	case InternalMedicineSpecialty:
		return "internal_medicine"
	case DentistrySpecialty:
		return "dentistry"
	case DermatologySpecialty:
		return "dermatology"
	case OncologySpecialty:
		return "oncology"
	case CardiologySpecialty:
		return "cardiology"
	case NeurologySpecialty:
		return "neurology"
	case OphthalmologySpecialty:
		return "ophthalmology"
	case RadiologySpecialty:
		return "radiology"
	case EmergencyCriticalCareSpecialty:
		return "emergency_critical_care"
	case AnesthesiologySpecialty:
		return "anesthesiology"
	case PathologySpecialty:
		return "pathology"
	case PreventiveMedicineSpecialty:
		return "preventive_medicine"
	case ExoticAnimalMedicineSpecialty:
		return "exotic_animal_medicine"
	case EquineMedicineSpecialty:
		return "equine_medicine"
	case AvianMedicineSpecialty:
		return "avian_medicine"
	case ZooAnimalMedicineSpecialty:
		return "zoo_animal_medicine"
	case FoodAnimalMedicineSpecialty:
		return "food_animal_medicine"
	case PublicHealthSpecialty:
		return "public_health"
	default:
		return fmt.Sprintf("Specialty(%d)", s)
	}
}

func VetSpecialtyFromString(s string) VetSpecialty {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "unknown_specialty":
		return UnknownSpecialty
	case "general_practice":
		return GeneralPracticeSpecialty
	case "surgery":
		return SurgerySpecialty
	case "internal_medicine":
		return InternalMedicineSpecialty
	case "dentistry":
		return DentistrySpecialty
	case "dermatology":
		return DermatologySpecialty
	case "oncology":
		return OncologySpecialty
	case "cardiology":
		return CardiologySpecialty
	case "neurology":
		return NeurologySpecialty
	case "ophthalmology":
		return OphthalmologySpecialty
	case "radiology":
		return RadiologySpecialty
	case "emergency_critical_care":
		return EmergencyCriticalCareSpecialty
	case "anesthesiology":
		return AnesthesiologySpecialty
	case "pathology":
		return PathologySpecialty
	case "preventive_medicine":
		return PreventiveMedicineSpecialty
	case "exotic_animal_medicine":
		return ExoticAnimalMedicineSpecialty
	case "equine_medicine":
		return EquineMedicineSpecialty
	case "avian_medicine":
		return AvianMedicineSpecialty
	case "zoo_animal_medicine":
		return ZooAnimalMedicineSpecialty
	case "food_animal_medicine":
		return FoodAnimalMedicineSpecialty
	case "public_health":
		return PublicHealthSpecialty
	default:
		return UnknownSpecialty
	}
}
