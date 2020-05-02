package models

import (
	"github.com/jinzhu/gorm"
)

// User - database
type User struct {
	gorm.Model
	Name   string      `gorm:"not null" json:"name"`
	Email  string      `gorm:"unique;not null" json:"email"`
	Images []ImageLink `gorm:"foreignkey:UserID"`
}

// ImageLink - database
type ImageLink struct {
	gorm.Model
	Link                string `gorm:"unique_index:idx_name_code;not null" json:"link"`
	AssistantPrediction string `gorm:"not null" json:"assistant_prediction"`
	DoctorPrediction    string `json:"doctor_prediction"`
	UserID              uint   `gorm:"unique_index:idx_name_code;not null" json:"user_id"`
}
