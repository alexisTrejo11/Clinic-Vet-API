package application

import (
	"context"

	"clinic-vet-api/app/modules/medical/deworm/application/command"
	"clinic-vet-api/app/modules/medical/deworm/application/query"
	"clinic-vet-api/app/shared/cqrs"
	p "clinic-vet-api/app/shared/page"
)

type DewormingFacadeService interface {
	RegisterDeworm(ctx context.Context, cmd command.DewormCreateCommand) cqrs.CommandResult
	UpdateDeworm(ctx context.Context, cmd command.DewormUpdateCommand) cqrs.CommandResult
	DeleteDeworm(ctx context.Context, cmd command.DewormDeleteCommand) cqrs.CommandResult

	FindDewormByID(ctx context.Context, qry query.FindDewormByIDQuery) (query.DewormResult, error)
	FindDewormsByPet(ctx context.Context, qry query.FindDewormsByPetQuery) (p.Page[query.DewormResult], error)
	FindDewormsByCustomer(ctx context.Context, qry query.FindDewormsByCustomerQuery) (p.Page[query.DewormResult], error)
	FindDewormsByEmployee(ctx context.Context, qry query.FindDewormsByEmployeeQuery) (p.Page[query.DewormResult], error)
	FindDewormsByDateRange(ctx context.Context, qry query.FindDewormsByDateRangeQuery) (p.Page[query.DewormResult], error)
}

type dewormingFacadeService struct {
	queryHandler   query.DewormQueryHandler
	commandHandler command.DewormCommandHandler
}

func NewDewormingFacadeService(
	queryHandler query.DewormQueryHandler,
	commandHandler command.DewormCommandHandler,
) DewormingFacadeService {
	return &dewormingFacadeService{
		queryHandler:   queryHandler,
		commandHandler: commandHandler,
	}
}

// Command

func (d *dewormingFacadeService) RegisterDeworm(ctx context.Context, cmd command.DewormCreateCommand) cqrs.CommandResult {
	return d.commandHandler.HandleCreate(ctx, cmd)
}

func (d *dewormingFacadeService) UpdateDeworm(ctx context.Context, cmd command.DewormUpdateCommand) cqrs.CommandResult {
	return d.commandHandler.HandleUpdate(ctx, cmd)
}

func (d *dewormingFacadeService) DeleteDeworm(ctx context.Context, cmd command.DewormDeleteCommand) cqrs.CommandResult {
	return d.commandHandler.HandleDelete(ctx, cmd)
}

// Query

func (d *dewormingFacadeService) FindDewormByID(ctx context.Context, qry query.FindDewormByIDQuery) (query.DewormResult, error) {
	return d.queryHandler.HandleByIDQuery(ctx, qry)
}

func (d *dewormingFacadeService) FindDewormsByDateRange(ctx context.Context, qry query.FindDewormsByDateRangeQuery) (p.Page[query.DewormResult], error) {
	return d.queryHandler.HandleByDateRangeQuery(ctx, qry)
}

func (d *dewormingFacadeService) FindDewormsByEmployee(ctx context.Context, qry query.FindDewormsByEmployeeQuery) (p.Page[query.DewormResult], error) {
	return d.queryHandler.HandleByEmployeeQuery(ctx, qry)
}

func (d *dewormingFacadeService) FindDewormsByPet(ctx context.Context, qry query.FindDewormsByPetQuery) (p.Page[query.DewormResult], error) {
	return d.queryHandler.HandleByPetQuery(ctx, qry)
}

func (d *dewormingFacadeService) FindDewormsByCustomer(ctx context.Context, qry query.FindDewormsByCustomerQuery) (p.Page[query.DewormResult], error) {
	return d.queryHandler.HandleByCustomerQuery(ctx, qry)
}
