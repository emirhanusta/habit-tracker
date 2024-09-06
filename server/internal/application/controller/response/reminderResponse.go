package response

import (
	"habit-tracker/internal/domain"
	"time"
)

type ReminderResponse struct {
	Id       string    `json:"id"`
	HabitId  string    `json:"habitId"`
	RemindAt time.Time `json:"remindAt"`
	Message  string    `json:"message"`
}

func ToReminderResponse(reminder *domain.Reminder) ReminderResponse {
	return ReminderResponse{
		Id:       reminder.Id,
		HabitId:  reminder.HabitId,
		RemindAt: reminder.RemindAt,
		Message:  reminder.Message,
	}
}

func ToReminderResponseList(reminders []domain.Reminder) []ReminderResponse {
	var response []ReminderResponse

	for _, reminder := range reminders {
		response = append(response, ToReminderResponse(&reminder))
	}

	return response
}
