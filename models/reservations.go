package models

import (
	"github.com/jinzhu/gorm"
)

type Reservations struct {
	gorm.Model
	UserName string `gorm:"association_foreignkey:UserName:not null;unique;size:235" json:"user_name"`
	RestaurantName string `gorm:"association_foreignkey:UserName:not null;unique;size:235" json:"restaurante_name"`
	PickUpTime string `json:"time_schedule"`
	ReservationDate string `json:"reservation_date"`
	MealType string `json:"meal_type"`
	Notes string `json:"notes,omitempty"`
	PickUpFood bool `sql:"default:false"`
}
