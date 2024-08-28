package query

import (
	"context"
	"fmt"
	"habit-tracker/internal/application/repository"
	"habit-tracker/internal/domain"
)

type IUserQueryService interface {
	GetById(context context.Context, id string) (*domain.User, error)
	GetByEmail(context context.Context, email string) (*domain.User, error)
}

type userQueryService struct {
	userRepository repository.IUserRepository
}

func NewUserQueryService(userRepository repository.IUserRepository) IUserQueryService {
	return &userQueryService{
		userRepository: userRepository,
	}
}

func (u *userQueryService) GetById(ctx context.Context, id string) (*domain.User, error) {
	user, err := u.userRepository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	fmt.Printf("userQueryService.GetById Started with userId: %s\n", id)

	return user, nil
}

func (u *userQueryService) GetByEmail(ctx context.Context, email string) (*domain.User, error) {
	user, err := u.userRepository.GetUserByEmail(ctx, email)

	if err != nil {
		return nil, err
	}

	return user, nil
}
