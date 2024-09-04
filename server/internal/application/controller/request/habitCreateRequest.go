package request

import "habit-tracker/internal/application/handler/habit"

type HabitCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	UserId      string `json:"userId" validate:"required"`
}

func (h *HabitCreateRequest) ToCommand() habit.CreateCommand {
	return habit.CreateCommand{
		Name:        h.Name,
		Description: h.Description,
		UserId:      h.UserId,
	}
}
