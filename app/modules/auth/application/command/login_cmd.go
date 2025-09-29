package command

import apperror "clinic-vet-api/app/shared/error/application"

type LoginCommand struct {
	identifier string
	password   string
	metadata   LoginMetadata
}

type LoginMetadata struct {
	ip         string
	userAgent  string
	deviceInfo string
}

func NewLoginCommand(identifier, password string, metadata LoginMetadata) (LoginCommand, error) {
	cmd := &LoginCommand{
		identifier: identifier,
		password:   password,
		metadata:   metadata,
	}

	if err := cmd.validate(); err != nil {
		return LoginCommand{}, err
	}

	return *cmd, nil
}

func NewLoginMetadata(ip, userAgent, deviceInfo string) *LoginMetadata {
	return &LoginMetadata{
		ip:         ip,
		userAgent:  userAgent,
		deviceInfo: deviceInfo,
	}
}

func (c *LoginCommand) validate() error {
	if c.identifier == "" {
		return LoginCmdError("identifier", "identifier is required")
	}
	if c.password == "" {
		return LoginCmdError("password", "password is required")
	}
	return nil
}

func (c *LoginMetadata) validate() error {
	return nil
}

func (c *LoginCommand) Identifier() string      { return c.identifier }
func (c *LoginCommand) Password() string        { return c.password }
func (c *LoginCommand) Metadata() LoginMetadata { return c.metadata }

func (m *LoginMetadata) IP() string         { return m.ip }
func (m *LoginMetadata) UserAgent() string  { return m.userAgent }
func (m *LoginMetadata) DeviceInfo() string { return m.deviceInfo }

func LoginCmdError(field, issue string) error {
	return apperror.CommandDataValidationError("LoginCommand", field, issue)
}
