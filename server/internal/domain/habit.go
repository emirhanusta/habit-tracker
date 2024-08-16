package domain

import (
	"time"
)

type Habit struct {
	ID          string    `json:"id"`      // UUID olarak tanımlı string
	UserID      string    `json:"user_id"` // User tablosuna referans, UUID olarak tanımlı string
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
