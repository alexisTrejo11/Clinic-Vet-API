package userQueries

import (
	"context"

	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type GetUserByIdQuery struct {
	Id             int             `json:"id"`
	IncludeProfile bool            `json:"include_profile"`
	Ctx            context.Context `json:"-"`
}

type GetUserByIdHandler interface {
	Handle(query GetUserByIdQuery) (*UserResponse, error)
}

type GetUserByIdHandlerImpl struct {
	userRepository userRepository.UserRepository
}

func NewGetUserByIdHandler(userRepository userRepository.UserRepository) GetUserByIdHandler {
	return &GetUserByIdHandlerImpl{
		userRepository: userRepository,
	}
}

func (h *GetUserByIdHandlerImpl) Handle(query GetUserByIdQuery) (*UserResponse, error) {
	user, err := h.userRepository.GetById(query.Ctx, query.Id)
	if err != nil {
		return nil, err
	}

	return toResponse(*user), nil
}
