package user

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4"
)

// Repository defines the contract for user data access.
type Repository interface {
	SaveUser(ctx context.Context, user *User) error
	FindByEmail(ctx context.Context, email string) (*User, error)
}

// pgRepository implements the Repository interface for PostgreSQL.
type pgRepository struct {
	conn *pgx.Conn
}

// NewPostgresRepository creates a new instance of the PostgreSQL user repository.
func NewPostgresRepository(conn *pgx.Conn) Repository {
	return &pgRepository{
		conn: conn,
	}
}

// SaveUser inserts or updates a user in the database.
func (r *pgRepository) SaveUser(ctx context.Context, user *User) error {
	query := `INSERT INTO users (id, email, name, created_at) VALUES ($1, $2, $3, $4) ON CONFLICT (email) DO NOTHING`
	_, err := r.conn.Exec(ctx, query, user.ID, user.Email, user.Name, user.CreatedAt)
	if err != nil {
		return fmt.Errorf("failed to save user: %w", err)
	}
	return nil
}

// FindByEmail finds a user by their email.
func (r *pgRepository) FindByEmail(ctx context.Context, email string) (*User, error) {
	user := &User{}
	query := `SELECT id, email, name, created_at FROM users WHERE email = $1`
	err := r.conn.QueryRow(ctx, query, email).Scan(&user.ID, &user.Email, &user.Name, &user.CreatedAt)
	if err != nil {
		if err == pgx.ErrNoRows {
			return nil, nil // User not found, which is not an error
		}
		return nil, fmt.Errorf("failed to find user by email: %w", err)
	}
	return user, nil
} 	