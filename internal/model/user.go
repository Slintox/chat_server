package model

import "time"

type UserRole int

// User описывает модель пользователя
type User struct {
	Username  string // Unique
	Email     string
	Role      UserRole
	CreatedAt time.Time
	UpdatedAt time.Time
}
