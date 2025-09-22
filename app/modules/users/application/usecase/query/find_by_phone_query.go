package query

import (
	"context"

	"clinic-vet-api/app/modules/core/domain/valueobject"
	"clinic-vet-api/app/modules/core/repository"
)

type FindUserByPhoneQuery struct {
	phone          valueobject.PhoneNumber
	includeProfile bool
}

func NewFindUserByPhoneQuery(phone string, includeProfile bool) (*FindUserByPhoneQuery, error) {
	phoneVO, err := valueobject.NewPhoneNumber(phone)
	if err != nil {
		return nil, err
	}

	return &FindUserByPhoneQuery{
		phone:          phoneVO,
		includeProfile: includeProfile,
	}, nil
}

type FindUserByPhoneHandler struct {
	userRepository repository.UserRepository
}

func NewFindUserByPhoneHandler(userRepository repository.UserRepository) *FindUserByPhoneHandler {
	return &FindUserByPhoneHandler{
		userRepository: userRepository,
	}
}

func (h *FindUserByPhoneHandler) Handle(ctx context.Context, query FindUserByPhoneQuery) (UserResult, error) {
	user, err := h.userRepository.FindByPhone(ctx, query.phone.Value())
	if err != nil {
		return UserResult{}, err
	}

	return userToResult(user), nil
}
