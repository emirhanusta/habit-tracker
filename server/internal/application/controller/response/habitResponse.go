package response

import "habit-tracker/internal/domain"

type HabitResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserId      string `json:"userId"`
}

func ToHabitResponse(habit *domain.Habit) HabitResponse {
	return HabitResponse{
		Id:          habit.Id,
		Name:        habit.Name,
		Description: habit.Description,
		UserId:      habit.UserId,
	}
}

func ToHabitResponseList(habits []domain.Habit) []HabitResponse {
	var response []HabitResponse

	for _, habit := range habits {
		response = append(response, ToHabitResponse(&habit))
	}

	return response
}
