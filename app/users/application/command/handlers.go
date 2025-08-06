package userCommand

import userUsecase "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application/usecases"

type CreateUserCommandHandler interface {
	Handle(command CreateUserCommand) error
}

type CreateUserCommandHandlerImpl struct {
	useCase userUsecase.CreateUserUseCase
}

func NewCreateUserCommandHandler(useCase userUsecase.CreateUserUseCase) *CreateUserCommandHandlerImpl {
	return &CreateUserCommandHandlerImpl{
		useCase: useCase,
	}
}

func (h *CreateUserCommandHandlerImpl) Handle(command CreateUserCommand) error {
	// Here you would typically map the command to a domain model and call the use case
	return h.useCase.Execute()
}
