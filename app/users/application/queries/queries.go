package userQueries

type UserByIdQuery struct {
	ID string `json:"id"`
}

type UserByEmailQuery struct {
	Email string `json:"email"`
}

type UserSearchQuery struct {
	EmailLike     string `json:"email_like"`
	FirstNameLike string `json:"first_name_like"`
	LastNameLike  string `json:"last_name_like"`
	Limit         int    `json:"limit"`
	Offset        int    `json:"offset"`
}

type UserListByRoleQuery struct {
	Role   string `json:"role"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}
