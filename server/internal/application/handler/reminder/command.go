package reminder

import "time"

type CreateCommand struct {
	HabitId  string
	RemindAt time.Time
	Message  string
}

type UpdateCommand struct {
	Id       string
	HabitId  string
	RemindAt time.Time
	Message  string
}
