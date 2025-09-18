package config

import (
	"github.com/twilio/twilio-go"
)

type TwilioConfig struct {
	AccountSID string `json:"account_sid"`
	AuthToken  string `json:"auth_token"`
	FromPhone  string `json:"from_phone"`
}

var twilioClient *twilio.RestClient

func InitTwilio(config TwilioConfig) {
	if config.AccountSID == "" || config.AuthToken == "" {
		panic("Twilio credentials are not set in environment variables")
	}

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: config.AccountSID,
		Password: config.AuthToken,
	})

	twilioClient = client
}

func GetTwilioClient() *twilio.RestClient {
	if twilioClient == nil {
		panic("Twilio client is not initialized. Call InitTwilio first.")
	}

	return twilioClient
}
