package query

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type GetByEmailQuery struct {
	Email          valueobject.Email `json:"email"`
	IncludeProfile bool              `json:"include_profile"`
	Ctx            context.Context   `json:"-"`
}

type GetByEmailHandler struct {
	userRepository repository.UserRepository
}

func NewGetByEmailHandler(userRepository repository.UserRepository) cqrs.QueryHandler[UserResponse] {
	return &GetByEmailHandler{
		userRepository: userRepository,
	}
}

func (h *GetByEmailHandler) Handle(q cqrs.Query) (UserResponse, error) {
	query := q.(GetByEmailQuery)
	user, err := h.userRepository.GetByEmail(query.Ctx, query.Email.Value())
	if err != nil {
		return UserResponse{}, err
	}

	return toResponse(user), nil
}
