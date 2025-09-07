package userDomainQueries

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
)

type GetByPhoneQuery struct {
	Phone          valueobject.PhoneNumber `json:"phone"`
	Ctx            context.Context         `json:"-"`
	IncludeProfile bool                    `json:"include_profile"`
}

type GetByPhoneHandler interface {
	Handle(ctx context.Context, query GetByPhoneQuery) (UserResponse, error)
}

type getByPhoneHandler struct {
	userRepository repository.UserRepository
}

func NewGetByPhoneHandler(userRepository repository.UserRepository) GetByPhoneHandler {
	return &getByPhoneHandler{
		userRepository: userRepository,
	}
}

func (h *getByPhoneHandler) Handle(ctx context.Context, query GetByPhoneQuery) (UserResponse, error) {
	user, err := h.userRepository.GetByPhone(ctx, query.Phone.Value())
	if err != nil {
		return UserResponse{}, err
	}

	return toResponse(user), nil
}
