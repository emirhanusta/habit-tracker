package user

import (
	"context"
	"habit-tracker/internal/application/repository"
	"habit-tracker/internal/domain"
)

type ICommandHandler interface {
	Save(ctx context.Context, command Command) error
}

type commandHandler struct {
	userRepository repository.IUserRepository
}

func NewCommandHandler(userRepository repository.IUserRepository) ICommandHandler {
	return &commandHandler{
		userRepository: userRepository,
	}
}

func (c *commandHandler) Save(ctx context.Context, command Command) error {
	if err := c.userRepository.SaveUser(ctx, c.BuildEntity(command)); err != nil {
		return err
	}

	return nil
}

func (c *commandHandler) BuildEntity(command Command) *domain.User {
	return &domain.User{
		Username: command.Username,
		Email:    command.Email,
		Password: command.Password,
	}
}
