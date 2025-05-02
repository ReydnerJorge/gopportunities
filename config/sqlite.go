package config

import (
	"os"

	"github.com/Reydner96/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"
	// Check if  the database file exist
	_, err := os.Stat(dbPath)
	if err != nil {
		if os.IsNotExist(err) {
			logger.Info("database file not found, creating...")

			// Create the database file and directory
			if err = os.MkdirAll("./db", os.ModePerm); err != nil {
				return nil, err
			}
			file, err := os.Create(dbPath)
			if err != nil {
				return nil, err
			}
			defer file.Close()
		} else {
			return nil, err
		}
	}

	// Create DB and Connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	// Migrate the Schema
	if err = db.AutoMigrate(&schemas.Opening{}); err != nil {
		logger.Errorf("sqlite automigration error: %v", err)
		return nil, err
	}

	// Return the DB
	return db, nil
}
