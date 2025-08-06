package userUsecase

import (
	"context"

	"github.com/alexisTrejo11/Clinic-Vet-API/app/shared/page"
	queries "github.com/alexisTrejo11/Clinic-Vet-API/app/users/application/queries"
	userDomain "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain"
	userRepository "github.com/alexisTrejo11/Clinic-Vet-API/app/users/domain/repositories"
)

type UserQueryUsecases struct {
	userRepo userRepository.UserRepository
}

func NewUserQueryUsecases(userRepo userRepository.UserRepository) *UserQueryUsecases {
	return &UserQueryUsecases{
		userRepo: userRepo,
	}
}

func (u *UserQueryUsecases) GetUserByID(context context.Context, id int) (*queries.UserResponse, error) {
	user, err := u.userRepo.GetById(context, id)
	if err != nil {
		return nil, err
	}
	return toResponse(*user), nil
}

func (u *UserQueryUsecases) SearchUsers(context context.Context, query queries.UserSearchQuery) (page.Page[[]queries.UserResponse], error) {
	userPage, err := u.userRepo.Search(context, query.ToMap(), query.Pagination)
	if err != nil {
		return page.EmptyPage[[]queries.UserResponse](), err
	}
	return toResponsePage(userPage), nil
}

func (u *UserQueryUsecases) GetUserByEmail(context context.Context, email string) (queries.UserResponse, error) {
	user, err := u.userRepo.GetByEmail(context, email)
	if err != nil {
		return queries.UserResponse{}, err
	}
	return *toResponse(*user), nil
}

func (u *UserQueryUsecases) GetUserByPhone(context context.Context, phone string) (queries.UserResponse, error) {
	user, err := u.userRepo.GetByPhone(context, phone)
	if err != nil {
		return queries.UserResponse{}, err
	}
	return *toResponse(*user), nil
}

func toResponse(user userDomain.User) *queries.UserResponse {
	return &queries.UserResponse{
		Id:          user.Id().String(),
		Email:       user.Email().String(),
		PhoneNumber: user.PhoneNumber().String(),
		Role:        user.Role().String(),
		Status:      string(user.Status()),
		JoinedAt:    user.JoinedAt().Format("2006-01-02 15:04:05"),
		LastLoginAt: user.LastLoginAt().Format("2006-01-02 15:04:05"),
	}
}

func toResponsePage(userPage page.Page[[]userDomain.User]) page.Page[[]queries.UserResponse] {
	var userResponses []queries.UserResponse

	users := userPage.Data
	for _, user := range users {
		userResponses = append(userResponses, *toResponse(user))
	}

	return *page.NewPage(userResponses, userPage.Metadata)
}
