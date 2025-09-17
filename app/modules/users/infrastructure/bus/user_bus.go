package bus

type UserBus struct {
	QueryBus   UserQueryBus
	CommandBus UserCommandBus
}

func NewUserBus(queryBus UserQueryBus, commandBus UserCommandBus) *UserBus {

	return &UserBus{
		QueryBus:   queryBus,
		CommandBus: commandBus,
	}
}
