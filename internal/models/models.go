package models

import (
	"time"
)

type Article struct {
	ID        int        `json:"id" db:"id"`
	Name      string     `json:"name" db:"name"`
	Content   string     `json:"content" db:"content"`
	UserID    int        `json:"user_id" db:"user_id"`
	CreatedAt time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" db:"deleted_at"`
}

type User struct {
	ID        int       `json:"id" db:"id"`
	FullName  string    `json:"full_name,omitempty" db:"full_name"`
	Login     string    `json:"login" db:"login"`
	Password  string    `json:"-" db:"password"`
	CreatedAt time.Time `json:"created_at,omitempty" db:"created_at"`
}

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
