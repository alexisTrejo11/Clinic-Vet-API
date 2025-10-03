package query

import (
	"context"

	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
)

type FindUserByIDQuery struct {
	id valueobject.UserID
}

func NewFindUserByIDQuery(id uint) FindUserByIDQuery {
	return FindUserByIDQuery{id: valueobject.NewUserID(id)}
}

type FindUserByIDHandler struct {
	userRepository repository.UserRepository
}

func NewFindUserByIDHandler(userRepository repository.UserRepository) FindUserByIDHandler {
	return FindUserByIDHandler{
		userRepository: userRepository,
	}
}

func (h FindUserByIDHandler) Handle(ctx context.Context, query FindUserByIDQuery) (UserResult, error) {
	user, err := h.userRepository.FindByID(ctx, query.id)
	if err != nil {
		return UserResult{}, err
	}

	return userToResult(user), nil
}
