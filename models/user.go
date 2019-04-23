package models

import (
	"github.com/jinzhu/gorm"
	"github.com/lib/pq"
)

type User struct {
	gorm.Model
	user_name  string `gorm:"size:235"`
	user_email string `gorm:"size:235"`
	user_phone int
	zip_code   int
}

type UserPreferences struct {
	gorm.Model
	prefer_food             pq.StringArray `gorm:"varchar(1000)"`
	prefer_launch_schedules pq.StringArray `gorm:"varchar(1000)"`
	prefer_restaurant_style pq.StringArray `gorm:"varchar(1000)"`
	prefer_music            pq.StringArray `gorm:"varchar(1000)"`
}

func DBMigration(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&User{}, &UserPreferences{})
	return db
}
