package dto

type CreateUpdateCompanyDto struct {
	Name    string `json:"name" validate:"required,printascii"`
	Balance uint64 `json:"balance" validate:"required,numeric"`
	Address string `json:"address" validate:"required,printascii"`
}
