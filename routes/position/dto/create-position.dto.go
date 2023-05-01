package dto

type CreateUserDto struct {
	Name   string `json:"name" validate:"required"`
	Salary uint32 `json:"salary" validate:"required,numeric"`
}
