package application

import (
	"context"
	"fmt"

	domain "github.com/alexisTrejo11/Clinic-Vet-API/app/notifications/domain"
)

type Notifier struct {
	sender Sender
}

type Sender interface {
	Send(ctx context.Context, notification *domain.Notification) error
}

func (n *Notifier) SetSender(sender Sender) {
	n.sender = sender
}

func (n *Notifier) SendNotification(ctx context.Context, notification *domain.Notification) error {
	if n.sender == nil {
		return fmt.Errorf("no sender strategy set")
	}
	return n.sender.Send(ctx, notification)
}

type EmailSender interface {
	Send(ctx context.Context, notification *domain.Notification) error
}
