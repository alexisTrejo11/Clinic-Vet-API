package bus

import (
	"errors"
	"reflect"

	"clinic-vet-api/app/core/repository"
	"clinic-vet-api/app/modules/users/application/usecase/query"
	"clinic-vet-api/app/shared/cqrs"
)

type UserQueryBus struct {
	handlers map[reflect.Type]any
}

func NewUserQueryBus(userRepo repository.UserRepository) *UserQueryBus {
	bus := &UserQueryBus{
		handlers: make(map[reflect.Type]any),
	}

	bus.RegisterQueries(userRepo)
	return bus
}

func (b *UserQueryBus) RegisterQueries(userRepo repository.UserRepository) error {
	b.Register(reflect.TypeOf(query.FindUserByEmailQuery{}), query.NewFindUserByEmailHandler(userRepo))
	b.Register(reflect.TypeOf(query.FindUserByIDQuery{}), query.NewFindUserByIDHandler(userRepo))
	b.Register(reflect.TypeOf(query.FindUserByPhoneQuery{}), query.NewFindUserByPhoneHandler(userRepo))
	b.Register(reflect.TypeOf(query.FindUsersByRoleQuery{}), query.NewFindUsersByRoleHandler(userRepo))
	b.Register(reflect.TypeOf(query.UserFindBySpecificationQuery{}), query.NewFindBySpecificationUsersHandler(userRepo))
	return nil
}

func (b *UserQueryBus) Register(queryType reflect.Type, handler any) error {
	if handler == nil {
		return errors.New("handler cannot be nil")
	}

	b.handlers[queryType] = handler
	return nil
}

func (b *UserQueryBus) Execute(q cqrs.Query) (any, error) {
	queryType := reflect.TypeOf(q)
	handler, ok := b.handlers[queryType]
	if !ok {
		return nil, errors.New("no handler registered for this query")
	}

	switch qry := q.(type) {
	case query.FindUserByEmailQuery:
		h := handler.(query.FindUserByEmailHandler)
		return h.Handle(qry)
	case query.FindUserByPhoneQuery:
		h := handler.(query.FindUserByPhoneHandler)
		return h.Handle(qry)
	case query.FindUserByIDQuery:
		h := handler.(query.FindUserByIDHandler)
		return h.Handle(qry)
	case query.UserFindBySpecificationQuery:
		h := handler.(query.FindBySpecificationUsersHandler)
		return h.Handle(qry)
	case query.FindUsersByRoleQuery:
		h := handler.(query.FindUsersByRoleHandler)
		return h.Handle(qry)
	default:
		return nil, errors.New("unhandled query type")
	}
}
