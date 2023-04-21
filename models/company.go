package models

type Tabler interface {
	TableName() string
}

type Company struct {
	ID        string `gorm:"primarykey" json:"id"`
	Name      string `gorm:"column:name" json:"name"`
	Address   string `gorm:"column:address" json:"address"`
	Balance   uint32 `gorm:"column:balance" json:"balance"`
	CreatedAt string `gorm:"column:created_at" json:"created_at"`
	UpdatedAt string `gorm:"column:updated_at" json:"updated_at"`
}

func (Company) TableName() string {
	return "company"
}
