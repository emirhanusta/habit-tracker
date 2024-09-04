package habit

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
	habitRepository repository.IHabitRepository
}

func NewCommandHandler(habitRepository repository.IHabitRepository) ICommandHandler {
	return &commandHandler{
		habitRepository: habitRepository,
	}
}

func (c *commandHandler) Save(ctx context.Context, command CreateCommand) error {
	if err := c.habitRepository.Save(ctx, c.buildEntityCreate(command)); err != nil {
		return err
	}
	fmt.Printf("commandHandler.Save Started with name: %s\n", command.Name)
	return nil
}

func (c *commandHandler) Update(ctx context.Context, command UpdateCommand) error {
	if err := c.habitRepository.Update(ctx, c.buildEntityUpdate(command)); err != nil {
		return err
	}
	fmt.Printf("commandHandler.Update Started with name: %s\n", command.Name)
	return nil
}

func (c *commandHandler) Delete(ctx context.Context, id string) error {
	if err := c.habitRepository.Delete(ctx, id); err != nil {
		return err
	}
	fmt.Printf("commandHandler.Delete Started with id: %s\n", id)
	return nil
}

func (c *commandHandler) buildEntityCreate(command CreateCommand) *domain.Habit {
	return &domain.Habit{
		Name:        command.Name,
		Description: command.Description,
		UserId:      command.UserId,
	}
}

func (c *commandHandler) buildEntityUpdate(command UpdateCommand) *domain.Habit {
	return &domain.Habit{
		Id:          command.Id,
		Name:        command.Name,
		Description: command.Description,
		UserId:      command.UserId,
	}
}
