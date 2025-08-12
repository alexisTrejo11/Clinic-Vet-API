package models

import (
	"database/sql/driver"
	"fmt"
)

// PersonGender represents the person_gender ENUM in the database.
type PersonGender string

const (
	PersonGenderMale         PersonGender = "male"
	PersonGenderFemale       PersonGender = "female"
	PersonGenderNotSpecified PersonGender = "not_specified"
)

func (e *PersonGender) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PersonGender(s)
	case string:
		*e = PersonGender(s)
	default:
		return fmt.Errorf("unsupported type for PersonGender: %T", src)
	}
	return nil
}

func (e PersonGender) Value() (driver.Value, error) {
	return string(e), nil
}

// VeterinarianSpeciality represents the veterinarian_speciality ENUM.
type VeterinarianSpeciality string

const (
	VeterinarianSpecialityUnknownSpecialty      VeterinarianSpeciality = "unknown_specialty"
	VeterinarianSpecialityGeneralPractice       VeterinarianSpeciality = "general_practice"
	VeterinarianSpecialitySurgery               VeterinarianSpeciality = "surgery"
	VeterinarianSpecialityInternalMedicine      VeterinarianSpeciality = "internal_medicine"
	VeterinarianSpecialityDentistry             VeterinarianSpeciality = "dentistry"
	VeterinarianSpecialityDermatology           VeterinarianSpeciality = "dermatology"
	VeterinarianSpecialityOncology              VeterinarianSpeciality = "oncology"
	VeterinarianSpecialityCardiology            VeterinarianSpeciality = "cardiology"
	VeterinarianSpecialityNeurology             VeterinarianSpeciality = "neurology"
	VeterinarianSpecialityOphthalmology         VeterinarianSpeciality = "ophthalmology"
	VeterinarianSpecialityRadiology             VeterinarianSpeciality = "radiology"
	VeterinarianSpecialityEmergencyCriticalCare VeterinarianSpeciality = "emergency_critical_care"
	VeterinarianSpecialityAnesthesiology        VeterinarianSpeciality = "anesthesiology"
	VeterinarianSpecialityPathology             VeterinarianSpeciality = "pathology"
	VeterinarianSpecialityPreventiveMedicine    VeterinarianSpeciality = "preventive_medicine"
	VeterinarianSpecialityExoticAnimalMedicine  VeterinarianSpeciality = "exotic_animal_medicine"
	VeterinarianSpecialityEquineMedicine        VeterinarianSpeciality = "equine_medicine"
	VeterinarianSpecialityAvianMedicine         VeterinarianSpeciality = "avian_medicine"
	VeterinarianSpecialityZooAnimalMedicine     VeterinarianSpeciality = "zoo_animal_medicine"
	VeterinarianSpecialityFoodAnimalMedicine    VeterinarianSpeciality = "food_animal_medicine"
	VeterinarianSpecialityPublicHealth          VeterinarianSpeciality = "public_health"
)

func (e *VeterinarianSpeciality) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = VeterinarianSpeciality(s)
	case string:
		*e = VeterinarianSpeciality(s)
	default:
		return fmt.Errorf("unsupported type for VeterinarianSpeciality: %T", src)
	}
	return nil
}

func (e VeterinarianSpeciality) Value() (driver.Value, error) {
	return string(e), nil
}

// UserStatus represents the user_status ENUM.
type UserStatus string

const (
	UserStatusActive   UserStatus = "active"
	UserStatusInactive UserStatus = "inactive"
	UserStatusPending  UserStatus = "pending"
	UserStatusBanned   UserStatus = "banned"
	UserStatusDeleted  UserStatus = "deleted"
)

func (e *UserStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserStatus(s)
	case string:
		*e = UserStatus(s)
	default:
		return fmt.Errorf("unsupported type for UserStatus: %T", src)
	}
	return nil
}

func (e UserStatus) Value() (driver.Value, error) {
	return string(e), nil
}

// UserRole represents the user_role ENUM.
type UserRole string

const (
	UserRoleOwner        UserRole = "owner"
	UserRoleReceptionist UserRole = "receptionist"
	UserRoleVeterinarian UserRole = "veterinarian"
	UserRoleAdmin        UserRole = "admin"
)

func (e *UserRole) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = UserRole(s)
	case string:
		*e = UserRole(s)
	default:
		return fmt.Errorf("unsupported type for UserRole: %T", src)
	}
	return nil
}

func (e UserRole) Value() (driver.Value, error) {
	return string(e), nil
}

// Currency represents the currency ENUM.
type Currency string

const (
	CurrencyUSD Currency = "USD"
	CurrencyMXN Currency = "MXN"
	CurrencyEUR Currency = "EUR"
	CurrencyGBP Currency = "GBP"
	CurrencyJPY Currency = "JPY"
	CurrencyAUD Currency = "AUD"
	CurrencyCAD Currency = "CAD"
	CurrencyCHF Currency = "CHF"
	CurrencyCNY Currency = "CNY"
	CurrencySEK Currency = "SEK"
	CurrencyNZD Currency = "NZD"
)

func (e *Currency) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Currency(s)
	case string:
		*e = Currency(s)
	default:
		return fmt.Errorf("unsupported type for Currency: %T", src)
	}
	return nil
}

func (e Currency) Value() (driver.Value, error) {
	return string(e), nil
}

// PaymentStatus represents the payment_status ENUM.
type PaymentStatus string

const (
	PaymentStatusPending   PaymentStatus = "pending"
	PaymentStatusCompleted PaymentStatus = "completed"
	PaymentStatusFailed    PaymentStatus = "failed"
	PaymentStatusRefunded  PaymentStatus = "refunded"
)

func (e *PaymentStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PaymentStatus(s)
	case string:
		*e = PaymentStatus(s)
	default:
		return fmt.Errorf("unsupported type for PaymentStatus: %T", src)
	}
	return nil
}

func (e PaymentStatus) Value() (driver.Value, error) {
	return string(e), nil
}

// PaymentMethod represents the payment_method ENUM.
type PaymentMethod string

const (
	PaymentMethodCash         PaymentMethod = "cash"
	PaymentMethodCreditCard   PaymentMethod = "credit_card"
	PaymentMethodDebitCard    PaymentMethod = "debit_card"
	PaymentMethodBankTransfer PaymentMethod = "bank_transfer"
	PaymentMethodPaypal       PaymentMethod = "paypal"
	PaymentMethodStripe       PaymentMethod = "stripe"
	PaymentMethodCheck        PaymentMethod = "check"
)

func (e *PaymentMethod) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = PaymentMethod(s)
	case string:
		*e = PaymentMethod(s)
	default:
		return fmt.Errorf("unsupported type for PaymentMethod: %T", src)
	}
	return nil
}

func (e PaymentMethod) Value() (driver.Value, error) {
	return string(e), nil
}

// ClinicService represents the clinic_service ENUM.
type ClinicService string

const (
	ClinicServiceGeneralConsultation ClinicService = "general_consultation"
	ClinicServiceVaccination         ClinicService = "vaccination"
	ClinicServiceSurgery             ClinicService = "surgery"
	ClinicServiceDentalCare          ClinicService = "dental_care"
	ClinicServiceEmergencyCare       ClinicService = "emergency_care"
	ClinicServiceGrooming            ClinicService = "grooming"
	ClinicServiceNutritionConsult    ClinicService = "nutrition_consult"
	ClinicServiceBehaviorConsult     ClinicService = "behavior_consult"
	ClinicServiceWellnessExam        ClinicService = "wellness_exam"
	ClinicServiceOther               ClinicService = "other"
)

func (e *ClinicService) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = ClinicService(s)
	case string:
		*e = ClinicService(s)
	default:
		return fmt.Errorf("unsupported type for ClinicService: %T", src)
	}
	return nil
}

func (e ClinicService) Value() (driver.Value, error) {
	return string(e), nil
}

// AppointmentStatus represents the appointment_status ENUM.
type AppointmentStatus string

const (
	AppointmentStatusPending      AppointmentStatus = "pending"
	AppointmentStatusCancelled    AppointmentStatus = "cancelled"
	AppointmentStatusConfirmed    AppointmentStatus = "confirmed"
	AppointmentStatusRescheduled  AppointmentStatus = "rescheduled"
	AppointmentStatusCompleted    AppointmentStatus = "completed"
	AppointmentStatusNotPresented AppointmentStatus = "not_presented"
)

func (e *AppointmentStatus) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = AppointmentStatus(s)
	case string:
		*e = AppointmentStatus(s)
	default:
		return fmt.Errorf("unsupported type for AppointmentStatus: %T", src)
	}
	return nil
}

func (e AppointmentStatus) Value() (driver.Value, error) {
	return string(e), nil
}
