package domain

import (
	"time"
)

type Reminder struct {
	ID       string
	HabitID  string
	RemindAt time.Time
	Message  string
}
