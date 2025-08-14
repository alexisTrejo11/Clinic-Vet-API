package userQueries

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/valueObjects"
	user "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepo "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type GetByPhoneQuery struct {
	Phone          valueObjects.PhoneNumber `json:"phone"`
	Ctx            context.Context          `json:"-"`
	IncludeProfile bool                     `json:"include_profile"`
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
	user, err := h.userRepository.GetByPhone(ctx, query.Phone.Value())
	if err != nil {
		return nil, err
	}

	return user, nil
}
