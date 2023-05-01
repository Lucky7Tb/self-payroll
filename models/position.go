package models

import "time"

type Position struct {
	ID        string    `gorm:"default:uuid_generate_v4()" json:"id"`
	Name      string    `json:"name,omitempty"`
	Salary    uint64    `json:"salary,omitempty"`
	createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP(3)"`
	updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP(3)"`
}
