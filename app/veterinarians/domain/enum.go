package vetDomain

import "fmt"

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
