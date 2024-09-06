package reminder

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
	reminderRepository repository.IReminderRepository
}

func NewCommandHandler(reminderRepository repository.IReminderRepository) ICommandHandler {
	return &commandHandler{
		reminderRepository: reminderRepository,
	}
}

func (c *commandHandler) Save(ctx context.Context, command CreateCommand) error {
	if err := c.reminderRepository.Save(ctx, c.buildEntityCreate(command)); err != nil {
		return err
	}
	fmt.Printf("ReminderCommandHandler.Save Started with message: %s\n", command.Message)
	return nil
}

func (c *commandHandler) Update(ctx context.Context, command UpdateCommand) error {
	if err := c.reminderRepository.Update(ctx, c.buildEntityUpdate(command)); err != nil {
		return err
	}
	fmt.Printf("ReminderCommandHandler.Update Started with message: %s\n", command.Message)
	return nil
}

func (c *commandHandler) Delete(ctx context.Context, id string) error {
	if err := c.reminderRepository.Delete(ctx, id); err != nil {
		return err
	}
	fmt.Printf("ReminderCommandHandler.Delete Started with id: %s\n", id)
	return nil
}

func (c *commandHandler) buildEntityCreate(command CreateCommand) *domain.Reminder {
	return &domain.Reminder{
		Message:  command.Message,
		HabitId:  command.HabitId,
		RemindAt: command.RemindAt,
	}
}

func (c *commandHandler) buildEntityUpdate(command UpdateCommand) *domain.Reminder {
	return &domain.Reminder{
		Id:       command.Id,
		Message:  command.Message,
		HabitId:  command.HabitId,
		RemindAt: command.RemindAt,
	}
}
