package command

import apperror "clinic-vet-api/app/shared/error/application"

type OAuthLoginCommand struct {
	provider string
	token    string
	metadata LoginMetadata
}

func NewOAuthLoginCommand(provider, token string, metadata LoginMetadata) (OAuthLoginCommand, error) {
	cmd := &OAuthLoginCommand{provider: provider, token: token, metadata: metadata}
	if err := cmd.validate(); err != nil {
		return OAuthLoginCommand{}, err
	}
	return *cmd, nil
}

func (c *OAuthLoginCommand) validate() error {
	if c.provider == "" {
		return OAuthLoginCmdError("provider", "is required")
	}
	if c.token == "" {
		return OAuthLoginCmdError("token", "is required")
	}

	return nil
}

func (c *OAuthLoginCommand) Provider() string        { return c.provider }
func (c *OAuthLoginCommand) Token() string           { return c.token }
func (c *OAuthLoginCommand) Metadata() LoginMetadata { return c.metadata }

func OAuthLoginCmdError(field, issue string) error {
	return apperror.CommandDataValidationError("OAuthLoginCommand", field, issue)
}
