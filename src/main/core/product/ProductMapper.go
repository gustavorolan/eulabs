package product

import "eulabs/src/main/core/dto"

func (product Product) ToResponse() dto.ProductResponse {
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

func NewProductRequestToEntity(request *dto.NewProductRequest) Product {
	return Product{
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
	}
}

func UpdateProductRequestToEntity(request *dto.UpdateProductRequest) Product {
	return Product{
		ID:          request.ID,
		Name:        request.Name,
		Description: request.Description,
		Price:       request.Price,
		Stock:       request.Stock,
	}
}

func MapToProductResponseList(page dto.Page) dto.Page {

	newProductResponseList := make([]dto.ProductResponse, 0)

	if dataSlice, ok := page.Data.([]*Product); ok {
		for _, p := range dataSlice {
			newProductResponseList = append(newProductResponseList, p.ToResponse())
		}
	}

	page.Data = newProductResponseList
	return page
}
