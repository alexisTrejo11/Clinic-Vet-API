package sms

import (
	"context"
	"encoding/json"
	"fmt"

	"clinic-vet-api/app/modules/core/domain/entity/notification"
	service "clinic-vet-api/app/modules/notifications/application"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioPhoneSender struct {
	client            *twilio.RestClient
	twilioPhoneNumber string
}

func NewTwilioPhoneSender(client *twilio.RestClient, twilioPhoneNumber string) service.SMSender {
	return &TwilioPhoneSender{client: client, twilioPhoneNumber: twilioPhoneNumber}
}

func (s *TwilioPhoneSender) Send(ctx context.Context, notification *notification.Notification) error {
	params := &twilioApi.CreateMessageParams{}

	if notification.Phone() == "" {
		return fmt.Errorf("user phone number is required")
	}

	if s.twilioPhoneNumber == "" {
		return fmt.Errorf("twilio phone number is required")
	}

	params.SetTo(notification.Phone())
	params.SetFrom(s.twilioPhoneNumber)
	params.SetBody(s.GenerateBody(notification))

	resp, err := s.client.Api.CreateMessage(params)
	if err != nil {
		fmt.Println("Error sending SMS message: " + err.Error())
		return err
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
		return nil
	}
}

func (s *TwilioPhoneSender) GenerateBody(notification *notification.Notification) string {
	return fmt.Sprintf("Hello from Clinic Vet! %s", notification.Message)
}
