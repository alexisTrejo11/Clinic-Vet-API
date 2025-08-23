package userDomainQueries

import (
	"context"

	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type GetByPhoneQuery struct {
	Phone          userDomain.PhoneNumber `json:"phone"`
	Ctx            context.Context        `json:"-"`
	IncludeProfile bool                   `json:"include_profile"`
}

type GetByPhoneHandler interface {
	Handle(ctx context.Context, query GetByPhoneQuery) (*userDomain.User, error)
}

type getByPhoneHandler struct {
	userRepository userDomain.UserRepository
}

func NewGetByPhoneHandler(userRepository userDomain.UserRepository) GetByPhoneHandler {
	return &getByPhoneHandler{
		userRepository: userRepository,
	}
}

func (h *getByPhoneHandler) Handle(ctx context.Context, query GetByPhoneQuery) (*userDomain.User, error) {
	user, err := h.userRepository.GetByPhone(ctx, query.Phone.Value())
	if err != nil {
		return nil, err
	}

	return user, nil
}
