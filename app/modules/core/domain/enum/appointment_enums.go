// Package enum for all domain core logic
package enum

import (
	domainerr "clinic-vet-api/app/modules/core/error"
	"context"
	"fmt"
	"slices"
	"strings"
)

func InvalidEnumParserError(enumName, enumValue string) error {
	return domainerr.InvalidEnumValue(context.Background(), enumValue, enumValue, fmt.Sprintf("invalid %s value: %s", enumName, enumValue), "enum parse")

}

// AppointmentStatus represents the status of an appointment
type AppointmentStatus string

const (
	AppointmentStatusPending      AppointmentStatus = "pending"
	AppointmentStatusCancelled    AppointmentStatus = "cancelled"
	AppointmentStatusCompleted    AppointmentStatus = "completed"
	AppointmentStatusRescheduled  AppointmentStatus = "rescheduled"
	AppointmentStatusConfirmed    AppointmentStatus = "confirmed"
	AppointmentStatusNotPresented AppointmentStatus = "not_presented"
)

// AppointmentStatus constants and methods
var (
	ValidAppointmentStatuses = []AppointmentStatus{
		AppointmentStatusPending,
		AppointmentStatusCancelled,
		AppointmentStatusCompleted,
		AppointmentStatusRescheduled,
		AppointmentStatusConfirmed,
		AppointmentStatusNotPresented,
	}

	appointmentStatusMap = map[string]AppointmentStatus{
		"pending":       AppointmentStatusPending,
		"cancelled":     AppointmentStatusCancelled,
		"completed":     AppointmentStatusCompleted,
		"rescheduled":   AppointmentStatusRescheduled,
		"confirmed":     AppointmentStatusConfirmed,
		"not_presented": AppointmentStatusNotPresented,
		"not presented": AppointmentStatusNotPresented,
	}

	appointmentStatusDisplayNames = map[AppointmentStatus]string{
		AppointmentStatusPending:      "Pending",
		AppointmentStatusCancelled:    "Cancelled",
		AppointmentStatusCompleted:    "Completed",
		AppointmentStatusRescheduled:  "Rescheduled",
		AppointmentStatusConfirmed:    "Confirmed",
		AppointmentStatusNotPresented: "Not Presented",
	}
)

func (as AppointmentStatus) IsValid() bool {
	_, exists := appointmentStatusMap[string(as)]
	return exists
}

func ParseAppointmentStatus(status string) (AppointmentStatus, error) {
	normalized := strings.TrimSpace(strings.ToLower(status))
	if val, exists := appointmentStatusMap[normalized]; exists {
		return val, nil
	}

	return "", InvalidEnumParserError("AppointmentStatus", status)
}

func MustParseAppointmentStatus(status string) AppointmentStatus {
	parsed, err := ParseAppointmentStatus(status)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (as AppointmentStatus) String() string {
	return string(as)
}

func (as AppointmentStatus) DisplayName() string {
	if displayName, exists := appointmentStatusDisplayNames[as]; exists {
		return displayName
	}
	return "Unknown Status"
}

func (as AppointmentStatus) Values() []AppointmentStatus {
	return ValidAppointmentStatuses
}

// CanBeCancelled checks if an appointment status allows cancellation
func (as AppointmentStatus) CanBeCancelled() bool {
	return as == AppointmentStatusPending || as == AppointmentStatusConfirmed
}

// IsFinalStatus checks if the status is a final state (no further changes expected)
func (as AppointmentStatus) IsFinalStatus() bool {
	return as == AppointmentStatusCompleted ||
		as == AppointmentStatusCancelled ||
		as == AppointmentStatusNotPresented
}

// CanBeRescheduled checks if an appointment can be rescheduled from current status
func (as AppointmentStatus) CanBeRescheduled() bool {
	return as == AppointmentStatusPending ||
		as == AppointmentStatusConfirmed ||
		as == AppointmentStatusRescheduled
}

// ClinicService represents the types of services offered by the clinic
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

// ClinicService constants and methods
var (
	ValidClinicServices = []ClinicService{
		ClinicServiceGeneralConsultation,
		ClinicServiceVaccination,
		ClinicServiceSurgery,
		ClinicServiceDentalCare,
		ClinicServiceEmergencyCare,
		ClinicServiceGrooming,
		ClinicServiceNutritionConsult,
		ClinicServiceBehaviorConsult,
		ClinicServiceWellnessExam,
		ClinicServiceOther,
	}

	clinicServiceMap = map[string]ClinicService{
		"general_consultation": ClinicServiceGeneralConsultation,
		"general consultation": ClinicServiceGeneralConsultation,
		"consultation":         ClinicServiceGeneralConsultation,
		"vaccination":          ClinicServiceVaccination,
		"surgery":              ClinicServiceSurgery,
		"dental_care":          ClinicServiceDentalCare,
		"dental care":          ClinicServiceDentalCare,
		"dental":               ClinicServiceDentalCare,
		"emergency_care":       ClinicServiceEmergencyCare,
		"emergency care":       ClinicServiceEmergencyCare,
		"emergency":            ClinicServiceEmergencyCare,
		"grooming":             ClinicServiceGrooming,
		"nutrition_consult":    ClinicServiceNutritionConsult,
		"nutrition consult":    ClinicServiceNutritionConsult,
		"nutrition":            ClinicServiceNutritionConsult,
		"behavior_consult":     ClinicServiceBehaviorConsult,
		"behavior consult":     ClinicServiceBehaviorConsult,
		"behavior":             ClinicServiceBehaviorConsult,
		"wellness_exam":        ClinicServiceWellnessExam,
		"wellness exam":        ClinicServiceWellnessExam,
		"wellness":             ClinicServiceWellnessExam,
		"other":                ClinicServiceOther,
	}

	clinicServiceDisplayNames = map[ClinicService]string{
		ClinicServiceGeneralConsultation: "General Consultation",
		ClinicServiceVaccination:         "Vaccination",
		ClinicServiceSurgery:             "Surgery",
		ClinicServiceDentalCare:          "Dental Care",
		ClinicServiceEmergencyCare:       "Emergency Care",
		ClinicServiceGrooming:            "Grooming",
		ClinicServiceNutritionConsult:    "Nutrition Consultation",
		ClinicServiceBehaviorConsult:     "Behavior Consultation",
		ClinicServiceWellnessExam:        "Wellness Exam",
		ClinicServiceOther:               "Other Service",
	}

	clinicServiceDurations = map[ClinicService]int{
		ClinicServiceGeneralConsultation: 30,
		ClinicServiceVaccination:         15,
		ClinicServiceSurgery:             120,
		ClinicServiceDentalCare:          45,
		ClinicServiceEmergencyCare:       60,
		ClinicServiceGrooming:            60,
		ClinicServiceNutritionConsult:    45,
		ClinicServiceBehaviorConsult:     60,
		ClinicServiceWellnessExam:        30,
		ClinicServiceOther:               30,
	}
)

func (cs ClinicService) IsValid() bool {
	_, exists := clinicServiceMap[string(cs)]
	return exists
}

func ParseClinicService(service string) (ClinicService, error) {
	normalized := normalizeServiceInput(service)
	if val, exists := clinicServiceMap[normalized]; exists {
		return val, nil
	}
	return "", fmt.Errorf("invalid clinic service: %s", service)
}

func MustParseClinicService(service string) ClinicService {
	parsed, err := ParseClinicService(service)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (cs ClinicService) String() string {
	return string(cs)
}

func (cs ClinicService) DisplayName() string {
	if displayName, exists := clinicServiceDisplayNames[cs]; exists {
		return displayName
	}
	return "Unknown Service"
}

func (cs ClinicService) Values() []ClinicService {
	return ValidClinicServices
}

func (cs ClinicService) DefaultDuration() int {
	if duration, exists := clinicServiceDurations[cs]; exists {
		return duration
	}
	return 30 // default duration in minutes
}

func (cs ClinicService) IsMedicalService() bool {
	medicalServices := []ClinicService{
		ClinicServiceGeneralConsultation,
		ClinicServiceVaccination,
		ClinicServiceSurgery,
		ClinicServiceDentalCare,
		ClinicServiceEmergencyCare,
		ClinicServiceNutritionConsult,
		ClinicServiceBehaviorConsult,
		ClinicServiceWellnessExam,
	}

	return slices.Contains(medicalServices, cs)
}

func (cs ClinicService) IsGroomingService() bool {
	return cs == ClinicServiceGrooming
}

func normalizeServiceInput(input string) string {
	input = strings.TrimSpace(strings.ToLower(input))
	input = strings.ReplaceAll(input, " ", "_")
	return input
}

func GetAllAppointmentStatuses() []AppointmentStatus {
	return ValidAppointmentStatuses
}

func GetAllClinicServices() []ClinicService {
	return ValidClinicServices
}

// GetActiveAppointmentStatuses returns statuses that represent active appointments
func GetActiveAppointmentStatuses() []AppointmentStatus {
	return []AppointmentStatus{
		AppointmentStatusPending,
		AppointmentStatusConfirmed,
		AppointmentStatusRescheduled,
	}
}
