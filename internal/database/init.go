package database

import (
	"workshop3/internal/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Init initializes the SQLite database connection and runs migrations
func Init(dbPath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate the models
	if err := db.AutoMigrate(
		&models.User{},
		&models.Transfer{},
		&models.PointLedger{},
	); err != nil {
		return nil, err
	}

	return db, nil
}
