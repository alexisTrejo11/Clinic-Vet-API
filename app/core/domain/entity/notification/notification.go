package notification

import (
	"time"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/enum"
)

type Notification struct {
	ID        string                   `bson:"_id,omitempty"`
	UserID    string                   `bson:"user_id"`
	UserEmail string                   `bson:"user_email"`
	UserPhone string                   `bson:"user_phone"`
	Title     string                   `bson:"title"`
	Subject   string                   `bson:"subject"`
	Message   string                   `bson:"message"`
	Token     string                   `bson:"token"`
	CreatedAt time.Time                `bson:"created_at"`
	NType     enum.NotificationType    `bson:"n_type"`
	Channel   enum.NotificationChannel `bson:"channel"`
}

func NewActivateAccountNotification(userID, userEmail string) Notification {
	return Notification{
		UserID:    userID,
		UserEmail: userEmail,
		Subject:   "Activate your account",
		Message:   "Please click the link below to activate your account.",
		NType:     enum.NotificationTypeActivationToken,
		Channel:   enum.NotificationChannelEmail,
	}
}
