package utils

import "github.com/gofiber/fiber/v2"

// SuccessResponse defines a standard success response structure.
type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

// ErrorResponse defines a standard error response structure.
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

// SendSuccessResponse sends a standardized success response.
func SendSuccessResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(SuccessResponse{
		Status:  "success",
		Message: message,
	})
}

// SendErrorResponse sends a standardized error response.
func SendErrorResponse(c *fiber.Ctx, statusCode int, message string) error {
	return c.Status(statusCode).JSON(ErrorResponse{
		Status:  "error",
		Message: message,
	})
}