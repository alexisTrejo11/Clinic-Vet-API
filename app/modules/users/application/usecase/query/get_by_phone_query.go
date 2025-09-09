package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type GetUserByPhoneQuery struct {
	Phone          valueobject.PhoneNumber `json:"phone"`
	Ctx            context.Context         `json:"-"`
	IncludeProfile bool                    `json:"include_profile"`
}

type GetUserByPhoneHandler struct {
	userRepository repository.UserRepository
}

func NewGetUserByPhoneHandler(userRepository repository.UserRepository) cqrs.QueryHandler[UserResponse] {
	return &GetUserByPhoneHandler{
		userRepository: userRepository,
	}
}

func (h *GetUserByPhoneHandler) Handle(q cqrs.Query) (UserResponse, error) {
	query, valid := q.(GetUserByPhoneQuery)
	if !valid {
		return UserResponse{}, errors.New("invalid query type")
	}

	user, err := h.userRepository.GetByPhone(query.Ctx, query.Phone.Value())
	if err != nil {
		return UserResponse{}, err
	}

	return toResponse(user), nil
}
