package response

import "github.com/gofiber/fiber/v2"

// SuccessResponse represents a successful API response
type SuccessResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// ErrorResponse represents an error API response
type ErrorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

// Success sends a successful response
func Success(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	return c.Status(statusCode).JSON(SuccessResponse{
		Status:  statusCode,
		Message: message,
		Data:    data,
	})
}

// Error sends an error response
func Error(c *fiber.Ctx, statusCode int, message string, err string) error {
	return c.Status(statusCode).JSON(ErrorResponse{
		Status:  statusCode,
		Message: message,
		Error:   err,
	})
}
