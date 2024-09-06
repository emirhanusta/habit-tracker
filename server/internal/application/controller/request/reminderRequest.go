package request

import (
	"habit-tracker/internal/application/handler/reminder"
	"time"
)

type ReminderCreateRequest struct {
	HabitId  string    `json:"habitId"`
	RemindAt time.Time `json:"remindAt"`
	Message  string    `json:"message"`
}

type ReminderUpdateRequest struct {
	Id       string    `json:"id"`
	RemindAt time.Time `json:"remindAt"`
	Message  string    `json:"message"`
}

func (r *ReminderCreateRequest) ToCommand() reminder.CreateCommand {
	return reminder.CreateCommand{
		HabitId:  r.HabitId,
		RemindAt: r.RemindAt,
		Message:  r.Message,
	}
}

func (r *ReminderUpdateRequest) ToCommand() reminder.UpdateCommand {
	return reminder.UpdateCommand{
		Id:       r.Id,
		RemindAt: r.RemindAt,
		Message:  r.Message,
	}
}
