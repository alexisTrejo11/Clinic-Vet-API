package userDomainQueries

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/entity/user"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
)

type GetByEmailQuery struct {
	Email          valueobject.Email `json:"email"`
	IncludeProfile bool              `json:"include_profile"`
	Ctx            context.Context   `json:"-"`
}

type GetByEmailHandler interface {
	Handle(ctx context.Context, query GetByEmailQuery) (user.User, error)
}

type getByEmailHandler struct {
	userRepository repository.UserRepository
}

func NewGetByEmailHandler(userRepository repository.UserRepository) GetByEmailHandler {
	return &getByEmailHandler{
		userRepository: userRepository,
	}
}

func (h *getByEmailHandler) Handle(ctx context.Context, query GetByEmailQuery) (user.User, error) {
	userEntity, err := h.userRepository.GetByEmail(ctx, query.Email.Value())
	if err != nil {
		return user.User{}, err
	}

	return userEntity, nil
}
