package bus

import (
	"clinic-vet-api/app/modules/account/user/application/query"
	"clinic-vet-api/app/modules/core/repository"
	"clinic-vet-api/app/shared/page"
	"context"
)

type UserQueryBus interface {
	FindUserByID(ctx context.Context, q query.FindUserByIDQuery) (query.UserResult, error)
	FindUserByEmail(ctx context.Context, q query.FindUserByEmailQuery) (query.UserResult, error)
	FindUserByPhone(ctx context.Context, q query.FindUserByPhoneQuery) (query.UserResult, error)
	FindUsersByRole(ctx context.Context, q query.FindUsersByRoleQuery) (page.Page[query.UserResult], error)
	FindUsersBySpecification(ctx context.Context, q query.FindUserBySpecificationQuery) (page.Page[query.UserResult], error)
}

type userQueryBus struct {
	findByIDHandler            query.FindUserByIDHandler
	findByEmailHandler         query.FindUserByEmailHandler
	findByPhoneHandler         query.FindUserByPhoneHandler
	findByRoleHandler          query.FindUsersByRoleHandler
	findBySpecificationHandler query.FindUserBySpecificationHandler
}

func NewUserQueryBus(repository repository.UserRepository) UserQueryBus {
	findByIDHandler := query.NewFindUserByIDHandler(repository)
	findByEmailHandler := query.NewFindUserByEmailHandler(repository)
	findByPhoneHandler := query.NewFindUserByPhoneHandler(repository)
	findByRoleHandler := query.NewFindUsersByRoleHandler(repository)
	findBySpecHandler := query.NewFindUserBySpecificationHandler(repository)

	return &userQueryBus{
		findByIDHandler:            *findByIDHandler,
		findByEmailHandler:         *findByEmailHandler,
		findByPhoneHandler:         *findByPhoneHandler,
		findByRoleHandler:          *findByRoleHandler,
		findBySpecificationHandler: *findBySpecHandler,
	}
}

func (b *userQueryBus) FindUserByID(ctx context.Context, q query.FindUserByIDQuery) (query.UserResult, error) {
	return b.findByIDHandler.Handle(ctx, q)
}

func (b *userQueryBus) FindUserByEmail(ctx context.Context, q query.FindUserByEmailQuery) (query.UserResult, error) {
	return b.findByEmailHandler.Handle(ctx, q)
}

func (b *userQueryBus) FindUserByPhone(ctx context.Context, q query.FindUserByPhoneQuery) (query.UserResult, error) {
	return b.findByPhoneHandler.Handle(ctx, q)
}

func (b *userQueryBus) FindUsersByRole(ctx context.Context, q query.FindUsersByRoleQuery) (page.Page[query.UserResult], error) {
	return b.findByRoleHandler.Handle(ctx, q)
}

func (b *userQueryBus) FindUsersBySpecification(ctx context.Context, q query.FindUserBySpecificationQuery) (page.Page[query.UserResult], error) {
	return b.findBySpecificationHandler.Handle(ctx, q)
}
