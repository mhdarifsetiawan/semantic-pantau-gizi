package domain

import (
	"context"
	"time"
)

type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

type UserRepository interface {
	FindAll(ctx context.Context) ([]*User, error)
	FindByID(ctx context.Context, id string) (*User, error)
}
