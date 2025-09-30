package service

import (
	"context"

	"clinic-vet-api/app/modules/core/domain/entity/notification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/page"
)

type NotificationQueryService interface {
	ListNotificationByUserID(ctx context.Context, userID valueobject.UserID, pagination page.PaginationRequest) (page.Page[NotificationResponse], error)
	ListNotificationBySpecies(ctx context.Context, notificationType string, pagination page.PaginationRequest) (page.Page[NotificationResponse], error)
	ListNotificationByChannel(ctx context.Context, channel string, pagination page.PaginationRequest) (page.Page[NotificationResponse], error)
	GenerateSummary(ctx context.Context, senderID string) ([]notification.Notification, error)
	GetNotificationByID(ctx context.Context, id string) (notification.Notification, error)
}

type notificationQueryServiceImpl struct {
	notificationRepo repository.NotificationRepository
}

func NewNotificationQueryService(notificationRepo repository.NotificationRepository) NotificationQueryService {
	return &notificationQueryServiceImpl{
		notificationRepo: notificationRepo,
	}
}

func (s *notificationQueryServiceImpl) ListNotificationByUserID(ctx context.Context, userID valueobject.UserID, pagination page.PaginationRequest) (page.Page[NotificationResponse], error) {
	notificationPage, err := s.notificationRepo.ListByUser(ctx, userID, pagination)
	if err != nil {
		return page.Page[NotificationResponse]{}, err
	}

	return page.MapItems(notificationPage, fromDomain), nil
}

func (s *notificationQueryServiceImpl) GetNotificationByID(ctx context.Context, id string) (notification.Notification, error) {
	return s.notificationRepo.GetByID(ctx, id)
}

func (s *notificationQueryServiceImpl) ListNotificationBySpecies(ctx context.Context, notificationType string, pagination page.PaginationRequest) (page.Page[NotificationResponse], error) {
	notificationPage, err := s.notificationRepo.ListBySpecies(ctx, notificationType, pagination)
	if err != nil {
		return page.Page[NotificationResponse]{}, err
	}

	return page.MapItems(notificationPage, fromDomain), nil
}

func (s *notificationQueryServiceImpl) ListNotificationByChannel(ctx context.Context, channel string, pagination page.PaginationRequest) (page.Page[NotificationResponse], error) {
	notificationPage, err := s.notificationRepo.ListByChannel(ctx, channel, pagination)
	if err != nil {
		return page.Page[NotificationResponse]{}, err
	}

	return page.MapItems(notificationPage, fromDomain), nil
}

func (s *notificationQueryServiceImpl) GenerateSummary(ctx context.Context, senderID string) ([]notification.Notification, error) {
	return []notification.Notification{}, nil
}
