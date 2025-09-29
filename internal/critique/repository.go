package critique

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// Repository defines the interface for critique data access.
type Repository interface {
	CreateCritique(ctx context.Context, critique *Critique) error
}

// pgRepository implements the Repository interface for PostgreSQL.
type pgRepository struct {
	conn *pgx.Conn
}

// NewPostgresRepository creates a new instance of the PostgreSQL repository.
func NewPostgresRepository(conn *pgx.Conn) Repository {
	return &pgRepository{
		conn: conn,
	}
}

// CreateCritique inserts a new critique into the database.
func (r *pgRepository) CreateCritique(ctx context.Context, critique *Critique) error {
	query := `INSERT INTO critiques (title, content, created_at) VALUES ($1, $2, $3)`
	_, err := r.conn.Exec(ctx, query, critique.Title, critique.Content, critique.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to insert critique: %v", err)
	}
	return nil
}
