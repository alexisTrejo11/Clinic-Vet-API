package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net/smtp"
	"strings"

	"clinic-vet-api/app/core/domain/entity/notification"
	emailTemplates "clinic-vet-api/app/modules/notifications/infrastructure/sending/email/templates"
)

func (s *emailSenderImpl) assignTemplate(notification *notification.Notification) (*template.Template, error) {
	templateName := s.getTemplateName(notification)
	tmpl, exists := s.templates[templateName]
	if !exists {
		return nil, fmt.Errorf("template %s not found", templateName)
	}

	return tmpl, nil
}

func (s *emailSenderImpl) renderTemplate(tmpl *template.Template, data EmailTemplateData) (bytes.Buffer, error) {
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return body, fmt.Errorf("error executing template: %w", err)
	}

	return body, nil
}

func (s *emailSenderImpl) loadTemplates() {
	activationTemplate := emailTemplates.ActivationTemplate
	mfaTemplate := emailTemplates.MfaTemplate
	marketingTemplate := emailTemplates.MarketingTemplate

	var err error
	s.templates["activation"], err = template.New("activation").Parse(activationTemplate)
	if err != nil {
		panic(fmt.Sprintf("Error parsing activation template: %v", err))
	}

	s.templates["mfa"], err = template.New("mfa").Parse(mfaTemplate)
	if err != nil {
		panic(fmt.Sprintf("Error parsing MFA template: %v", err))
	}

	s.templates["marketing"], err = template.New("marketing").Parse(marketingTemplate)
	if err != nil {
		panic(fmt.Sprintf("Error parsing marketing template: %v", err))
	}
}

func (s *emailSenderImpl) getTemplateName(notification *notification.Notification) string {
	switch {
	case strings.Contains(strings.ToLower(notification.Subject), "activación") ||
		strings.Contains(strings.ToLower(notification.Subject), "activation"):
		return "activation"
	case strings.Contains(strings.ToLower(notification.Subject), "verificación") ||
		strings.Contains(strings.ToLower(notification.Subject), "mfa") ||
		strings.Contains(strings.ToLower(notification.Subject), "código"):
		return "mfa"
	default:
		return "marketing"
	}
}

func (s *emailSenderImpl) prepareTemplateData(notification *notification.Notification) EmailTemplateData {
	return EmailTemplateData{
		ProjectName: s.config.ProjectName,
		LogoURL:     s.config.LogoURL,
		UserName:    notification.UserEmail,
		Token:       notification.Token,
		Message:     notification.Message,
		Title:       notification.Subject,
		ButtonText:  "", // TODO: Add ButtonText
		ButtonURL:   "",
		Year:        2024,
	}
}

func (s *emailSenderImpl) sendEmail(to, subject, htmlBody string) error {
	// Configuración del servidor SMTP
	auth := smtp.PlainAuth("", s.config.SMTPUsername, s.config.SMTPPassword, s.config.SMTPHost)

	// Crear el mensaje
	headers := make(map[string]string)
	headers["From"] = fmt.Sprintf("%s <%s>", s.config.FromName, s.config.FromEmail)
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	// Construir el mensaje completo
	var message bytes.Buffer
	for k, v := range headers {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	message.WriteString("\r\n")
	message.WriteString(htmlBody)

	tlsconfig := &tls.Config{
		InsecureSkipVerify: false,
		ServerName:         s.config.SMTPHost,
	}

	conn, err := tls.Dial("tcp", s.config.SMTPHost+":"+s.config.SMTPPort, tlsconfig)
	if err != nil {
		return fmt.Errorf("error connecting to SMTP server: %w", err)
	}
	defer conn.Close()

	client, err := smtp.NewClient(conn, s.config.SMTPHost)
	if err != nil {
		return fmt.Errorf("error creating SMTP client: %w", err)
	}
	defer client.Quit()

	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("error authenticating: %w", err)
	}

	if err = client.Mail(s.config.FromEmail); err != nil {
		return fmt.Errorf("error setting sender: %w", err)
	}

	if err = client.Rcpt(to); err != nil {
		return fmt.Errorf("error setting recipient: %w", err)
	}

	writer, err := client.Data()
	if err != nil {
		return fmt.Errorf("error getting data writer: %w", err)
	}

	_, err = writer.Write(message.Bytes())
	if err != nil {
		return fmt.Errorf("error writing message: %w", err)
	}

	err = writer.Close()
	if err != nil {
		return fmt.Errorf("error closing writer: %w", err)
	}

	return nil
}
