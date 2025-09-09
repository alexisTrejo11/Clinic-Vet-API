package query

import (
	"context"
	"errors"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/core/domain/valueobject"
	repository "github.com/alexisTrejo11/Clinic-Vet-API/app/core/repositories"
	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/cqrs"
)

type GetUserByIDQuery struct {
	id             valueobject.UserID
	includeProfile bool
	ctx            context.Context
}

func NewGetUserByIDQuery(ctx context.Context, id int, includeProfile bool) (GetUserByIDQuery, error) {
	userID, err := valueobject.NewUserID(id)
	if err != nil {
		return GetUserByIDQuery{}, err
	}

	return GetUserByIDQuery{
		id:             userID,
		includeProfile: includeProfile,
		ctx:            ctx,
	}, nil
}

type GetUserByIDHandler struct {
	userRepository repository.UserRepository
}

func NewGetUserByIDHandler(userRepository repository.UserRepository) cqrs.QueryHandler[UserResponse] {
	return &GetUserByIDHandler{
		userRepository: userRepository,
	}
}

func (h *GetUserByIDHandler) Handle(q cqrs.Query) (UserResponse, error) {
	query, valid := q.(GetUserByIDQuery)
	if !valid {
		return UserResponse{}, errors.New("invalid query type")
	}

	user, err := h.userRepository.GetByID(query.ctx, query.id)
	if err != nil {
		return UserResponse{}, err
	}

	return toResponse(user), nil
}
