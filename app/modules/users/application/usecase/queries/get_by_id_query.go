package userDomainQueries

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
)

type GetUserByIDQuery struct {
	ID             valueobject.UserID `json:"id"`
	IncludeProfile bool               `json:"include_profile"`
	Ctx            context.Context    `json:"-"`
}

type GetUserByIDHandler interface {
	Handle(query GetUserByIDQuery) (UserResponse, error)
}

type GetUserByIDHandlerImpl struct {
	userRepository repository.UserRepository
}

func NewGetUserByIDHandler(userRepository repository.UserRepository) GetUserByIDHandler {
	return &GetUserByIDHandlerImpl{
		userRepository: userRepository,
	}
}

func (h *GetUserByIDHandlerImpl) Handle(query GetUserByIDQuery) (UserResponse, error) {
	user, err := h.userRepository.GetByID(query.Ctx, query.ID)
	if err != nil {
		return UserResponse{}, err
	}

	return toResponse(user), nil
}
