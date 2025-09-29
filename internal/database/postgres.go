package database

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v4"
	"myapp/internal/config"
)

// DB represents the database connection.
var DB *pgx.Conn

// ConnectDB establishes a connection to PostgreSQL.
func ConnectDB(cfg *config.Config) {
	conn, err := pgx.Connect(context.Background(), cfg.PostgreSQLURI)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = conn.Ping(ctx)
	if err != nil {
		log.Fatalf("Could not ping database: %v\n", err)
	}

	DB = conn
	log.Println("Connected to PostgreSQL successfully!")
}

// CloseDB closes the database connection.
func CloseDB() {
	if DB != nil {
		DB.Close(context.Background())
		log.Println("Database connection closed.")
	}
}