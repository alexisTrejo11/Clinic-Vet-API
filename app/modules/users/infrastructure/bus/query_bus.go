package bus

import (
	"errors"
	"reflect"

	repository "clinic-vet-api/app/core/repositories"
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
	b.Register(reflect.TypeOf(query.GetByEmailQuery{}), query.NewGetByEmailHandler(userRepo))
	b.Register(reflect.TypeOf(query.GetUserByIDQuery{}), query.NewGetUserByIDHandler(userRepo))
	b.Register(reflect.TypeOf(query.GetUserByPhoneQuery{}), query.NewGetUserByPhoneHandler(userRepo))
	b.Register(reflect.TypeOf(query.ListUsersByRoleQuery{}), query.NewListUsersByRoleHandler(userRepo))
	b.Register(reflect.TypeOf(query.UserSearchQuery{}), query.NewSearchUsersHandler(userRepo))
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
	case query.GetByEmailQuery:
		h := handler.(query.GetByEmailHandler)
		return h.Handle(qry)
	case query.GetUserByPhoneQuery:
		h := handler.(query.GetUserByPhoneHandler)
		return h.Handle(qry)
	case query.GetUserByIDQuery:
		h := handler.(query.GetUserByIDHandler)
		return h.Handle(qry)
	case query.UserSearchQuery:
		h := handler.(query.SearchUsersHandler)
		return h.Handle(qry)
	case query.ListUsersByRoleQuery:
		h := handler.(query.ListUsersByRoleHandler)
		return h.Handle(qry)
	default:
		return nil, errors.New("unhandled query type")
	}
}
