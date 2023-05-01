package dto

type WithdrawDto struct {
	Id       string `json:"id" validate:"required,uuid4"`
	SecretId string `json:"secret_id" validate:"required,alphanum"`
}
