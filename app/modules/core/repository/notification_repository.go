package repository

import (
	"context"

	"clinic-vet-api/app/modules/core/domain/entity/notification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/shared/page"
)

// entity.NotificationRepository defines methods for persisting and tracking notifications
//go:generate mockgen -source=repository.go -destination=../../test/mock/notification_repository_mock.go -package=mock

type NotificationRepository interface {
	Create(ctx context.Context, notification *notification.Notification) error
	GetByID(ctx context.Context, id string) (notification.Notification, error)
	ListByUser(ctx context.Context, userID valueobject.UserID, pagination page.PaginationRequest) (page.Page[notification.Notification], error)
	ListBySpecies(ctx context.Context, notificationType string, pagination page.PaginationRequest) (page.Page[notification.Notification], error)
	ListByChannel(ctx context.Context, channel string, pagination page.PaginationRequest) (page.Page[notification.Notification], error)
}
