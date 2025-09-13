package server

import (
	"myapp/internal/config"
	"myapp/internal/database"
	"myapp/internal/critique"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
    "github.com/gofiber/fiber/v2/middleware/cors" // Tambahkan ini
)

// Server holds the Fiber app and configurations.
type Server struct {
	app *fiber.App
	cfg *config.Config
}

// NewServer creates a new instance of the server.
func NewServer(cfg *config.Config) *Server {
	app := fiber.New()
	app.Use(logger.New())
    app.Use(cors.New()) // Tambahkan ini untuk mengaktifkan CORS

	return &Server{
		app: app,
		cfg: cfg,
	}
}

// SetupRoutes configures the application routes.
func (s *Server) SetupRoutes() {
	// Initialize database connection
	database.ConnectDB(s.cfg)

	// Dependency Injection for the critique module
	critiqueCollection := database.GetCollection(s.cfg, "critiques")
	critiqueRepo := critique.NewCritiqueRepository(critiqueCollection)
	critiqueService := critique.NewCritiqueService(critiqueRepo)
	critiqueHandler := critique.NewCritiqueHandler(critiqueService)

	// API routes
	api := s.app.Group("/api/v1")
	api.Post("/critiques", critiqueHandler.CreateCritique)
}

// Start runs the server.
func (s *Server) Start() {
	s.SetupRoutes()
	s.app.Listen(":" + s.cfg.AppPort)
}