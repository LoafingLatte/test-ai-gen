package models

// User represents a user in the system with membership information
type User struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	MembershipID    string `gorm:"uniqueIndex" json:"membership_id"`
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	Phone           string `json:"phone"`
	Email           string `gorm:"uniqueIndex" json:"email"`
	MembershipDate  string `json:"membership_date"`
	MembershipLevel string `json:"membership_level"`
	Points          int64  `json:"points"`
	CreatedAt       int64  `json:"created_at"`
	UpdatedAt       int64  `json:"updated_at"`
}

// TableName specifies the table name for the User model
func (User) TableName() string {
	return "users"
}
