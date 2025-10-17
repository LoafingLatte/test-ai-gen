package handler

import (
	"crypto/rand"
	"encoding/hex"
	"strconv"
	"time"
	"workshop3/internal/models"
	"workshop3/pkg/response"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// TransferCreateRequest represents the request for creating a transfer
type TransferCreateRequest struct {
	FromUserID uint   `json:"fromUserId" binding:"required"`
	ToUserID   uint   `json:"toUserId" binding:"required"`
	Amount     int64  `json:"amount" binding:"required,min=1"`
	Note       string `json:"note,omitempty"`
}

// TransferResponse represents the API response for a transfer
type TransferResponse struct {
	Transfer *TransferResponseData `json:"transfer"`
}

// TransferResponseData represents the transfer data in response
type TransferResponseData struct {
	IdemKey     string  `json:"idemKey"`
	TransferID  uint    `json:"transferId,omitempty"`
	FromUserID  uint    `json:"fromUserId"`
	ToUserID    uint    `json:"toUserId"`
	Amount      int64   `json:"amount"`
	Status      string  `json:"status"`
	Note        *string `json:"note,omitempty"`
	CreatedAt   string  `json:"createdAt"`
	UpdatedAt   string  `json:"updatedAt"`
	CompletedAt *string `json:"completedAt,omitempty"`
	FailReason  *string `json:"failReason,omitempty"`
}

// TransferListResponse represents the response for listing transfers
type TransferListResponse struct {
	Data     []TransferResponseData `json:"data"`
	Page     int                    `json:"page"`
	PageSize int                    `json:"pageSize"`
	Total    int64                  `json:"total"`
}

// generateIdempotencyKey generates a random idempotency key
func generateIdempotencyKey() (string, error) {
	b := make([]byte, 16)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}

// convertTransferToResponse converts a Transfer model to TransferResponseData
func convertTransferToResponse(t *models.Transfer) *TransferResponseData {
	response := &TransferResponseData{
		IdemKey:    t.IdempotencyKey,
		TransferID: t.ID,
		FromUserID: t.FromUserID,
		ToUserID:   t.ToUserID,
		Amount:     t.Amount,
		Status:     string(t.Status),
		Note:       t.Note,
		CreatedAt:  time.Unix(t.CreatedAt, 0).UTC().Format(time.RFC3339),
		UpdatedAt:  time.Unix(t.UpdatedAt, 0).UTC().Format(time.RFC3339),
	}

	if t.CompletedAt != nil {
		completedAtStr := time.Unix(*t.CompletedAt, 0).UTC().Format(time.RFC3339)
		response.CompletedAt = &completedAtStr
	}

	response.FailReason = t.FailReason

	return response
}

// CreateTransfer handles POST /transfers - Creates a new transfer
func (h *Handler) CreateTransfer(c *fiber.Ctx) error {
	var req TransferCreateRequest
	if err := c.BodyParser(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}

	// Validate request
	if req.FromUserID == 0 || req.ToUserID == 0 || req.Amount <= 0 {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request parameters", "fromUserId, toUserId, and amount must be valid")
	}

	// Validate amount (max 2 decimal places, max 200 cents = 2.00)
	if req.Amount > 200 {
		return response.Error(c, fiber.StatusConflict, "Amount exceeds maximum", "Maximum transfer amount is 2.00")
	}

	// Check if from and to users are different
	if req.FromUserID == req.ToUserID {
		return response.Error(c, fiber.StatusUnprocessableEntity, "Cannot transfer to yourself", "fromUserId and toUserId must be different")
	}

	// Check if users exist
	var fromUser models.User
	if err := h.db.First(&fromUser, req.FromUserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.Error(c, fiber.StatusNotFound, "From user not found", "")
		}
		return response.Error(c, fiber.StatusInternalServerError, "Database error", err.Error())
	}

	var toUser models.User
	if err := h.db.First(&toUser, req.ToUserID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.Error(c, fiber.StatusNotFound, "To user not found", "")
		}
		return response.Error(c, fiber.StatusInternalServerError, "Database error", err.Error())
	}

	// Check if from user has enough points
	if fromUser.Points < req.Amount {
		return response.Error(c, fiber.StatusConflict, "Insufficient points", "User does not have enough points to transfer")
	}

	// Generate idempotency key
	idemKey, err := generateIdempotencyKey()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to generate idempotency key", err.Error())
	}

	// Create transfer record
	now := time.Now().Unix()
	transfer := &models.Transfer{
		FromUserID:     req.FromUserID,
		ToUserID:       req.ToUserID,
		Amount:         req.Amount,
		Status:         models.StatusCompleted,
		IdempotencyKey: idemKey,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	if req.Note != "" {
		transfer.Note = &req.Note
	}

	// Start transaction
	tx := h.db.Begin()

	// Create transfer
	if err := tx.Create(transfer).Error; err != nil {
		tx.Rollback()
		return response.Error(c, fiber.StatusBadRequest, "Failed to create transfer", err.Error())
	}

	// Update user points
	if err := tx.Model(&models.User{}).Where("id = ?", req.FromUserID).
		Update("points", gorm.Expr("points - ?", req.Amount)).Error; err != nil {
		tx.Rollback()
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update from user points", err.Error())
	}

	if err := tx.Model(&models.User{}).Where("id = ?", req.ToUserID).
		Update("points", gorm.Expr("points + ?", req.Amount)).Error; err != nil {
		tx.Rollback()
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update to user points", err.Error())
	}

	// Create point ledger entries
	fromLedger := &models.PointLedger{
		UserID:       req.FromUserID,
		Change:       -req.Amount,
		BalanceAfter: fromUser.Points - req.Amount,
		EventType:    models.EventTransferOut,
		TransferID:   &transfer.ID,
		CreatedAt:    now,
	}

	toLedger := &models.PointLedger{
		UserID:       req.ToUserID,
		Change:       req.Amount,
		BalanceAfter: toUser.Points + req.Amount,
		EventType:    models.EventTransferIn,
		TransferID:   &transfer.ID,
		CreatedAt:    now,
	}

	if err := tx.Create(fromLedger).Error; err != nil {
		tx.Rollback()
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create from ledger", err.Error())
	}

	if err := tx.Create(toLedger).Error; err != nil {
		tx.Rollback()
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create to ledger", err.Error())
	}

	if err := tx.Commit().Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to commit transaction", err.Error())
	}

	completedAtStr := time.Unix(now, 0).UTC().Format(time.RFC3339)
	transfer.CompletedAt = &now

	responseData := convertTransferToResponse(transfer)
	responseData.CompletedAt = &completedAtStr

	c.Set("Idempotency-Key", idemKey)
	return response.Success(c, fiber.StatusCreated, "Transfer created successfully", &TransferResponse{Transfer: responseData})
}

// GetTransfers handles GET /transfers - Lists transfers for a user
func (h *Handler) GetTransfers(c *fiber.Ctx) error {
	userIDStr := c.Query("userId")
	pageStr := c.Query("page", "1")
	pageSizeStr := c.Query("pageSize", "20")

	if userIDStr == "" {
		return response.Error(c, fiber.StatusBadRequest, "userId query parameter is required", "")
	}

	userID, err := strconv.ParseUint(userIDStr, 10, 64)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid userId", err.Error())
	}

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		page = 1
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 || pageSize > 200 {
		pageSize = 20
	}

	// Check if user exists
	var user models.User
	if err := h.db.First(&user, userID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.Error(c, fiber.StatusNotFound, "User not found", "")
		}
		return response.Error(c, fiber.StatusInternalServerError, "Database error", err.Error())
	}

	var transfers []models.Transfer
	var total int64

	// Get transfers where user is sender or receiver
	query := h.db.Where("from_user_id = ? OR to_user_id = ?", userID, userID)
	if err := query.Model(&models.Transfer{}).Count(&total).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to count transfers", err.Error())
	}

	offset := (page - 1) * pageSize
	if err := query.Offset(offset).Limit(pageSize).Order("created_at DESC").Find(&transfers).Error; err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve transfers", err.Error())
	}

	responseList := make([]TransferResponseData, 0, len(transfers))
	for i := range transfers {
		responseList = append(responseList, *convertTransferToResponse(&transfers[i]))
	}

	return response.Success(c, fiber.StatusOK, "Transfers retrieved successfully", &TransferListResponse{
		Data:     responseList,
		Page:     page,
		PageSize: pageSize,
		Total:    total,
	})
}

// GetTransferByID handles GET /transfers/:id - Gets transfer by idempotency key
func (h *Handler) GetTransferByID(c *fiber.Ctx) error {
	id := c.Params("id")
	if id == "" {
		return response.Error(c, fiber.StatusBadRequest, "Transfer ID is required", "")
	}

	var transfer models.Transfer
	if err := h.db.Where("idempotency_key = ?", id).First(&transfer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return response.Error(c, fiber.StatusNotFound, "Transfer not found", "")
		}
		return response.Error(c, fiber.StatusInternalServerError, "Database error", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Transfer retrieved successfully", &TransferResponse{Transfer: convertTransferToResponse(&transfer)})
}
