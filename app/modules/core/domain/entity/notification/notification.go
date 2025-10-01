package notification

import (
	"clinic-vet-api/app/modules/core/domain/enum"
	"clinic-vet-api/app/modules/core/domain/valueobject"
	"fmt"
	"time"
)

type Notification struct {
	id        string                   `bson:"_id,omitempty"`
	userID    valueobject.UserID       `bson:"user_id"`
	email     string                   `bson:"email"`
	phone     string                   `bson:"phone"`
	title     string                   `bson:"title"`
	subject   string                   `bson:"subject"`
	message   string                   `bson:"message"`
	token     string                   `bson:"token"`
	createdAt time.Time                `bson:"created_at"`
	nType     enum.NotificationType    `bson:"n_type"`
	channel   enum.NotificationChannel `bson:"channel"`
}

func (n *Notification) ID() string                        { return n.id }
func (n *Notification) UserID() valueobject.UserID        { return n.userID }
func (n *Notification) Email() string                     { return n.email }
func (n *Notification) Phone() string                     { return n.phone }
func (n *Notification) Title() string                     { return n.title }
func (n *Notification) Subject() string                   { return n.subject }
func (n *Notification) Message() string                   { return n.message }
func (n *Notification) Token() string                     { return n.token }
func (n *Notification) CreatedAt() time.Time              { return n.createdAt }
func (n *Notification) NType() enum.NotificationType      { return n.nType }
func (n *Notification) Channel() enum.NotificationChannel { return n.channel }

type NotificationBuilder struct {
	notification *Notification
}

func NewNotificationBuilder() *NotificationBuilder {
	return &NotificationBuilder{notification: &Notification{}}
}

func (b *NotificationBuilder) WithID(id string) *NotificationBuilder {
	b.notification.id = id
	return b
}

func (b *NotificationBuilder) WithUserID(userID valueobject.UserID) *NotificationBuilder {
	b.notification.userID = userID
	return b
}

func (b *NotificationBuilder) WithEmail(userEmail string) *NotificationBuilder {
	b.notification.email = userEmail
	return b
}

func (b *NotificationBuilder) WithPhone(userPhone string) *NotificationBuilder {
	b.notification.phone = userPhone
	return b
}

func (b *NotificationBuilder) WithTitle(title string) *NotificationBuilder {
	b.notification.title = title
	return b
}

func (b *NotificationBuilder) WithSubject(subject string) *NotificationBuilder {
	b.notification.subject = subject
	return b
}

func (b *NotificationBuilder) WithMessage(message string) *NotificationBuilder {
	b.notification.message = message
	return b
}

func (b *NotificationBuilder) WithToken(token string) *NotificationBuilder {
	b.notification.token = token
	return b
}

func (b *NotificationBuilder) WithCreatedAt(createdAt time.Time) *NotificationBuilder {
	b.notification.createdAt = createdAt
	return b
}

func (b *NotificationBuilder) WithNType(nType enum.NotificationType) *NotificationBuilder {
	b.notification.nType = nType
	return b
}

func (b *NotificationBuilder) WithChannel(channel enum.NotificationChannel) *NotificationBuilder {
	b.notification.channel = channel
	return b
}

func (b *NotificationBuilder) Build() *Notification {
	if b.notification.createdAt.IsZero() {
		b.notification.createdAt = time.Now()
	}
	return b.notification
}

func NewActivationEmail(userID valueobject.UserID, email, name, token string) *Notification {
	notif := NewNotificationBuilder().
		WithUserID(userID).
		WithEmail(email).
		WithNType(enum.NotificationTypeActivationToken).
		WithChannel(enum.NotificationChannelEmail).
		WithSubject("Activa tu cuenta").
		WithMessage(fmt.Sprintf("Hola %s, activa tu cuenta usando el enlace.", name)).
		WithToken(token).
		Build()

	return notif
}

func NewWelcomeEmail(email, name string) *Notification {
	notif := NewNotificationBuilder().
		WithEmail(email).
		WithNType(enum.NotificationTypeWelcome).
		WithChannel(enum.NotificationChannelEmail).
		WithSubject("Bienvenido a Clinic Vet").
		WithMessage(fmt.Sprintf("Hola %s, bienvenido a Clinic Vet.", name)).
		Build()

	return notif
}

func NewEnable2FANotificationEmail(userID valueobject.UserID, email valueobject.Email, token string) *Notification {
	title := "Habilitación de 2FA - Método: Email"
	message := fmt.Sprintf("Has solicitado habilitar la autenticación de dos factores (2FA) usando %s. Usa el siguiente token para completar la configuración: %s", "Email", token)

	notif := NewNotificationBuilder().
		WithUserID(userID).
		WithEmail(email.String()).
		WithNType(enum.NotificationTypeVerificationToken).
		WithChannel(enum.NotificationChannelEmail).
		WithTitle(title).
		WithSubject("Habilitación de 2FA").
		WithMessage(message).
		WithToken(token).
		Build()

	return notif
}

func NewEnable2FANotificationSMS(userID valueobject.UserID, phone valueobject.PhoneNumber, token string) *Notification {
	title := "Habilitación de 2FA - Método: SMS"
	message := fmt.Sprintf("Has solicitado habilitar la autenticación de dos factores (2FA) usando %s. Usa el siguiente token para completar la configuración: %s", "SMS", token)

	notif := NewNotificationBuilder().
		WithUserID(userID).
		WithPhone(phone.Value).
		WithNType(enum.NotificationTypeVerificationToken).
		WithChannel(enum.NotificationChannelSMS).
		WithTitle(title).
		WithSubject("Habilitación de 2FA").
		WithMessage(message).
		WithToken(token).
		Build()

	return notif
}

func (b *Notification) SetID(id string) {
	b.id = id
}
