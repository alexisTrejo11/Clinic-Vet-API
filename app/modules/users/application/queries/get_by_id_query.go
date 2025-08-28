package userDomainQueries

import (
	"context"

	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type GetUserByIdQuery struct {
	Id             int             `json:"id"`
	IncludeProfile bool            `json:"include_profile"`
	Ctx            context.Context `json:"-"`
}

type GetUserByIdHandler interface {
	Handle(query GetUserByIdQuery) (UserResponse, error)
}

type GetUserByIdHandlerImpl struct {
	userRepository userDomain.UserRepository
}

func NewGetUserByIdHandler(userRepository userDomain.UserRepository) GetUserByIdHandler {
	return &GetUserByIdHandlerImpl{
		userRepository: userRepository,
	}
}

func (h *GetUserByIdHandlerImpl) Handle(query GetUserByIdQuery) (UserResponse, error) {
	user, err := h.userRepository.GetByID(query.Ctx, query.Id)
	if err != nil {
		return UserResponse{}, err
	}

	return toResponse(user), nil
}
