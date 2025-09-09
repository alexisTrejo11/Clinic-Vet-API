package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/notification"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

// entity.NotificationRepository defines methods for persisting and tracking notifications
//go:generate mockgen -source=repository.go -destination=../../test/mock/notification_repository_mock.go -package=mock

type NotificationRepository interface {
	Create(ctx context.Context, notification *notification.Notification) error
	GetByID(ctx context.Context, id string) (notification.Notification, error)
	ListByUser(ctx context.Context, userID valueobject.UserID, pagination page.PageInput) (page.Page[[]notification.Notification], error)
	ListByType(ctx context.Context, notificationType string, pagination page.PageInput) (page.Page[[]notification.Notification], error)
	ListByChannel(ctx context.Context, channel string, pagination page.PageInput) (page.Page[[]notification.Notification], error)
}
