package handler

import (
	"strconv"
	"time"
	"workshop3/internal/models"
	"workshop3/pkg/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Handler holds dependencies for user handlers
type Handler struct {
	db *gorm.DB
}

// New creates a new handler instance
func New(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

// CreateUser handles POST /users - Creates a new user
func (h *Handler) CreateUser(c *fiber.Ctx) error {
	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}

	user.CreatedAt = time.Now().Unix()
	user.UpdatedAt = time.Now().Unix()

	result := h.db.Create(user)
	if result.Error != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed to create user", result.Error.Error())
	}

	return response.Success(c, fiber.StatusCreated, "User created successfully", user)
}

// GetAllUsers handles GET /users - Retrieves all users
func (h *Handler) GetAllUsers(c *fiber.Ctx) error {
	var users []models.User
	result := h.db.Find(&users)
	if result.Error != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve users", result.Error.Error())
	}

	return response.Success(c, fiber.StatusOK, "Users retrieved successfully", users)
}

// GetUserByID handles GET /users/:id - Retrieves a user by ID
func (h *Handler) GetUserByID(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	user := new(models.User)
	result := h.db.First(user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return response.Error(c, fiber.StatusNotFound, "User not found", "")
		}
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve user", result.Error.Error())
	}

	return response.Success(c, fiber.StatusOK, "User retrieved successfully", user)
}

// UpdateUser handles PUT /users/:id - Updates a user
func (h *Handler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	user := new(models.User)
	if err := c.BodyParser(user); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}

	user.UpdatedAt = time.Now().Unix()

	result := h.db.Where("id = ?", id).Updates(user)
	if result.Error != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed to update user", result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return response.Error(c, fiber.StatusNotFound, "User not found", "")
	}

	return response.Success(c, fiber.StatusOK, "User updated successfully", user)
}

// DeleteUser handles DELETE /users/:id - Deletes a user
func (h *Handler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.ParseUint(c.Params("id"), 10, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	result := h.db.Delete(&models.User{}, id)
	if result.Error != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete user", result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return response.Error(c, fiber.StatusNotFound, "User not found", "")
	}

	return response.Success(c, fiber.StatusOK, "User deleted successfully", nil)
}
