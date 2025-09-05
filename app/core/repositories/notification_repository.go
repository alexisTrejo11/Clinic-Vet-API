package repository

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

// entity.NotificationRepository defines methods for persisting and tracking notifications
//go:generate mockgen -source=repository.go -destination=../../test/mock/notification_repository_mock.go -package=mock

type NotificationRepository interface {
	Create(ctx context.Context, notification *entity.Notification) error
	GetByID(ctx context.Context, id string) (entity.Notification, error)
	ListByUser(ctx context.Context, userID string, pagination page.PageInput) (page.Page[[]entity.Notification], error)
	ListByType(ctx context.Context, notificationType string, pagination page.PageInput) (page.Page[[]entity.Notification], error)
	ListByChannel(ctx context.Context, channel string, pagination page.PageInput) (page.Page[[]entity.Notification], error)
}
