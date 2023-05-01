package models

import "time"

type Transaction struct {
	id        string    `gorm:"default:uuid_generate_v4()"`
	Type      string    `json:"type"`
	Amount    uint64    `json:"amount"`
	Note      string    `json:"note"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP(3)" json:"date"`
	updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP(3)"`
}
