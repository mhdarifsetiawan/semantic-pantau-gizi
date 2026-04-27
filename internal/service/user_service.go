package service

import (
	"context"

	"github.com/mhdarifsetiawan/semantic-pantau-gizi/internal/domain"
	apperrors "github.com/mhdarifsetiawan/semantic-pantau-gizi/pkg/errors"
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*domain.User, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, apperrors.FromRuntime(err)
	}
	return users, nil
}

func (s *UserService) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	if id == "" {
		return nil, apperrors.BadRequest("user id cannot be empty")
	}

	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, apperrors.FromRuntime(err)
	}
	return user, nil
}
