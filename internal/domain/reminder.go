package domain

type Reminder struct {
	ID             int    `json:"id"`
	HabitID        int    `json:"habit_id"`
	RemindAt       string `json:"remind_at"`
	RepeatInterval string `json:"repeat_interval"`
}
