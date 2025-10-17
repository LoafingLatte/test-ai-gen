package models

// TransferStatus represents the status of a transfer
type TransferStatus string

const (
	StatusPending    TransferStatus = "pending"
	StatusProcessing TransferStatus = "processing"
	StatusCompleted  TransferStatus = "completed"
	StatusFailed     TransferStatus = "failed"
	StatusCancelled  TransferStatus = "cancelled"
	StatusReversed   TransferStatus = "reversed"
)

// Transfer represents a transfer record
type Transfer struct {
	ID             uint           `gorm:"primaryKey" json:"transfer_id"`
	FromUserID     uint           `json:"from_user_id" binding:"required"`
	ToUserID       uint           `json:"to_user_id" binding:"required"`
	Amount         int64          `json:"amount" binding:"required,min=1"`
	Status         TransferStatus `json:"status"`
	Note           *string        `json:"note"`
	IdempotencyKey string         `gorm:"uniqueIndex" json:"idemKey"`
	CreatedAt      int64          `json:"created_at"`
	UpdatedAt      int64          `json:"updated_at"`
	CompletedAt    *int64         `json:"completed_at"`
	FailReason     *string        `json:"fail_reason"`
}

// TableName specifies the table name for the Transfer model
func (Transfer) TableName() string {
	return "transfers"
}

// EventType represents the type of point ledger event
type EventType string

const (
	EventTransferOut EventType = "transfer_out"
	EventTransferIn  EventType = "transfer_in"
	EventAdjust      EventType = "adjust"
	EventEarn        EventType = "earn"
	EventRedeem      EventType = "redeem"
)

// PointLedger represents a point ledger entry (append-only)
type PointLedger struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	UserID       uint      `json:"user_id" binding:"required"`
	Change       int64     `json:"change"` // positive for incoming, negative for outgoing
	BalanceAfter int64     `json:"balance_after"`
	EventType    EventType `json:"event_type"`
	TransferID   *uint     `json:"transfer_id"`
	Reference    *string   `json:"reference"`
	Metadata     *string   `json:"metadata"` // JSON text
	CreatedAt    int64     `json:"created_at"`
}

// TableName specifies the table name for the PointLedger model
func (PointLedger) TableName() string {
	return "point_ledger"
}
