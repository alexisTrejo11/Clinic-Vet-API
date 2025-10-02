package bus

type UserBus struct {
	UserQueryBus
	UserCommandBus
}

func NewUserBus(queryBus UserQueryBus, commandBus UserCommandBus) *UserBus {
	return &UserBus{
		UserQueryBus:   queryBus,
		UserCommandBus: commandBus,
	}
}
