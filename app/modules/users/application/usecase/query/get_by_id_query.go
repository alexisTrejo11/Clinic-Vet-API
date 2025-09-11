package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/repository"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type GetUserByIDQuery struct {
	id             valueobject.UserID
	includeProfile bool
	ctx            context.Context
}

func NewGetUserByIDQuery(ctx context.Context, id uint, includeProfile bool) *GetUserByIDQuery {
	userID := valueobject.NewUserID(id)
	return &GetUserByIDQuery{
		id:             userID,
		includeProfile: includeProfile,
		ctx:            ctx,
	}
}

type GetUserByIDHandler struct {
	userRepository repository.UserRepository
}

func NewGetUserByIDHandler(userRepository repository.UserRepository) cqrs.QueryHandler[UserResult] {
	return &GetUserByIDHandler{
		userRepository: userRepository,
	}
}

func (h *GetUserByIDHandler) Handle(q cqrs.Query) (UserResult, error) {
	query, valid := q.(GetUserByIDQuery)
	if !valid {
		return UserResult{}, errors.New("invalid query type")
	}

	user, err := h.userRepository.FindByID(query.ctx, query.id)
	if err != nil {
		return UserResult{}, err
	}

	return userToResult(user), nil
}
