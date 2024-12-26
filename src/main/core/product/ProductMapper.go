package product

import (
	"eulabs/src/main/core/dto"
)

func (product Product) ToResponse() dto.ProductResponse {
	return dto.ProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}
}

func ToResponse(productsMap map[*Product]error) *dto.ProductsResponse {
	var products = make([]*dto.ProductResponse, 0)
	var response = new(dto.ProductsResponse)

	for key, value := range productsMap {
		product := &dto.ProductResponse{
			ID:           key.ID,
			Name:         key.Name,
			Description:  key.Description,
			Price:        key.Price,
			Stock:        key.Stock,
			ErrorMessage: "",
		}

		if value != nil {
			product.ErrorMessage = value.Error()
		}

		products = append(products, product)
	}

	response.Products = products

	return response
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
