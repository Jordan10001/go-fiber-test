package auth

import (
	"encoding/json"
	"fmt"
	"log"
	"myapp/internal/config"
	"myapp/internal/user"

	"github.com/gofiber/fiber/v2"
)

// Struktur untuk menampung data pengguna dari Google
type GoogleUser struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

// Handler handles OAuth requests.
type Handler struct {
	cfg        *config.Config
	userService user.Service
}

// NewAuthHandler creates a new auth handler.
func NewAuthHandler(cfg *config.Config, userService user.Service) *Handler {
	return &Handler{
		cfg:        cfg,
		userService: userService,
	}
}

// HandleGoogleLogin mengarahkan pengguna ke halaman login Google.
func (h *Handler) HandleGoogleLogin(c *fiber.Ctx) error {
	// Untuk demo, kita gunakan state statis, tapi di produksi harus acak dan disimpan di sesi.
	state := "random_state_string"
	url := GetGoogleOAuthURL(state)
	return c.Redirect(url, fiber.StatusTemporaryRedirect)
}

// HandleGoogleCallback menangani panggilan balik dari Google.
func (h *Handler) HandleGoogleCallback(c *fiber.Ctx) error {
	state := c.Query("state")
	if state != "random_state_string" {
		return c.Status(fiber.StatusBadRequest).SendString("State parameter does not match.")
	}

	code := c.Query("code")
	if code == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Code parameter not found.")
	}

	token, err := GetGoogleTokens(c.Context(), code)
	if err != nil {
		log.Printf("Failed to exchange code for token: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get token.")
	}

	// Langkah 1: Dapatkan informasi pengguna dari Google
	client := GoogleConfig.Client(c.Context(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		log.Printf("Failed to get user info: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to get user info.")
	}
	defer resp.Body.Close()

	var googleUser GoogleUser
	if err := json.NewDecoder(resp.Body).Decode(&googleUser); err != nil {
		log.Printf("Failed to decode user info: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to decode user info.")
	}

	// Langkah 2: Simpan atau perbarui data pengguna di database
	appUser := &user.User{
		ID:    googleUser.ID,
		Email: googleUser.Email,
		Name:  googleUser.Name,
	}
	savedUser, err := h.userService.GetOrCreateUser(c.Context(), appUser)
	if err != nil {
		log.Printf("Failed to save user: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Failed to save user.")
	}
	
	log.Printf("User logged in: %+v\n", savedUser)

	// Langkah 3: Arahkan kembali ke frontend dengan token
	frontendURL := fmt.Sprintf("%s?token=%s", h.cfg.FrontendURL, token.AccessToken)
	return c.Redirect(frontendURL, fiber.StatusFound)
}