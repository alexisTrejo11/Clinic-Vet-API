package command

type RefreshSessionCommand struct {
	refreshToken string
}

func NewRefreshSessionCommand(refreshToken string) RefreshSessionCommand {
	return RefreshSessionCommand{
		refreshToken: refreshToken,
	}
}

func (c RefreshSessionCommand) RefreshToken() string {
	return c.refreshToken
}
