package models

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Email  string `json:"email"`
	Message string `json:"message"`
}