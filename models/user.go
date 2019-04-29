package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	User_name  string `gorm:"not null;unique;size:235"`
	User_email string `gorm:"not null;unique;size:235"`
	User_phone int
	Zip_code   int
}

type UserPreferences struct {
	gorm.Model
	Prefer_food             string `gorm:"varchar(1000)"`
	Prefer_launch_schedules string `gorm:"varchar(1000)"`
	Prefer_restaurant_style string `gorm:"varchar(1000)"`
	Prefer_music            string `gorm:"varchar(1000)"`
}

func DBMigration(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &UserPreferences{})
	return db
}
