package repository

import (
	"context"

	"example.com/at/backend/api-vet/sqlc"
)

type ReminderRepository interface {
	CreateReminder(params sqlc.CreateReminderParams) error
	GetReminderByID(ReminderId int32) (*sqlc.Reminder, error)
	UpdateReminder(updateParams sqlc.UpdateReminderParams) error
	DeleteReminder(ReminderId int32) error
}

type reminderRepositoryImpl struct {
	queries *sqlc.Queries
}

func NewReminderRepository(queries *sqlc.Queries) ReminderRepository {
	return &reminderRepositoryImpl{
		queries: queries,
	}
}

func (ar reminderRepositoryImpl) CreateReminder(params sqlc.CreateReminderParams) error {
	if _, err := ar.queries.CreateReminder(context.Background(), params); err != nil {
		return err
	}

	return nil
}

func (ar reminderRepositoryImpl) GetReminderByID(reminderId int32) (*sqlc.Reminder, error) {
	Reminder, err := ar.queries.GetReminderByID(context.Background(), reminderId)
	if err != nil {
		return nil, err
	}

	return &Reminder, nil
}

func (ar reminderRepositoryImpl) UpdateReminder(updateParams sqlc.UpdateReminderParams) error {
	err := ar.queries.UpdateReminder(context.Background(), updateParams)
	if err != nil {
		return err
	}

	return nil
}

func (ar reminderRepositoryImpl) DeleteReminder(reminderId int32) error {
	if err := ar.queries.DeleteReminder(context.Background(), reminderId); err != nil {
		return err
	}
	return nil
}
