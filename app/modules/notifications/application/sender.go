package service

import (
	"context"
	"fmt"

	"clinic-vet-api/app/core/domain/entity/notification"
)

type Notifier struct {
	sender Sender
}

type Sender interface {
	Send(ctx context.Context, notification *notification.Notification) error
}

func (n *Notifier) SetSender(sender Sender) {
	n.sender = sender
}

func (n *Notifier) SendNotification(ctx context.Context, notification *notification.Notification) error {
	if n.sender == nil {
		return fmt.Errorf("no sender strategy set")
	}
	return n.sender.Send(ctx, notification)
}

type EmailSender interface {
	Send(ctx context.Context, notification *notification.Notification) error
}
