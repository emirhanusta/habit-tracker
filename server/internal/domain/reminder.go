package domain

import (
	"time"
)

type Reminder struct {
	ID        string    `json:"id"`
	HabitID   string    `json:"habit_id"`
	RemindAt  time.Time `json:"remind_at"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}
