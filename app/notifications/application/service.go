package application

import (
	"context"
	"fmt"

	domain "github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/domain"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
)

type NotificationFacade interface {
	SendNotification(ctx context.Context, notification *domain.Notification) error
	ListNotificationByUserId(ctx context.Context, userId string, pagination page.PageData) (page.Page[[]NotificationResponse], error)
	ListNotificationByType(ctx context.Context, notificationType string, pagination page.PageData) (page.Page[[]NotificationResponse], error)
	ListNotificationByChannel(ctx context.Context, channel string, pagination page.PageData) (page.Page[[]NotificationResponse], error)
	GenerateSummary(ctx context.Context, senderID string) ([]domain.Notification, error)
	GetNotificationByID(ctx context.Context, id string) (domain.Notification, error)
}

type notificationServiceImpl struct {
	notificationRepo domain.NotificationRepository
	senders          map[string]Sender
}

func NewNotificationService(notificationRepo domain.NotificationRepository, senders map[string]Sender) NotificationFacade {
	return &notificationServiceImpl{
		notificationRepo: notificationRepo,
		senders:          senders,
	}
}

func (s *notificationServiceImpl) SendNotification(ctx context.Context, notification *domain.Notification) error {
	typeStr := notification.NType.String()
	if s.senders[typeStr] == nil {
		return fmt.Errorf("no sender registered for notification type: %s", typeStr)
	}
	return s.senders[typeStr].Send(ctx, notification)
}

func (s *notificationServiceImpl) ListNotificationByUserId(ctx context.Context, userId string, pagination page.PageData) (page.Page[[]NotificationResponse], error) {
	notificationPage, err := s.notificationRepo.ListByUser(ctx, userId, pagination)
	if err != nil {
		return page.Page[[]NotificationResponse]{}, err
	}

	responsePage := page.NewPage(fromDomainList(notificationPage.Data), notificationPage.Metadata)
	return responsePage, nil
}

func (s *notificationServiceImpl) GetNotificationByID(ctx context.Context, id string) (domain.Notification, error) {
	return s.notificationRepo.GetById(ctx, id)
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

func (s *notificationServiceImpl) GenerateSummary(ctx context.Context, senderID string) ([]domain.Notification, error) {
	// return s.notificationRepo.GetBySender(ctx, senderID)
	return []domain.Notification{}, nil
}
