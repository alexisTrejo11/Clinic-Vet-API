package enum

type NotificationType int

const (
	Alert NotificationType = iota
	Reminder
	Info
	VerificationToken
	ActivationToken
	Unknown
)

func (n NotificationType) String() string {
	switch n {
	case Alert:
		return "alert"
	case Reminder:
		return "reminder"
	case Info:
		return "info"
	case VerificationToken:
		return "verification_token"
	case ActivationToken:
		return "activation_token"
	default:
		return "unknown"
	}
}

type NotificationChannel int

const (
	SMS NotificationChannel = iota
	Email
	Push
)

func (n NotificationChannel) String() string {
	switch n {
	case SMS:
		return "sms"
	case Email:
		return "email"
	case Push:
		return "push"
	default:
		return "unknown"
	}
}
