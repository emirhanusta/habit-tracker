package user

import (
	"context"
	"fmt"
	"habit-tracker/internal/application/repository"
	"habit-tracker/internal/domain"
)

type ICommandHandler interface {
	Save(ctx context.Context, command Command) error
	Update(ctx context.Context, command Command) error
	Delete(ctx context.Context, id string) error
}

type commandHandler struct {
	userRepository repository.IUserRepository
}

func NewCommandHandler(userRepository repository.IUserRepository) ICommandHandler {
	return &commandHandler{
		userRepository: userRepository,
	}
}

func (c *commandHandler) Update(ctx context.Context, command Command) error {
	if err := c.userRepository.UpdateUser(ctx, c.BuildEntity(command)); err != nil {
		return err
	}
	fmt.Printf("commandHandler.Update Started with username: %s\n", command.Username)
	return nil
}

func (c *commandHandler) Save(ctx context.Context, command Command) error {
	if err := c.userRepository.SaveUser(ctx, c.BuildEntity(command)); err != nil {
		return err
	}
	fmt.Printf("commandHandler.Save Started with username: %s\n", command.Username)
	return nil
}

func (c *commandHandler) Delete(ctx context.Context, id string) error {
	if err := c.userRepository.DeleteUser(ctx, id); err != nil {
		return err
	}
	fmt.Printf("commandHandler.Delete Started with id: %s\n", id)
	return nil
}

func (c *commandHandler) BuildEntity(command Command) *domain.User {
	return &domain.User{
		Id:       command.Id,
		Username: command.Username,
		Email:    command.Email,
		Password: command.Password,
	}
}
