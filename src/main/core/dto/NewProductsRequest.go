package dto

type NewProductsRequest struct {
	Products []NewProductRequest `json:"products" validate:"required"`
}
