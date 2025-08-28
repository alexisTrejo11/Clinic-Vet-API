package userDomainQueries

import (
	"context"

	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
)

type GetByEmailQuery struct {
	Email          userDomain.Email `json:"email"`
	IncludeProfile bool             `json:"include_profile"`
	Ctx            context.Context  `json:"-"`
}

type GetByEmailHandler interface {
	Handle(ctx context.Context, query GetByEmailQuery) (userDomain.User, error)
}

type getByEmailHandler struct {
	userRepository userDomain.UserRepository
}

func NewGetByEmailHandler(userRepository userDomain.UserRepository) GetByEmailHandler {
	return &getByEmailHandler{
		userRepository: userRepository,
	}
}

func (h *getByEmailHandler) Handle(ctx context.Context, query GetByEmailQuery) (userDomain.User, error) {
	user, err := h.userRepository.GetByEmail(ctx, query.Email.Value())
	if err != nil {
		return userDomain.User{}, err
	}

	return user, nil
}
