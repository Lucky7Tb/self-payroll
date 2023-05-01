package models

import "time"

type User struct {
	ID         string    `gorm:"default:uuid_generate_v4()" json:"id"`
	PositionId string    `json:"position_id"`
	EmployeeId string    `json:"employee_id"`
	SecretId   string    `json:"secret_id,omitempty"`
	Name       string    `json:"name"`
	Email      string    `json:"email,omitempty"`
	Phone      string    `json:"phone"`
	Address    string    `json:"address,omitempty"`
	Position   Position  `gorm:"foreignKey:PositionId" json:"position"`
	createdAt  time.Time `gorm:"default:CURRENT_TIMESTAMP(3)"`
	updatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP(3)"`
}
