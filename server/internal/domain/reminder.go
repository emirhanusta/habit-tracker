package domain

import (
	"time"
)

type Reminder struct {
	Id       string
	HabitId  string
	RemindAt time.Time
	Message  string
}
