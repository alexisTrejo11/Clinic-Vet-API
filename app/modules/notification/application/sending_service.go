package service

import (
	"clinic-vet-api/app/modules/core/domain/entity/notification"
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/modules/core/service"
	"context"
	"fmt"
)

type EmailSender interface {
	Send(ctx context.Context, notification *notification.Notification) error
}

type SMSender interface {
	Send(ctx context.Context, notification *notification.Notification) error
}

type notificationService struct {
	emailSender EmailSender
	smsSender   SMSender
	notiRepo    repository.NotificationRepository
}

func NewNotificationService(emailSender EmailSender, smsSender SMSender, notiRepo repository.NotificationRepository) service.NotificationService {
	return &notificationService{
		emailSender: emailSender,
		smsSender:   smsSender,
		notiRepo:    notiRepo,
	}
}

func (ns *notificationService) Send(ctx context.Context, notif *notification.Notification) error {
	switch notif.Channel() {
	case enum.NotificationChannelEmail:
		return ns.emailSender.Send(ctx, notif)
	case enum.NotificationChannelSMS:
		return ns.smsSender.Send(ctx, notif)
	default:
		return fmt.Errorf("unsupported notification channel: %s", notif.Channel())
	}
}
