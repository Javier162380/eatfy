package models

import (
	"github.com/jinzhu/gorm"
)

type UserLogin struct {
	UserIdentifier string `json:"user_identifier"`
	UserPassword   string `json:"user_password"`
}

type User struct {
	gorm.Model
	UserName     string `gorm:"not null;unique;size:235" json:"user_name"`
	UserEmail    string `gorm:"not null;unique;size:235" json:"user_email"`
	UserPassword string `gorm:"not null;size:235" json:"user_password"`
	UserPhone    int    `json:"user_phone"`
	ZipCode      int    `json:"zip_code,omitempty"`
}

type UserPreferences struct {
	gorm.Model
	PreferFood            string `gorm:"varchar(1000)" json:"prefer_food"`
	PreferLaunchSchedules string `gorm:"varchar(1000)" json:"prefer_launch_schedules"`
	PreferRestaurantStyle string `gorm:"varchar(1000)" json:"prefer_restaurant_style"`
	PreferMusic           string `gorm:"varchar(1000)" json:"prefer_music"`
}

func DBMigration(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &UserPreferences{})
	return db
}
