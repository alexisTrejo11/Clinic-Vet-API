package command

import "time"

type CreateProfileCommand struct {
	firstName   string
	lastName    string
	gender      string
	profilePic  string
	bio         string
	dateOfBirth time.Time
	address     string
}
