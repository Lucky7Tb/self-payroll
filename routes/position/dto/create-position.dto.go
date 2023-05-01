package dto

type CreatePositionDto struct {
	Name   string `json:"name" validate:"required"`
	Salary uint64 `json:"salary" validate:"required,numeric"`
}
