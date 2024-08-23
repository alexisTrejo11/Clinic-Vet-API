package services

type ReminderService interface {
}

type reminderServiceImpl struct {
}

func NewReminderService() ReminderService {
	return &reminderServiceImpl{}
}
