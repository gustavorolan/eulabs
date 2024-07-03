package dto

type NewProductRequest struct {
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description" validate:"required"`
	Price       float64 `json:"price" validate:"required"`
	Stock       int     `json:"stock" validate:"required"`
}
