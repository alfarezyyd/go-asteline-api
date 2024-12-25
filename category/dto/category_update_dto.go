package dto

type CategoryUpdateDto struct {
	Name        string `json:"name" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required"`
}
