package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect(path string) error {
	var err error
	DB, err = gorm.Open(sqlite.Open(path), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic("failed to connect database")
	}
	if err := DB.AutoMigrate(
		&User{},
		&Waf{},
	); err != nil {
		panic("failed to migrate database:" + err.Error())
	}
	return nil
}
