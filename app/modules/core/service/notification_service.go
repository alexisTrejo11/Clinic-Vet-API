package service

import (
	"clinic-vet-api/app/modules/core/domain/entity/notification"
	"context"
)

type NotificationService interface {
	Send(ctx context.Context, notif *notification.Notification) error
}
