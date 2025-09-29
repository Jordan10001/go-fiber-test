package user

import (
	"context"
	"time"
)

// Service defines the business logic for users.
type Service interface {
	GetOrCreateUser(ctx context.Context, googleUser *User) (*User, error)
}

// userService implements the Service interface.
type userService struct {
	repo Repository
}

// NewUserService creates a new instance of the user service.
func NewUserService(repo Repository) Service {
	return &userService{
		repo: repo,
	}
}

// GetOrCreateUser checks if a user exists, and if not, saves them.
func (s *userService) GetOrCreateUser(ctx context.Context, googleUser *User) (*User, error) {
	existingUser, err := s.repo.FindByEmail(ctx, googleUser.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return existingUser, nil
	}

	googleUser.CreatedAt = time.Now()
	err = s.repo.SaveUser(ctx, googleUser)
	if err != nil {
		return nil, err
	}

	return googleUser, nil
}