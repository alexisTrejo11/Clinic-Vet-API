package userQueries

import (
	"context"

	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type GetByPhoneQuery struct {
	Email          user.Email      `json:"email"`
	Ctx            context.Context `json:"-"`
	IncludeProfile bool            `json:"include_profile"`
}

type GetByPhoneHandler interface {
	Handle(ctx context.Context, query GetByPhoneQuery) (*user.User, error)
}

type getByPhoneHandler struct {
	userRepository userRepo.UserRepository
}

func NewGetByPhoneHandler(userRepository userRepo.UserRepository) GetByPhoneHandler {
	return &getByPhoneHandler{
		userRepository: userRepository,
	}
}

func (h *getByPhoneHandler) Handle(ctx context.Context, query GetByPhoneQuery) (*user.User, error) {
	user, err := h.userRepository.GetByPhone(ctx, query.Email.Value())
	if err != nil {
		return nil, err
	}

	return user, nil
}
