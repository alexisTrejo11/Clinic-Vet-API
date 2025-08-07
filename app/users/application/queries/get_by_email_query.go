package userQueries

import (
	"context"

	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type GetByEmailQuery struct {
	Email          user.Email      `json:"email"`
	IncludeProfile bool            `json:"include_profile"`
	Ctx            context.Context `json:"-"`
}

type GetByEmailHandler interface {
	Handle(ctx context.Context, query GetByEmailQuery) (*user.User, error)
}

type getByEmailHandler struct {
	userRepository userRepo.UserRepository
}

func NewGetByEmailHandler(userRepository userRepo.UserRepository) GetByEmailHandler {
	return &getByEmailHandler{
		userRepository: userRepository,
	}
}

func (h *getByEmailHandler) Handle(ctx context.Context, query GetByEmailQuery) (*user.User, error) {
	user, err := h.userRepository.GetByEmail(ctx, query.Email.Value())
	if err != nil {
		return nil, err
	}

	return user, nil
}
