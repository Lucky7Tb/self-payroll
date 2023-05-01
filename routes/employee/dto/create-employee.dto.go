package dto

type CreateEmployeeDto struct {
	EmployeeId string `json:"employee_id" validate:"required,numeric"`
	SecretId   string `json:"secret_id" validate:"required,alphanum"`
	PositionId string `json:"position_id" validate:"required,uuid4"`
	Name       string `json:"name" validate:"required,printascii"`
	Email      string `json:"email" validate:"required,email"`
	Phone      string `json:"phone" validate:"required,numeric,min=10,max=14"`
	Address    string `json:"address" validate:"required,printascii"`
}
