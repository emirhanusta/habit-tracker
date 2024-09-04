package domain

import (
	"time"
)

type Reminder struct {
	Id       string
	HabitID  string
	RemindAt time.Time
	Message  string
}
