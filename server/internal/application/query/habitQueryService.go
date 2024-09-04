package query

import (
	"context"
	"habit-tracker/internal/application/repository"
	"habit-tracker/internal/domain"
)

type IHabitQueryService interface {
	GetAllByUserId(ctx context.Context, id string) ([]domain.Habit, error)
	GetById(ctx context.Context, id string) (*domain.Habit, error)
}

type habitQueryService struct {
	habitRepository repository.IHabitRepository
}

func NewHabitQueryService(habitRepository repository.IHabitRepository) IHabitQueryService {
	return &habitQueryService{
		habitRepository: habitRepository,
	}
}

func (h *habitQueryService) GetAllByUserId(ctx context.Context, id string) ([]domain.Habit, error) {
	habits, err := h.habitRepository.GetAllByUserId(ctx, id)

	if err != nil {
		return nil, err
	}

	return habits, nil
}

func (h *habitQueryService) GetById(ctx context.Context, id string) (*domain.Habit, error) {
	habit, err := h.habitRepository.GetById(ctx, id)

	if err != nil {
		return nil, err
	}

	return habit, nil
}
