package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Amount float64 `json:"amount"`
	Status string `json:"status"` // e.g., "pending", "completed"
}