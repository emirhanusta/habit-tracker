package domain

import (
	"time"
)

type User struct {
	ID           string    `json:"id"` // UUID olarak tanımlı string
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	CreatedAt    time.Time `json:"created_at"`
}
