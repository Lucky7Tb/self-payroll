package dto

type TopupCompanyBalanceDto struct {
	Balance uint64 `json:"balance" validate:"required,numeric"`
}
