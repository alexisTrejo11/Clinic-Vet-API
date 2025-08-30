package email

import (
	"context"
	"fmt"
	"html/template"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	service "github.com/alexisTrejo11/Clinic-Vet-API/app/modules/notifications/application"

	"github.com/alexisTrejo11/Clinic-Vet-API/config"
)

type EmailSender interface {
	Send(ctx context.Context, notification *entity.Notification) error
}
type EmailTemplateData struct {
	ProjectName string
	LogoURL     string
	UserName    string
	Token       string
	Message     string
	Title       string
	ButtonText  string
	ButtonURL   string
	Year        int
}

type emailSenderImpl struct {
	config    config.EmailConfig
	templates map[string]*template.Template
}

func NewEmailSender(config config.EmailConfig) service.Sender {
	sender := &emailSenderImpl{
		config:    config,
		templates: make(map[string]*template.Template),
	}
	sender.loadTemplates()
	return sender
}

func (s *emailSenderImpl) Send(ctx context.Context, notification *entity.Notification) error {
	tmpl, err := s.assignTemplate(notification)
	if err != nil {
		return fmt.Errorf("error assigning template: %w", err)
	}

	data := s.prepareTemplateData(notification)

	body, err := s.renderTemplate(tmpl, data)
	if err != nil {
		return fmt.Errorf("error rendering template: %w", err)
	}

	// Enviar el email
	return s.sendEmail(notification.UserEmail, notification.Subject, body.String())
}
