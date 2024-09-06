package query

import (
	"context"
	"habit-tracker/internal/application/repository"
	"habit-tracker/internal/domain"
)

type IReminderQueryService interface {
	GetAllByHabitId(ctx context.Context, habitId string) ([]domain.Reminder, error)
	GetById(ctx context.Context, id string) (*domain.Reminder, error)
}

type reminderQueryService struct {
	reminderRepository repository.IReminderRepository
}

func NewReminderQueryService(reminderRepository repository.IReminderRepository) IReminderQueryService {
	return &reminderQueryService{
		reminderRepository: reminderRepository,
	}
}

func (r *reminderQueryService) GetAllByHabitId(ctx context.Context, habitId string) ([]domain.Reminder, error) {
	reminders, err := r.reminderRepository.GetAllByHabitId(ctx, habitId)

	if err != nil {
		return nil, err
	}

	return reminders, nil
}

func (r *reminderQueryService) GetById(ctx context.Context, id string) (*domain.Reminder, error) {
	reminder, err := r.reminderRepository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return reminder, nil
}
