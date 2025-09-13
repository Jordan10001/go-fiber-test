package critique

import (
	"myapp/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// Handler handles HTTP requests for critiques.
type Handler struct {
	service Service
}

// NewCritiqueHandler creates a new instance of the handler.
func NewCritiqueHandler(svc Service) *Handler {
	return &Handler{
		service: svc,
	}
}

// CreateCritique handles the POST /critiques request.
func (h *Handler) CreateCritique(c *fiber.Ctx) error {
	critique := new(Critique)
	if err := c.BodyParser(critique); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusBadRequest, "Invalid request body")
	}

	if err := h.service.CreateCritique(c.Context(), critique); err != nil {
		return utils.SendErrorResponse(c, fiber.StatusInternalServerError, "Failed to save critique")
	}

	return utils.SendSuccessResponse(c, fiber.StatusCreated, "Critique submitted successfully!")
}