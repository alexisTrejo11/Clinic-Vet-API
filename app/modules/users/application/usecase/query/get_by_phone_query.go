package query

import (
	"context"
	"errors"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
)

type GetUserByPhoneQuery struct {
	phone          valueobject.PhoneNumber
	ctx            context.Context
	includeProfile bool
}

func NewGetUserByPhoneQuery(ctx context.Context, phone string, includeProfile bool) (*GetUserByPhoneQuery, error) {
	phoneVO, err := valueobject.NewPhoneNumber(phone)
	if err != nil {
		return nil, err
	}

	return &GetUserByPhoneQuery{
		phone:          phoneVO,
		ctx:            ctx,
		includeProfile: includeProfile,
	}, nil
}

type GetUserByPhoneHandler struct {
	userRepository repository.UserRepository
}

func NewGetUserByPhoneHandler(userRepository repository.UserRepository) cqrs.QueryHandler[UserResult] {
	return &GetUserByPhoneHandler{
		userRepository: userRepository,
	}
}

func (h *GetUserByPhoneHandler) Handle(q cqrs.Query) (UserResult, error) {
	query, valid := q.(GetUserByPhoneQuery)
	if !valid {
		return UserResult{}, errors.New("invalid query type")
	}

	user, err := h.userRepository.FindByPhone(query.ctx, query.phone.Value())
	if err != nil {
		return UserResult{}, err
	}

	return userToResult(user), nil
}
