package factory

import (
	"errors"
	"eulabs/src/main/core/dto"
	"eulabs/src/main/core/product"
)

func GetProductEntity() *product.Product {
	return &product.Product{
		ID:          1,
		Name:        "name",
		Description: "desc",
		Price:       10.00,
		Stock:       10,
	}
}

func GetNewProductRequest() *dto.NewProductRequest {
	return &dto.NewProductRequest{
		Name:        "name",
		Description: "desc",
		Price:       10.00,
		Stock:       10,
	}
}

func GetUpdateProductRequest() *dto.UpdateProductRequest {
	return &dto.UpdateProductRequest{
		ID:          1,
		Name:        "name",
		Description: "desc",
		Price:       10.00,
		Stock:       10,
	}
}

func GetEntityAsList() []*product.Product {
	products := make([]*product.Product, 0)
	products = append(products, GetProductEntity())
	return products
}

func GetResponseAsList() []dto.ProductResponse {
	products := make([]dto.ProductResponse, 0)
	products = append(products, GetProductResponse())
	return products
}

func GetTwoEntityAsList() []*product.Product {
	products := make([]*product.Product, 0)
	products = append(products, GetProductEntity())
	products = append(products, GetProductEntity())
	return products
}

func GetProductResponse() dto.ProductResponse {
	return dto.ProductResponse{
		ID:          1,
		Name:        "name",
		Description: "desc",
		Price:       10.00,
		Stock:       10,
	}
}

func GetResponseProductResponseOk() dto.Response {
	return dto.NewSuccessResponseOk(GetProductResponse())
}

func GetResponseDeleteProductResponseOk() dto.Response {
	return dto.NewSuccessResponseOk(map[string]interface{}{
		"message": "A Product Product has been deleted id: 1",
	})
}

func GetResponseProductResponseCreated() dto.Response {
	return dto.NewSuccessResponseCreated(GetProductResponse())
}

func GetResponseProductResponseNotFound() dto.Response {
	return dto.NewErrorResponseNotFound("Product was not found id: 1")
}

func GetResponseProductResponseInternalServerErrorMoreThanOneReturn() dto.Response {
	return dto.NewErrorResponseInternalServerError(
		"More than one entity has returned for this id: 1",
	)
}

func GetResponseProductResponse() dto.Response {
	return dto.NewSuccessResponseOk(GetPageProductResponse())
}

func GetPageProductResponse() dto.Page {
	return dto.Page{
		Page:       1,
		PageSize:   1,
		Sort:       "ID desc",
		Data:       GetResponseAsList(),
		TotalPages: 1,
		Total:      1,
	}
}

func GetPageProductEntity() dto.Page {
	return dto.Page{
		Page:       1,
		PageSize:   1,
		Sort:       "ID desc",
		Data:       GetEntityAsList(),
		TotalPages: 1,
		Total:      1,
	}
}

func GetPageable() *dto.Pageable {
	return &dto.Pageable{
		Page:     1,
		PageSize: 1,
		Sort:     "ID desc",
	}
}

func GetResponseInternalServerError() dto.Response {
	return dto.NewErrorResponseInternalServerError(GetServerError().Error())
}

func GetServerError() error {
	return errors.New("generic error")
}
