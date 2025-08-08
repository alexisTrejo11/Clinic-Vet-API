package appointDomain

type AppointmentStatus string

const (
	StatusPending      AppointmentStatus = "pending"
	StatusCancelled    AppointmentStatus = "cancelled"
	StatusCompleted    AppointmentStatus = "completed"
	StatusRescheduled  AppointmentStatus = "rescheduled"
	StatusNotPresented AppointmentStatus = "not_presented"
)

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
