package routes

import (
	"workshop3/internal/handler"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// RegisterUserRoutes registers all user-related routes
func RegisterUserRoutes(app *fiber.App, db *gorm.DB) {
	h := handler.New(db)

	users := app.Group("/api/v1/users")
	users.Post("", h.CreateUser)
	users.Get("", h.GetAllUsers)
	users.Get("/:id", h.GetUserByID)
	users.Put("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)

	// Register transfer routes
	transfers := app.Group("/api/v1/transfers")
	transfers.Post("", h.CreateTransfer)
	transfers.Get("", h.GetTransfers)
	transfers.Get("/:id", h.GetTransferByID)
}
