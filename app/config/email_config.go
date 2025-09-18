package config

import "os"

type EmailConfig struct {
	SMTPHost     string
	SMTPPort     string
	SMTPUsername string
	SMTPPassword string
	FromEmail    string
	FromName     string
	ProjectName  string
	LogoURL      string
}

func InitEmailConfig() EmailConfig {
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")
	smtpUsername := os.Getenv("SMTP_USERNAME")
	smtpPassword := os.Getenv("SMTP_PASSWORD")
	fromEmail := os.Getenv("FROM_EMAIL")
	fromName := os.Getenv("FROM_NAME")
	projectName := os.Getenv("PROJECT_NAME")
	logoURL := os.Getenv("LOGO_URL")

	if smtpHost == "" && smtpPort == "" && smtpUsername == "" && smtpPassword == "" && fromEmail == "" && fromName == "" && projectName == "" && logoURL == "" {
		panic("One or more email configuration variables are not set")
	}

	return EmailConfig{
		SMTPHost:     smtpHost,
		SMTPPort:     smtpPort,
		SMTPUsername: smtpUsername,
		SMTPPassword: smtpPassword,
		FromEmail:    fromEmail,
		FromName:     fromName,
		ProjectName:  projectName,
		LogoURL:      logoURL,
	}
}
