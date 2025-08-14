package notificationDomain

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

// NotificationRepository defines methods for persisting and tracking notifications
//go:generate mockgen -source=repository.go -destination=../../test/mock/notification_repository_mock.go -package=mock

type NotificationRepository interface {
	Create(ctx context.Context, notification *Notification) error
	GetById(ctx context.Context, id string) (Notification, error)
	ListByUser(ctx context.Context, userID string, pagination page.PageData) (page.Page[[]Notification], error)
	ListByType(ctx context.Context, notificationType string, pagination page.PageData) (page.Page[[]Notification], error)
	ListByChannel(ctx context.Context, channel string, pagination page.PageData) (page.Page[[]Notification], error)
}
