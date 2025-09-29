package server

import (
	"myapp/internal/auth"
	"myapp/internal/config"
	"myapp/internal/critique"
	"myapp/internal/database"
	"myapp/internal/user"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
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
	app.Use(cors.New())

	return &Server{
		app: app,
		cfg: cfg,
	}
}

// SetupRoutes configures the application routes.
func (s *Server) SetupRoutes() {
	// Inisialisasi koneksi database
	database.ConnectDB(s.cfg)

	// Inisialisasi konfigurasi Google OAuth
	auth.InitGoogleOAuth(s.cfg.GoogleClientID, s.cfg.GoogleSecret)

	// Dependency Injection untuk modul kritik
	critiqueRepo := critique.NewPostgresRepository(database.DB)
	critiqueService := critique.NewCritiqueService(critiqueRepo)
	critiqueHandler := critique.NewCritiqueHandler(critiqueService)

	// Dependency Injection untuk modul user
	userRepo := user.NewPostgresRepository(database.DB)
	userService := user.NewUserService(userRepo)

	// Dependency Injection untuk modul otentikasi
	authHandler := auth.NewAuthHandler(s.cfg, userService)

	// API routes
	api := s.app.Group("/api/v1")
	api.Post("/critiques", critiqueHandler.CreateCritique)

	// Rute Otentikasi Google
	s.app.Get("/auth/google", authHandler.HandleGoogleLogin)
	s.app.Get("/auth/google/callback", authHandler.HandleGoogleCallback)
}

// Start runs the server.
func (s *Server) Start() {
	s.SetupRoutes()
	s.app.Listen(":" + s.cfg.AppPort)
}
