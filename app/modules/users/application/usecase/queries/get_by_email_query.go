package userDomainQueries

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/entity/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
)

type GetByEmailQuery struct {
	Email          valueobject.Email `json:"email"`
	IncludeProfile bool              `json:"include_profile"`
	Ctx            context.Context   `json:"-"`
}

type GetByEmailHandler interface {
	Handle(ctx context.Context, query GetByEmailQuery) (entity.User, error)
}

type getByEmailHandler struct {
	userRepository repository.UserRepository
}

func NewGetByEmailHandler(userRepository repository.UserRepository) GetByEmailHandler {
	return &getByEmailHandler{
		userRepository: userRepository,
	}
}

func (h *getByEmailHandler) Handle(ctx context.Context, query GetByEmailQuery) (entity.User, error) {
	user, err := h.userRepository.GetByEmail(ctx, query.Email.Value())
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}
