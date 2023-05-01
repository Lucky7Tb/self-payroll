package models

type Tabler interface {
	TableName() string
}

type Company struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Balance   uint32 `json:"balance"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func (Company) TableName() string {
	return "company"
}
