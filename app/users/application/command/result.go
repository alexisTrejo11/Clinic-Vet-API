package userCommand

type CommandResult struct {
	Success bool
	Id      string
	Message string
	Error   error
}
