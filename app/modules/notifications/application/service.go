package service

import (
	"context"
	"fmt"

	"clinic-vet-api/app/modules/core/domain/entity/notification"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/page"
)

type NotificationService interface {
	SendNotification(ctx context.Context, notification *notification.Notification) error
	ListNotificationByUserID(ctx context.Context, userID valueobject.UserID, pagination page.PaginationRequest) (page.Page[NotificationResponse], error)
	ListNotificationBySpecies(ctx context.Context, notificationType string, pagination page.PaginationRequest) (page.Page[NotificationResponse], error)
	ListNotificationByChannel(ctx context.Context, channel string, pagination page.PaginationRequest) (page.Page[NotificationResponse], error)
	GenerateSummary(ctx context.Context, senderID string) ([]notification.Notification, error)
	GetNotificationByID(ctx context.Context, id string) (notification.Notification, error)
}

type notificationServiceImpl struct {
	notificationRepo repository.NotificationRepository
	senders          map[string]Sender
}

func NewNotificationService(notificationRepo repository.NotificationRepository, senders map[string]Sender) NotificationService {
	return &notificationServiceImpl{
		notificationRepo: notificationRepo,
		senders:          senders,
	}
}

func (s *notificationServiceImpl) SendNotification(ctx context.Context, notification *notification.Notification) error {
	typeStr := notification.NType.String()
	if s.senders[typeStr] == nil {
		return fmt.Errorf("no sender registered for notification type: %s", typeStr)
	}
	return s.senders[typeStr].Send(ctx, notification)
}

func (s *notificationServiceImpl) ListNotificationByUserID(ctx context.Context, userID valueobject.UserID, pagination page.PaginationRequest) (page.Page[NotificationResponse], error) {
	notificationPage, err := s.notificationRepo.ListByUser(ctx, userID, pagination)
	if err != nil {
		return page.Page[NotificationResponse]{}, err
	}

	return page.MapItems(notificationPage, fromDomain), nil
}

func (s *notificationServiceImpl) GetNotificationByID(ctx context.Context, id string) (notification.Notification, error) {
	return s.notificationRepo.GetByID(ctx, id)
}

func (s *notificationServiceImpl) ListNotificationBySpecies(ctx context.Context, notificationType string, pagination page.PaginationRequest) (page.Page[NotificationResponse], error) {
	notificationPage, err := s.notificationRepo.ListBySpecies(ctx, notificationType, pagination)
	if err != nil {
		return page.Page[NotificationResponse]{}, err
	}

	return page.MapItems(notificationPage, fromDomain), nil
}

func (s *notificationServiceImpl) ListNotificationByChannel(ctx context.Context, channel string, pagination page.PaginationRequest) (page.Page[NotificationResponse], error) {
	notificationPage, err := s.notificationRepo.ListByChannel(ctx, channel, pagination)
	if err != nil {
		return page.Page[NotificationResponse]{}, err
	}

	return page.MapItems(notificationPage, fromDomain), nil
}

func (s *notificationServiceImpl) GenerateSummary(ctx context.Context, senderID string) ([]notification.Notification, error) {
	return []notification.Notification{}, nil
}
