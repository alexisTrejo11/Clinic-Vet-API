package email

import (
	"context"
	"fmt"
	"html/template"

	"clinic-vet-api/app/modules/core/domain/entity/notification"
	service "clinic-vet-api/app/modules/notification/application"

	"clinic-vet-api/app/config"
)

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

func NewEmailSender(config config.EmailConfig) service.EmailSender {
	sender := &emailSenderImpl{
		config:    config,
		templates: make(map[string]*template.Template),
	}
	sender.loadTemplates()
	return sender
}

func (s *emailSenderImpl) Send(ctx context.Context, notif *notification.Notification) error {
	fmt.Println("Config:", s.config)
	tmpl, err := s.getTemplate(notif.NType())
	if err != nil {
		return fmt.Errorf("failed to get template: %w", err)
	}

	templateData := s.prepareTemplateData(notif)

	htmlBody, err := s.renderTemplate(tmpl, templateData)
	if err != nil {
		return fmt.Errorf("failed to render template: %w", err)
	}

	if err := s.sendEmail(notif.Email(), notif.Subject(), htmlBody.String()); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
