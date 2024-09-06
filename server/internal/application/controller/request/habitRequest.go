package request

import "habit-tracker/internal/application/handler/habit"

type HabitCreateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description"`
	UserId      string `json:"userId" validate:"required"`
}

type HabitUpdateRequest struct {
	Id          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (h *HabitUpdateRequest) ToCommand() habit.UpdateCommand {
	return habit.UpdateCommand{
		Id:          h.Id,
		Name:        h.Name,
		Description: h.Description,
	}
}

func (h *HabitCreateRequest) ToCommand() habit.CreateCommand {
	return habit.CreateCommand{
		Name:        h.Name,
		Description: h.Description,
		UserId:      h.UserId,
	}
}
