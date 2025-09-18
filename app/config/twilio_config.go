package config

import (
	"os"

	"github.com/twilio/twilio-go"
)

type TwilioConfig struct {
	AccountSID string `json:"account_sid"`
	AuthToken  string `json:"auth_token"`
	FromPhone  string `json:"from_phone"`
}

var twilioClient *twilio.RestClient

func InitTwilio() {
	accountSid := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	if accountSid == "" || authToken == "" {
		panic("Twilio credentials are not set in environment variables")
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})

	twilioClient = client
}

func GetTwilioClient() *twilio.RestClient {
	if twilioClient == nil {
		InitTwilio()
	}

	return twilioClient
}
