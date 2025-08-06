package userCommand

import "time"

type CreateUserCommand struct {
	Email       string
	Password    string
	FirstName   string
	LastName    string
	Gender      string
	Phone       string
	Address     string
	Role        string
	Status      string
	DateOfBirth time.Time
}
