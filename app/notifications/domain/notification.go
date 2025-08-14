package notificationDomain

import "time"

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

type Notification struct {
	ID        string              `bson:"_id,omitempty"`
	UserID    string              `bson:"user_id"`
	UserEmail string              `bson:"user_email"`
	UserPhone string              `bson:"user_phone"`
	Title     string              `bson:"title"`
	Subject   string              `bson:"subject"`
	Message   string              `bson:"message"`
	Token     string              `bson:"token"`
	CreatedAt time.Time           `bson:"created_at"`
	NType     NotificationType    `bson:"n_type"`
	Channel   NotificationChannel `bson:"channel"`
}
