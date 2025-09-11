package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type GetUserByEmailQuery struct {
	email          valueobject.Email
	includeProfile bool
	ctx            context.Context
}

func NewGetUserByEmailQuery(ctx context.Context, email string, includeProfile bool) (*GetUserByEmailQuery, error) {
	emailVO, err := valueobject.NewEmail(email)
	if err != nil {
		return nil, err
	}

	return &GetUserByEmailQuery{
		email:          emailVO,
		includeProfile: includeProfile,
		ctx:            ctx,
	}, nil
}

type GetUserByEmailHandler struct {
	userRepository repository.UserRepository
}

func NewGetUserByEmailHandler(userRepository repository.UserRepository) cqrs.QueryHandler[UserResult] {
	return &GetUserByEmailHandler{
		userRepository: userRepository,
	}
}

func (h *GetUserByEmailHandler) Handle(q cqrs.Query) (UserResult, error) {
	query, ok := q.(GetUserByEmailQuery)
	if !ok {
		return UserResult{}, errors.New("invalid query type")
	}

	user, err := h.userRepository.FindByEmail(query.ctx, query.email.Value())
	if err != nil {
		return UserResult{}, err
	}

	return userToResult(user), nil
}
