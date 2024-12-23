package dto

type UserRegisterDto struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8"`
	FullName  string `json:"fullName" validate:"required"`
	BirthDate string `json:"birthDate" validate:"required,date"`
	Gender    string `json:"gender" validate:"required,oneof=Male Female"`
}
