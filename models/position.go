package models

type Position struct {
	ID        string `gorm:"default:uuid_generate_v4()" json:"id"`
	Name      string `json:"name"`
	Salary    uint32 `json:"salary"`
	createdAt string `gorm:"default:CURRENT_TIMESTAMP(3)"`
	updatedAt string `gorm:"default:CURRENT_TIMESTAMP(3)"`
}
