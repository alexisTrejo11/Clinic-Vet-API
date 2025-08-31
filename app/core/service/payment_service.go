package service

import (
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/enum"
	domainerr "github.com/alexisTrejo11/Clinic-Vet-API/app/core/errors"
)

type PaymentProccesorService struct{}

func (ps *PaymentProccesorService) Cancel(payment *entity.Payment) error {
	if payment.GetStatus() == enum.CANCELLED {
		return domainerr.PaymentStatusConflict(payment.GetID(), errors.New("P"))
	}

	payment.SetStatus(enum.CANCELLED)

	return nil
}

func (ps *PaymentProccesorService) ValidateDelete(payment *entity.Payment) error {
	return nil
}

func (ps *PaymentProccesorService) Process(payment *entity.Payment, transactionID string) error {
	return nil
}

func (ps *PaymentProccesorService) Overdue(payment *entity.Payment) error {
	return nil
}

func (ps *PaymentProccesorService) Refund(payment *entity.Payment) error {
	return nil
}
