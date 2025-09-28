package facade

import (
	c "clinic-vet-api/app/modules/medical/session/application/command"
	q "clinic-vet-api/app/modules/medical/session/application/query"
	"clinic-vet-api/app/shared/cqrs"
	p "clinic-vet-api/app/shared/page"
	"context"
)

type MedicalSessionCommandBus interface {
	CreateMedicalSession(ctx context.Context, cmd c.CreateMedSessionCommand) cqrs.CommandResult
	UpdateMedicalSession(ctx context.Context, cmd c.UpdateMedSessionCommand) cqrs.CommandResult
	DeleteMedSessionCommand(ctx context.Context, cmd c.DeleteMedSessionCommand) cqrs.CommandResult
}

type MedicalSessionQueryBus interface {
	FindMedSessionByID(ctx context.Context, qry q.FindMedSessionByIDQuery) (*q.MedSessionResult, error)
	FindMedSessionBySpec(ctx context.Context, qry q.FindMedSessionBySpecQuery) (*p.Page[q.MedSessionResult], error)
	FindMedSessionByEmployeeID(ctx context.Context, qry q.FindMedSessionByEmployeeIDQuery) (*p.Page[q.MedSessionResult], error)
	FindMedSessionByPetID(ctx context.Context, qry q.FindMedSessionByPetIDQuery) (*p.Page[q.MedSessionResult], error)
	FindMedSessionByCustomerID(ctx context.Context, qry q.FindMedSessionByCustomerIDQuery) (*p.Page[q.MedSessionResult], error)
	FindRecentMedSessionByPetID(ctx context.Context, qry q.FindRecentMedSessionByPetIDQuery) ([]q.MedSessionResult, error)
	FindMedSessionByDateRange(ctx context.Context, qry q.FindMedSessionByDateRangeQuery) (*p.Page[q.MedSessionResult], error)
	FindMedSessionByPetAndDateRange(ctx context.Context, qry q.FindMedSessionByPetAndDateRangeQuery) ([]q.MedSessionResult, error)
	FindMedSessionByDiagnosis(ctx context.Context, qry q.FindMedSessionByDiagnosisQuery) (*p.Page[q.MedSessionResult], error)
}

type MedicalApplicationService interface {
	CommandBus() MedicalSessionCommandBus
	QueryBus() MedicalSessionQueryBus
}

type medicalApplicationService struct {
	commandHandlers MedicalSessionCommandBus
	queryHandlers   MedicalSessionQueryBus
}

func NewMedicalApplicationService(
	cmdHandlers MedicalSessionCommandBus, qryHandlers MedicalSessionQueryBus,
) MedicalApplicationService {
	return &medicalApplicationService{
		commandHandlers: cmdHandlers,
		queryHandlers:   qryHandlers,
	}
}

func (s *medicalApplicationService) CommandBus() MedicalSessionCommandBus {
	return s.commandHandlers
}

func (s *medicalApplicationService) QueryBus() MedicalSessionQueryBus {
	return s.queryHandlers
}
