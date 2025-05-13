package dto

type CreateCategoryRequest struct {
	Name string `json:"name" validate:"required,min=1,max=100"`
}

type UpdateCategoryRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required,min=1,max=200"`
}
