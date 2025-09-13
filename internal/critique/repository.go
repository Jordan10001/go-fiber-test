package critique

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

// Repository defines the interface for data access.
type Repository interface {
	CreateCritique(ctx context.Context, critique *Critique) error
}

// critiqueRepository implements the Repository interface.
type critiqueRepository struct {
	collection *mongo.Collection
}

// NewCritiqueRepository creates a new instance of the repository.
func NewCritiqueRepository(coll *mongo.Collection) Repository {
	return &critiqueRepository{
		collection: coll,
	}
}

// CreateCritique inserts a new critique into the database.
func (r *critiqueRepository) CreateCritique(ctx context.Context, critique *Critique) error {
	_, err := r.collection.InsertOne(ctx, critique)
	return err
}