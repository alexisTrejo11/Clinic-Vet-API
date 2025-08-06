package userUsecase

type UserUseCases interface {
	CreateUserUseCase()
	GetUserByIDUseCase()
	SearchUsersUseCase()
	GetUserByEmailUseCase()
	GetUserByPhoneUseCase()
	DeleteUserUseCase()
}

type userUseCasesImpl struct {
	createUserUseCase *CreateUserUseCase

	deleteUserUseCase *DeleteUserUseCase
}

func (u *userUseCasesImpl) CreateUserUseCase() *CreateUserUseCase {
	return u.createUserUseCase
}
