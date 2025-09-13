package critique

import (
	"context"
	"time"
)

// Service defines the business logic for critiques.
type Service interface {
	CreateCritique(ctx context.Context, critique *Critique) error
}

// critiqueService implements the Service interface.
type critiqueService struct {
	repo Repository
}

// NewCritiqueService creates a new instance of the service.
func NewCritiqueService(repo Repository) Service {
	return &critiqueService{
		repo: repo,
	}
}

// CreateCritique handles the creation of a new critique.
func (s *critiqueService) CreateCritique(ctx context.Context, critique *Critique) error {
	critique.CreatedAt = time.Now().Unix()
	return s.repo.CreateCritique(ctx, critique)
}