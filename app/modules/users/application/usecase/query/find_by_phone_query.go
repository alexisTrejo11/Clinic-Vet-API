package query

import (
	"context"
	"errors"

	"clinic-vet-api/app/core/domain/valueobject"
	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/shared/cqrs"
)

type FindUserByPhoneQuery struct {
	phone          valueobject.PhoneNumber
	ctx            context.Context
	includeProfile bool
}

func NewFindUserByPhoneQuery(ctx context.Context, phone string, includeProfile bool) (*FindUserByPhoneQuery, error) {
	phoneVO, err := valueobject.NewPhoneNumber(phone)
	if err != nil {
		return nil, err
	}

	return &FindUserByPhoneQuery{
		phone:          phoneVO,
		ctx:            ctx,
		includeProfile: includeProfile,
	}, nil
}

type FindUserByPhoneHandler struct {
	userRepository repository.UserRepository
}

func NewFindUserByPhoneHandler(userRepository repository.UserRepository) cqrs.QueryHandler[UserResult] {
	return &FindUserByPhoneHandler{
		userRepository: userRepository,
	}
}

func (h *FindUserByPhoneHandler) Handle(q cqrs.Query) (UserResult, error) {
	query, valid := q.(FindUserByPhoneQuery)
	if !valid {
		return UserResult{}, errors.New("invalid query type")
	}

	user, err := h.userRepository.FindByPhone(query.ctx, query.phone.Value())
	if err != nil {
		return UserResult{}, err
	}

	return userToResult(user), nil
}
