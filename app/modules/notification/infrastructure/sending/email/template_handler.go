package email

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"html/template"
	"net"
	"net/smtp"
	"time"

	"clinic-vet-api/app/modules/core/domain/entity/notification"
	"clinic-vet-api/app/modules/core/domain/enum"
	emailTemplates "clinic-vet-api/app/modules/notification/infrastructure/sending/email/templates"
)

func (s *emailSenderImpl) renderTemplate(tmpl *template.Template, data EmailTemplateData) (bytes.Buffer, error) {
	var body bytes.Buffer
	if err := tmpl.Execute(&body, data); err != nil {
		return body, fmt.Errorf("error executing template: %w", err)
	}

	return body, nil
}

func (s *emailSenderImpl) loadTemplates() {
	templates := map[enum.NotificationType]string{
		enum.NotificationTypeActivationToken: emailTemplates.ActivationTemplate,
		enum.NotificationTypeMFA:             emailTemplates.MfaTemplate,
		enum.NotificationTypeInfo:            emailTemplates.MarketingTemplate,
		enum.NotificationTypeWelcome:         emailTemplates.MarketingTemplate,
		enum.NotificationTypePasswordReset:   emailTemplates.ActivationTemplate,
	}

	for notifType, tmplString := range templates {
		templateName := string(notifType)
		tmpl, err := template.New(templateName).Parse(tmplString)
		if err != nil {
			panic(fmt.Sprintf("Error parsing %s template: %v", notifType, err))
		}
		s.templates[templateName] = tmpl
	}
}

func (s *emailSenderImpl) getTemplate(notifType enum.NotificationType) (*template.Template, error) {
	tmpl, exists := s.templates[notifType.String()]
	if !exists {
		return nil, fmt.Errorf("template not found for notification type: %s", notifType)
	}
	return tmpl, nil
}

func (s *emailSenderImpl) prepareTemplateData(notif *notification.Notification) EmailTemplateData {
	data := EmailTemplateData{
		ProjectName: s.config.ProjectName,
		LogoURL:     s.config.LogoURL,
		UserName:    notif.Email(),
		Token:       notif.Token(),
		Message:     notif.Message(),
		Title:       notif.Subject(),
		Year:        time.Now().Year(),
	}

	switch notif.NType() {
	case enum.NotificationTypeActivationToken:
		data.ButtonText = "Activar Cuenta"
		data.ButtonURL = fmt.Sprintf("https://localhost:8080/api/v2/auth/activate?token=%s", notif.Token())

	case enum.NotificationTypePasswordReset:
		data.ButtonText = "Restablecer Contraseña"
		data.ButtonURL = fmt.Sprintf("https://localhost:8080/api/v2/auth/reset-password?token=%s", notif.Token())

	case enum.NotificationTypeMFA:
		data.ButtonText = ""
		if notif.Message() == "" {
			data.Message = fmt.Sprintf("Tu código de verificación es: %s", notif.Token())
		}

	case enum.NotificationTypeWelcome:
		data.ButtonText = "Comenzar"
		data.ButtonURL = "https://localhost:8080/api/v2/welcome"
	}

	return data

}

func (s *emailSenderImpl) sendEmail(to, subject, htmlBody string) error {
	auth := smtp.PlainAuth("", s.config.SMTPUsername, s.config.SMTPPassword, s.config.SMTPHost)

	// Construir headers
	headers := make(map[string]string)
	headers["From"] = fmt.Sprintf("%s <%s>", s.config.FromName, s.config.FromEmail)
	headers["To"] = to
	headers["Subject"] = subject
	headers["MIME-Version"] = "1.0"
	headers["Content-Type"] = "text/html; charset=UTF-8"

	var message bytes.Buffer
	for k, v := range headers {
		message.WriteString(fmt.Sprintf("%s: %s\r\n", k, v))
	}
	message.WriteString("\r\n")
	message.WriteString(htmlBody)

	serverAddr := fmt.Sprintf("%s:%d", s.config.SMTPHost, s.config.SMTPPort)

	// 1. Conectar sin TLS
	conn, err := net.Dial("tcp", serverAddr)
	if err != nil {
		return fmt.Errorf("error connecting to SMTP server: %w", err)
	}
	defer conn.Close()

	// 2. Crear cliente SMTP
	client, err := smtp.NewClient(conn, s.config.SMTPHost)
	if err != nil {
		return fmt.Errorf("error creating SMTP client: %w", err)
	}
	defer client.Quit()

	// ✅ 3. Verificar soporte de STARTTLS
	if ok, _ := client.Extension("STARTTLS"); ok {
		tlsconfig := &tls.Config{
			ServerName:         s.config.SMTPHost,
			InsecureSkipVerify: false,
		}

		if err = client.StartTLS(tlsconfig); err != nil {
			return fmt.Errorf("error starting TLS: %w", err)
		}
	} else {
		return fmt.Errorf("server does not support STARTTLS")
	}

	// 4. Autenticar (ahora la conexión está encriptada)
	if err = client.Auth(auth); err != nil {
		return fmt.Errorf("error authenticating: %w", err)
	}

	// 5. Enviar email
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
