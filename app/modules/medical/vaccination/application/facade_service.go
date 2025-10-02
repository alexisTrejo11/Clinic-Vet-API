package application

import (
	"clinic-vet-api/app/modules/core/domain/entity/medical"
	c "clinic-vet-api/app/modules/medical/vaccination/application/command"
	h "clinic-vet-api/app/modules/medical/vaccination/application/handler"
	q "clinic-vet-api/app/modules/medical/vaccination/application/query"
	"clinic-vet-api/app/shared/cqrs"
	p "clinic-vet-api/app/shared/page"
	"context"
)

type VaccinationFacadeService interface {
	FindVaccinationByID(ctx context.Context, qry q.FindVaccinationByIDQuery) (h.VaccinationResult, error)
	FindVaccinationsByPet(ctx context.Context, qry q.FindVaccinationsByPetQuery) (p.Page[h.VaccinationResult], error)
	FindVaccinationsByCustomer(ctx context.Context, qry q.FindVaccinationsByCustomerQuery) (p.Page[h.VaccinationResult], error)
	FindVaccinationsByEmployee(ctx context.Context, qry q.FindVaccinationsByEmployeeQuery) (p.Page[h.VaccinationResult], error)
	FindVaccinationsByDateRange(ctx context.Context, qry q.FindVaccinationsByDateRangeQuery) (p.Page[h.VaccinationResult], error)

	RegisterVaccine(ctx context.Context, cmd c.RegisterVaccinationCommand) cqrs.CommandResult
	UpdateVaccine(ctx context.Context, cmd c.UpdateVaccinationCommand) cqrs.CommandResult
	GenerateVaccinePlan(ctx context.Context, cmd c.GenerateVaccinationPlanCommand) (medical.VaccinationPlan, error)
	DeleteVaccine(ctx context.Context, cmd c.DeleteVaccinationCommand) cqrs.CommandResult
}

type vaccinationFacadeService struct {
	qryHandler *h.VaccinationQueryHandler
	cmdHandler *h.PetVaccineCmdHandler
}

func NewVaccinationFacadeService(qryHandler *h.VaccinationQueryHandler, cmdHandler *h.PetVaccineCmdHandler) VaccinationFacadeService {
	return &vaccinationFacadeService{
		qryHandler: qryHandler,
		cmdHandler: cmdHandler,
	}
}

func (s *vaccinationFacadeService) FindVaccinationByID(ctx context.Context, qry q.FindVaccinationByIDQuery) (h.VaccinationResult, error) {
	return s.qryHandler.HandleVaccinationByID(ctx, qry)
}

func (s *vaccinationFacadeService) FindVaccinationsByPet(ctx context.Context, qry q.FindVaccinationsByPetQuery) (p.Page[h.VaccinationResult], error) {
	return s.qryHandler.HandleVaccinationsByPet(ctx, qry)
}

func (s *vaccinationFacadeService) FindVaccinationsByCustomer(ctx context.Context, qry q.FindVaccinationsByCustomerQuery) (p.Page[h.VaccinationResult], error) {
	return s.qryHandler.HandleVaccinationsByCustomer(ctx, qry)
}

func (s *vaccinationFacadeService) FindVaccinationsByEmployee(ctx context.Context, qry q.FindVaccinationsByEmployeeQuery) (p.Page[h.VaccinationResult], error) {
	return s.qryHandler.HandleVaccinationsByEmployee(ctx, qry)
}

func (s *vaccinationFacadeService) FindVaccinationsByDateRange(ctx context.Context, qry q.FindVaccinationsByDateRangeQuery) (p.Page[h.VaccinationResult], error) {
	return s.qryHandler.HandleVaccinationsByDateRange(ctx, qry)
}

func (s *vaccinationFacadeService) RegisterVaccine(ctx context.Context, cmd c.RegisterVaccinationCommand) cqrs.CommandResult {
	return s.cmdHandler.HandleRegister(ctx, cmd)
}

func (s *vaccinationFacadeService) UpdateVaccine(ctx context.Context, cmd c.UpdateVaccinationCommand) cqrs.CommandResult {
	return s.cmdHandler.HandleUpdate(ctx, cmd)
}

func (s *vaccinationFacadeService) GenerateVaccinePlan(ctx context.Context, cmd c.GenerateVaccinationPlanCommand) (medical.VaccinationPlan, error) {
	return s.cmdHandler.HandleGenerateVaccPlan(ctx, cmd)
}

func (s *vaccinationFacadeService) DeleteVaccine(ctx context.Context, cmd c.DeleteVaccinationCommand) cqrs.CommandResult {
	return s.cmdHandler.HandleDelete(ctx, cmd.VaccinationID)
}
