package dto

type UpdateProductRequest struct {
	ID          int     `json:"id" validate:"required"`
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
}
