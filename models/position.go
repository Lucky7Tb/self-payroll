package models

type Position struct {
	ID        string `gorm:"primarykey" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	Salary    int64  `gorm:"column:salary" json:"salary"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}
