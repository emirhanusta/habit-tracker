package query

import (
	"context"
	"fmt"
	"habit-tracker/internal/application/repository"
	"habit-tracker/internal/domain"
)

type IUserQueryService interface {
	GetAll(context context.Context) ([]domain.User, error)
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

func (u *userQueryService) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := u.userRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Printf("userQueryService.GetAll Started\n")

	return users, nil
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
	user, err := u.userRepository.GetByEmail(ctx, email)

	if err != nil {
		return nil, err
	}
	fmt.Printf("userQueryService.GetByEmail Started with email: %s\n", email)
	return user, nil
}
