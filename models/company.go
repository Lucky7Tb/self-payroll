package models

import "time"

type Tabler interface {
	TableName() string
}

type Company struct {
	ID        string    `gorm:"default:uuid_generate_v4()" json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Balance   uint64    `json:"balance"`
	createdAt string    `gorm:"default:CURRENT_TIMESTAMP(3)"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP(3)"`
}

func (Company) TableName() string {
	return "company"
}
