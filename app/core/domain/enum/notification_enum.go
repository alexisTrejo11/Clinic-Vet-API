package enum

import (
	domainerr "clinic-vet-api/app/core/error"
	"fmt"
	"slices"
)

// NotificationType represents the type of notification
type NotificationType string

const (
	NotificationTypeAlert              NotificationType = "alert"
	NotificationTypeReminder           NotificationType = "reminder"
	NotificationTypeInfo               NotificationType = "info"
	NotificationTypeVerificationToken  NotificationType = "verification_token"
	NotificationTypeActivationToken    NotificationType = "activation_token"
	NotificationTypePasswordReset      NotificationType = "password_reset"
	NotificationTypeAppointmentConfirm NotificationType = "appointment_confirm"
	NotificationTypeAppointmentRemind  NotificationType = "appointment_remind"
	NotificationTypeAppointmentCancel  NotificationType = "appointment_cancel"
	NotificationTypePaymentReceipt     NotificationType = "payment_receipt"
	NotificationTypePaymentReminder    NotificationType = "payment_reminder"
	NotificationTypeWelcome            NotificationType = "welcome"
	NotificationTypeSecurity           NotificationType = "security"
	NotificationTypePromotional        NotificationType = "promotional"
	NotificationTypeUnknown            NotificationType = "unknown"
)

// NotificationType constants and methods
var (
	ValidNotificationTypes = []NotificationType{
		NotificationTypeAlert,
		NotificationTypeReminder,
		NotificationTypeInfo,
		NotificationTypeVerificationToken,
		NotificationTypeActivationToken,
		NotificationTypePasswordReset,
		NotificationTypeAppointmentConfirm,
		NotificationTypeAppointmentRemind,
		NotificationTypeAppointmentCancel,
		NotificationTypePaymentReceipt,
		NotificationTypePaymentReminder,
		NotificationTypeWelcome,
		NotificationTypeSecurity,
		NotificationTypePromotional,
		NotificationTypeUnknown,
	}

	notificationTypeMap = map[string]NotificationType{
		"alert":               NotificationTypeAlert,
		"reminder":            NotificationTypeReminder,
		"info":                NotificationTypeInfo,
		"information":         NotificationTypeInfo,
		"verification_token":  NotificationTypeVerificationToken,
		"verification token":  NotificationTypeVerificationToken,
		"verification":        NotificationTypeVerificationToken,
		"activation_token":    NotificationTypeActivationToken,
		"activation token":    NotificationTypeActivationToken,
		"activation":          NotificationTypeActivationToken,
		"password_reset":      NotificationTypePasswordReset,
		"password reset":      NotificationTypePasswordReset,
		"reset":               NotificationTypePasswordReset,
		"appointment_confirm": NotificationTypeAppointmentConfirm,
		"appointment confirm": NotificationTypeAppointmentConfirm,
		"confirm":             NotificationTypeAppointmentConfirm,
		"appointment_remind":  NotificationTypeAppointmentRemind,
		"appointment remind":  NotificationTypeAppointmentRemind,
		"remind":              NotificationTypeAppointmentRemind,
		"appointment_cancel":  NotificationTypeAppointmentCancel,
		"appointment cancel":  NotificationTypeAppointmentCancel,
		"cancel":              NotificationTypeAppointmentCancel,
		"payment_receipt":     NotificationTypePaymentReceipt,
		"payment receipt":     NotificationTypePaymentReceipt,
		"receipt":             NotificationTypePaymentReceipt,
		"payment_reminder":    NotificationTypePaymentReminder,
		"payment reminder":    NotificationTypePaymentReminder,
		"welcome":             NotificationTypeWelcome,
		"security":            NotificationTypeSecurity,
		"promotional":         NotificationTypePromotional,
		"promo":               NotificationTypePromotional,
		"unknown":             NotificationTypeUnknown,
	}

	notificationTypeDisplayNames = map[NotificationType]string{
		NotificationTypeAlert:              "Alert",
		NotificationTypeReminder:           "Reminder",
		NotificationTypeInfo:               "Information",
		NotificationTypeVerificationToken:  "Verification Token",
		NotificationTypeActivationToken:    "Activation Token",
		NotificationTypePasswordReset:      "Password Reset",
		NotificationTypeAppointmentConfirm: "Appointment Confirmation",
		NotificationTypeAppointmentRemind:  "Appointment Reminder",
		NotificationTypeAppointmentCancel:  "Appointment Cancellation",
		NotificationTypePaymentReceipt:     "Payment Receipt",
		NotificationTypePaymentReminder:    "Payment Reminder",
		NotificationTypeWelcome:            "Welcome Message",
		NotificationTypeSecurity:           "Security Alert",
		NotificationTypePromotional:        "Promotional Message",
		NotificationTypeUnknown:            "Unknown Notification",
	}

	notificationTypeCategories = map[NotificationType]string{
		NotificationTypeAlert:              "urgent",
		NotificationTypeReminder:           "scheduled",
		NotificationTypeInfo:               "informational",
		NotificationTypeVerificationToken:  "authentication",
		NotificationTypeActivationToken:    "authentication",
		NotificationTypePasswordReset:      "authentication",
		NotificationTypeAppointmentConfirm: "appointment",
		NotificationTypeAppointmentRemind:  "appointment",
		NotificationTypeAppointmentCancel:  "appointment",
		NotificationTypePaymentReceipt:     "financial",
		NotificationTypePaymentReminder:    "financial",
		NotificationTypeWelcome:            "onboarding",
		NotificationTypeSecurity:           "security",
		NotificationTypePromotional:        "marketing",
		NotificationTypeUnknown:            "unknown",
	}

	notificationTypePriorities = map[NotificationType]int{
		NotificationTypeAlert:              1, // Highest priority
		NotificationTypeSecurity:           1,
		NotificationTypeVerificationToken:  2,
		NotificationTypeActivationToken:    2,
		NotificationTypePasswordReset:      2,
		NotificationTypeAppointmentCancel:  2,
		NotificationTypePaymentReminder:    3,
		NotificationTypeAppointmentRemind:  3,
		NotificationTypeAppointmentConfirm: 4,
		NotificationTypePaymentReceipt:     4,
		NotificationTypeReminder:           4,
		NotificationTypeWelcome:            5,
		NotificationTypeInfo:               5,
		NotificationTypePromotional:        6, // Lowest priority
		NotificationTypeUnknown:            6,
	}
)

func (nt NotificationType) IsValid() bool {
	_, exists := notificationTypeMap[string(nt)]
	return exists
}

func ParseNotificationType(notificationType string) (NotificationType, error) {
	normalized := normalizeInput(notificationType)
	if val, exists := notificationTypeMap[normalized]; exists {
		return val, nil
	}
	return NotificationTypeUnknown, domainerr.InvalidEnumValue("notification-type", notificationType, "invalid notification type")
}

func MustParseNotificationType(notificationType string) NotificationType {
	parsed, err := ParseNotificationType(notificationType)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (nt NotificationType) String() string {
	return string(nt)
}

func (nt NotificationType) DisplayName() string {
	if displayName, exists := notificationTypeDisplayNames[nt]; exists {
		return displayName
	}
	return "Unknown Notification Type"
}

func (nt NotificationType) Values() []NotificationType {
	return ValidNotificationTypes
}

func (nt NotificationType) Category() string {
	if category, exists := notificationTypeCategories[nt]; exists {
		return category
	}
	return "unknown"
}

func (nt NotificationType) Priority() int {
	if priority, exists := notificationTypePriorities[nt]; exists {
		return priority
	}
	return 6 // Default to lowest priority
}

func (nt NotificationType) IsUrgent() bool {
	return nt.Priority() <= 2
}

func (nt NotificationType) IsAuthentication() bool {
	return nt.Category() == "authentication"
}

func (nt NotificationType) IsAppointmentRelated() bool {
	return nt.Category() == "appointment"
}

func (nt NotificationType) IsFinancial() bool {
	return nt.Category() == "financial"
}

func (nt NotificationType) RequiresImmediateDelivery() bool {
	return nt.IsUrgent() || nt.IsAuthentication()
}

func (nt NotificationType) CanBeDelayed() bool {
	return nt.Category() == "marketing" || nt.Category() == "informational"
}

// NotificationChannel represents the delivery channel for notifications
type NotificationChannel string

const (
	NotificationChannelSMS      NotificationChannel = "sms"
	NotificationChannelEmail    NotificationChannel = "email"
	NotificationChannelPush     NotificationChannel = "push"
	NotificationChannelInApp    NotificationChannel = "in_app"
	NotificationChannelVoice    NotificationChannel = "voice"
	NotificationChannelWhatsApp NotificationChannel = "whatsapp"
	NotificationChannelUnknown  NotificationChannel = "unknown"
)

// NotificationChannel constants and methods
var (
	ValidNotificationChannels = []NotificationChannel{
		NotificationChannelSMS,
		NotificationChannelEmail,
		NotificationChannelPush,
		NotificationChannelInApp,
		NotificationChannelVoice,
		NotificationChannelWhatsApp,
		NotificationChannelUnknown,
	}

	notificationChannelMap = map[string]NotificationChannel{
		"sms":               NotificationChannelSMS,
		"text":              NotificationChannelSMS,
		"email":             NotificationChannelEmail,
		"mail":              NotificationChannelEmail,
		"push":              NotificationChannelPush,
		"push_notification": NotificationChannelPush,
		"in_app":            NotificationChannelInApp,
		"in app":            NotificationChannelInApp,
		"app":               NotificationChannelInApp,
		"voice":             NotificationChannelVoice,
		"call":              NotificationChannelVoice,
		"whatsapp":          NotificationChannelWhatsApp,
		"wa":                NotificationChannelWhatsApp,
		"unknown":           NotificationChannelUnknown,
	}

	notificationChannelDisplayNames = map[NotificationChannel]string{
		NotificationChannelSMS:      "SMS",
		NotificationChannelEmail:    "Email",
		NotificationChannelPush:     "Push Notification",
		NotificationChannelInApp:    "In-App Notification",
		NotificationChannelVoice:    "Voice Call",
		NotificationChannelWhatsApp: "WhatsApp",
		NotificationChannelUnknown:  "Unknown Channel",
	}

	notificationChannelCapabilities = map[NotificationChannel][]string{
		NotificationChannelSMS:      {"text", "urgent", "offline"},
		NotificationChannelEmail:    {"text", "html", "attachments", "async"},
		NotificationChannelPush:     {"text", "rich", "interactive", "real_time"},
		NotificationChannelInApp:    {"text", "rich", "interactive", "persistent"},
		NotificationChannelVoice:    {"audio", "urgent", "interactive"},
		NotificationChannelWhatsApp: {"text", "rich", "media", "interactive"},
	}
)

func (nc NotificationChannel) IsValid() bool {
	_, exists := notificationChannelMap[string(nc)]
	return exists
}

func ParseNotificationChannel(channel string) (NotificationChannel, error) {
	normalized := normalizeInput(channel)
	if val, exists := notificationChannelMap[normalized]; exists {
		return val, nil
	}
	return NotificationChannelUnknown, fmt.Errorf("invalid notification channel: %s", channel)
}

func MustParseNotificationChannel(channel string) NotificationChannel {
	parsed, err := ParseNotificationChannel(channel)
	if err != nil {
		panic(err)
	}
	return parsed
}

func (nc NotificationChannel) String() string {
	return string(nc)
}

func (nc NotificationChannel) DisplayName() string {
	if displayName, exists := notificationChannelDisplayNames[nc]; exists {
		return displayName
	}
	return "Unknown Channel"
}

func (nc NotificationChannel) Values() []NotificationChannel {
	return ValidNotificationChannels
}

func (nc NotificationChannel) Capabilities() []string {
	if capabilities, exists := notificationChannelCapabilities[nc]; exists {
		return capabilities
	}
	return []string{"text"}
}

func (nc NotificationChannel) SupportsRichContent() bool {
	capabilities := nc.Capabilities()
	for _, cap := range capabilities {
		if cap == "rich" || cap == "html" || cap == "media" {
			return true
		}
	}
	return false
}

func (nc NotificationChannel) IsRealTime() bool {
	capabilities := nc.Capabilities()
	for _, cap := range capabilities {
		if cap == "real_time" || cap == "push" {
			return true
		}
	}
	return false
}

func (nc NotificationChannel) CanDeliverUrgent() bool {
	capabilities := nc.Capabilities()
	return slices.Contains(capabilities, "urgent")
}

func (nc NotificationChannel) IsAsync() bool {
	capabilities := nc.Capabilities()
	return slices.Contains(capabilities, "async")
}

func (nc NotificationChannel) RecommendedForType(nt NotificationType) bool {
	switch {
	case nt.IsUrgent():
		return nc.CanDeliverUrgent() && nc.IsRealTime()
	case nt.IsAuthentication():
		return nc == NotificationChannelSMS || nc == NotificationChannelEmail
	case nt.IsAppointmentRelated():
		return nc == NotificationChannelSMS || nc == NotificationChannelPush || nc == NotificationChannelEmail
	case nt.IsFinancial():
		return nc == NotificationChannelEmail // Formal communications
	default:
		return true
	}
}

func GetAllNotificationTypes() []NotificationType {
	return ValidNotificationTypes
}

func GetAllNotificationChannels() []NotificationChannel {
	return ValidNotificationChannels
}

func GetUrgentNotificationTypes() []NotificationType {
	urgentTypes := []NotificationType{}
	for _, nt := range ValidNotificationTypes {
		if nt.IsUrgent() {
			urgentTypes = append(urgentTypes, nt)
		}
	}
	return urgentTypes
}

func GetAuthenticationNotificationTypes() []NotificationType {
	return []NotificationType{
		NotificationTypeVerificationToken,
		NotificationTypeActivationToken,
		NotificationTypePasswordReset,
	}
}

func GetRecommendedChannelsForType(nt NotificationType) []NotificationChannel {
	recommended := []NotificationChannel{}
	for _, nc := range ValidNotificationChannels {
		if nc.RecommendedForType(nt) {
			recommended = append(recommended, nc)
		}
	}
	return recommended
}
