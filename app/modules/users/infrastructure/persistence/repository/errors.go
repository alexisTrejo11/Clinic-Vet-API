package repositoryimpl

const (
	TableUsers = "users"

	// Mensajes de error específicos
	ErrMsgGetUser             = "failed to get user"
	ErrMsgGetUserByEmail      = "failed to get user by email"
	ErrMsgGetUserByPhone      = "failed to get user by phone"
	ErrMsgListUsers           = "failed to list users"
	ErrMsgSearchUsers         = "failed to search users"
	ErrMsgCreateUser          = "failed to create user"
	ErrMsgUpdateUser          = "failed to update user"
	ErrMsgSoftDeleteUser      = "failed to soft delete user"
	ErrMsgHardDeleteUser      = "failed to hard delete user"
	ErrMsgCheckUserExists     = "failed to check if user exists"
	ErrMsgUpdateLastLogin     = "failed to update last login"
	ErrMsgConvertUserToDomain = "failed to convert user to domain entity"

	DriverSQL = "SQLC"
	OpSelect  = "SELECT"
	OpInsert  = "INSERT"
	OpUpdate  = "UPDATE"
	OpDelete  = "DELETE"
	OpCount   = "COUNT"
)
