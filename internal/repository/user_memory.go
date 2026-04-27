package repository

import (
	"context"
	"time"

	"github.com/mhdarifsetiawan/semantic-pantau-gizi/internal/domain"
	apperrors "github.com/mhdarifsetiawan/semantic-pantau-gizi/pkg/errors"
)

type userMemoryRepo struct {
	data map[string]*domain.User
}

func NewUserMemoryRepository() domain.UserRepository {
	now := time.Now()
	data := map[string]*domain.User{
		"1": {
			ID:        "1",
			Name:      "Ahmad Fauzi",
			Email:     "ahmad@example.com",
			Role:      "admin",
			CreatedAt: now.Add(-72 * time.Hour),
		},
		"2": {
			ID:        "2",
			Name:      "Siti Rahma",
			Email:     "siti@example.com",
			Role:      "user",
			CreatedAt: now.Add(-24 * time.Hour),
		},
		"3": {
			ID:        "3",
			Name:      "Budi Santoso",
			Email:     "budi@example.com",
			Role:      "user",
			CreatedAt: now,
		},
	}
	return &userMemoryRepo{data: data}
}

func (r *userMemoryRepo) FindAll(ctx context.Context) ([]*domain.User, error) {
	users := make([]*domain.User, 0, len(r.data))
	for _, u := range r.data {
		users = append(users, u)
	}
	return users, nil
}

func (r *userMemoryRepo) FindByID(ctx context.Context, id string) (*domain.User, error) {
	user, ok := r.data[id]
	if !ok {
		return nil, apperrors.NotFound("user")
	}
	return user, nil
}
