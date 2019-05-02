package models

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	User_name     string `gorm:"not null;unique;size:235" json:"user_name"`
	User_email    string `gorm:"not null;unique;size:235" json:"user_email"`
	User_password string `gorm:"not null;size:235" json:"user_password"`
	User_phone    int    `json:"user_phone"`
	Zip_code      int    `json:"zip_code,omitempty"`
}

type UserPreferences struct {
	gorm.Model
	Prefer_food             string `gorm:"varchar(1000)" json:"prefer_food"`
	Prefer_launch_schedules string `gorm:"varchar(1000)" json:"prefer_launch_schedules"`
	Prefer_restaurant_style string `gorm:"varchar(1000)" json:"prefer_restaurant_style"`
	Prefer_music            string `gorm:"varchar(1000)" json:"prefer_music"`
}

func DBMigration(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &UserPreferences{})
	return db
}
