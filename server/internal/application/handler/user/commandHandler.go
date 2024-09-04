package user

import (
	"context"
	"fmt"
	"habit-tracker/internal/application/repository"
	"habit-tracker/internal/domain"
)

type ICommandHandler interface {
	Save(ctx context.Context, command CreateCommand) error
	Update(ctx context.Context, command UpdateCommand) error
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

func (c *commandHandler) Save(ctx context.Context, command CreateCommand) error {
	if err := c.userRepository.Save(ctx, c.buildEntityCreate(command)); err != nil {
		return err
	}
	fmt.Printf("commandHandler.Save Started with username: %s\n", command.Username)
	return nil
}

func (c *commandHandler) Update(ctx context.Context, command UpdateCommand) error {
	if err := c.userRepository.Update(ctx, c.buildEntityUpdate(command)); err != nil {
		return err
	}
	fmt.Printf("commandHandler.Update Started with username: %s\n", command.Username)
	return nil
}

func (c *commandHandler) Delete(ctx context.Context, id string) error {
	if err := c.userRepository.Delete(ctx, id); err != nil {
		return err
	}
	fmt.Printf("commandHandler.Delete Started with id: %s\n", id)
	return nil
}

func (c *commandHandler) buildEntityCreate(command CreateCommand) *domain.User {
	return &domain.User{
		Username: command.Username,
		Email:    command.Email,
		Password: command.Password,
	}
}

func (c *commandHandler) buildEntityUpdate(command UpdateCommand) *domain.User {
	return &domain.User{
		Id:       command.Id,
		Username: command.Username,
		Email:    command.Email,
		Password: command.Password,
	}
}
