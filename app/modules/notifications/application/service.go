package service

import (
	"context"
	"fmt"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type NotificationService interface {
	SendNotification(ctx context.Context, notification *entity.Notification) error
	ListNotificationByUserID(ctx context.Context, userID string, pagination page.PageData) (page.Page[[]NotificationResponse], error)
	ListNotificationByType(ctx context.Context, notificationType string, pagination page.PageData) (page.Page[[]NotificationResponse], error)
	ListNotificationByChannel(ctx context.Context, channel string, pagination page.PageData) (page.Page[[]NotificationResponse], error)
	GenerateSummary(ctx context.Context, senderID string) ([]entity.Notification, error)
	GetNotificationByID(ctx context.Context, id string) (entity.Notification, error)
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

func (s *notificationServiceImpl) SendNotification(ctx context.Context, notification *entity.Notification) error {
	typeStr := notification.NType.String()
	if s.senders[typeStr] == nil {
		return fmt.Errorf("no sender registered for notification type: %s", typeStr)
	}
	return s.senders[typeStr].Send(ctx, notification)
}

func (s *notificationServiceImpl) ListNotificationByUserID(ctx context.Context, userID string, pagination page.PageData) (page.Page[[]NotificationResponse], error) {
	notificationPage, err := s.notificationRepo.ListByUser(ctx, userID, pagination)
	if err != nil {
		return page.Page[[]NotificationResponse]{}, err
	}

	responsePage := page.NewPage(fromDomainList(notificationPage.Data), notificationPage.Metadata)
	return responsePage, nil
}

func (s *notificationServiceImpl) GetNotificationByID(ctx context.Context, id string) (entity.Notification, error) {
	return s.notificationRepo.GetByID(ctx, id)
}

func (s *notificationServiceImpl) ListNotificationByType(ctx context.Context, notificationType string, pagination page.PageData) (page.Page[[]NotificationResponse], error) {
	notificationPage, err := s.notificationRepo.ListByType(ctx, notificationType, pagination)
	if err != nil {
		return page.Page[[]NotificationResponse]{}, err
	}

	responsePage := page.NewPage(fromDomainList(notificationPage.Data), notificationPage.Metadata)
	return responsePage, nil
}

func (s *notificationServiceImpl) ListNotificationByChannel(ctx context.Context, channel string, pagination page.PageData) (page.Page[[]NotificationResponse], error) {
	notificationPage, err := s.notificationRepo.ListByChannel(ctx, channel, pagination)
	if err != nil {
		return page.Page[[]NotificationResponse]{}, err
	}

	responsePage := page.NewPage(fromDomainList(notificationPage.Data), notificationPage.Metadata)
	return responsePage, nil
}

func (s *notificationServiceImpl) GenerateSummary(ctx context.Context, senderID string) ([]entity.Notification, error) {
	// return s.notificationRepo.GetBySender(ctx, senderID)
	return []entity.Notification{}, nil
}
