package models

import (
	"github.com/jinzhu/gorm"
)

func DBMigration(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &UserPreferences{}, &Reservations{})
	return db
}