// Package enum for all domain core logic
package enum

import (
	"errors"
)

type AppointmentStatus string

const (
	StatusPending      AppointmentStatus = "pending"
	StatusCancelled    AppointmentStatus = "cancelled"
	StatusCompleted    AppointmentStatus = "completed"
	StatusRescheduled  AppointmentStatus = "rescheduled"
	StatusConfirmed    AppointmentStatus = "confirmed"
	StatusNotPresented AppointmentStatus = "not_presented"
)

func (as AppointmentStatus) IsValid() bool {
	switch as {
	case StatusPending, StatusCancelled, StatusCompleted, StatusRescheduled, StatusConfirmed, StatusNotPresented:
		return true
	default:
		return false
	}
}

func NewAppointmentStatus(status string) (AppointmentStatus, error) {
	as := AppointmentStatus(status)
	if !as.IsValid() {
		return "", errors.New("invalid appointment status")
	}
	return as, nil
}

type ClinicService string

const (
	ServiceGeneralConsultation ClinicService = "general_consultation"
	ServiceVaccination         ClinicService = "vaccination"
	ServiceSurgery             ClinicService = "surgery"
	ServiceDentalCare          ClinicService = "dental_care"
	ServiceEmergencyCare       ClinicService = "emergency_care"
	ServiceGrooming            ClinicService = "grooming"
	ServiceNutritionConsult    ClinicService = "nutrition_consult"
	ServiceBehaviorConsult     ClinicService = "behavior_consult"
	ServiceWellnessExam        ClinicService = "wellness_exam"
	ServiceOther               ClinicService = "other"
)

func NewClinicService(name string) (ClinicService, error) {
	return ServiceBehaviorConsult, nil
}
