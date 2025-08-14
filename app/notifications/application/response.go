package application

import (
	"time"

	domain "github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/domain"
)

type NotificationResponse struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	UserEmail string `json:"user_email"`
	UserPhone string `json:"user_phone"`
	Title     string `json:"title"`
	Subject   string `json:"subject"`
	Message   string `json:"message"`
	Token     string `json:"token"`
	CreatedAt string `json:"created_at"`
	NType     string `json:"type"`
	Channel   string `json:"channel"`
}

func fromDomain(notification domain.Notification) NotificationResponse {
	response := &NotificationResponse{
		ID:        notification.ID,
		UserID:    notification.UserID,
		UserEmail: notification.UserEmail,
		UserPhone: notification.UserPhone,
		Title:     notification.Title,
		Subject:   notification.Subject,
		Message:   notification.Message,
		Token:     notification.Token,
		CreatedAt: notification.CreatedAt.Format(time.RFC3339),
		NType:     notification.NType.String(),
		Channel:   notification.Channel.String(),
	}
	return *response
}

func fromDomainList(notifications []domain.Notification) []NotificationResponse {
	responses := make([]NotificationResponse, 0, len(notifications))
	for i, notification := range notifications {
		responses[i] = fromDomain(notification)
	}
	return responses
}
