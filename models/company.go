package models

import "time"

type Tabler interface {
	TableName() string
}

type Company struct {
	Id        string    `gorm:"default:uuid_generate_v4()" json:"-"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Balance   uint64    `json:"balance"`
	createdAt time.Time `gorm:"default:CURRENT_TIMESTAMP(3)"`
	updatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP(3)"`
}

func (Company) TableName() string {
	return "company"
}
