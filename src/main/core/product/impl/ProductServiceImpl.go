package impl

import (
	"eulabs/src/main/core/dto"
	"eulabs/src/main/core/product"
)

type ServiceImpl struct {
	Repository product.Repository
}

func (s ServiceImpl) Create(request *dto.NewProductRequest) dto.Response {

	productEntity := product.NewProductRequestToEntity(request)

	productEntitySaved, err := s.Repository.Create(&productEntity)

	if err != nil {
		return dto.NewErrorResponseInternalServerError(err.Error())
	}

	return dto.NewSuccessResponseCreated(productEntitySaved.ToResponse())
}

func (s ServiceImpl) GetById(id string) dto.Response {

	products, err := s.Repository.FindById(id)

	if err != nil {
		return dto.NewErrorResponseInternalServerError(err.Error())
	} else if len(products) > 1 {
		return dto.NewErrorResponseInternalServerError("More than one entity has returned for this id: " + id)
	} else if len(products) == 0 {
		return dto.NewErrorResponseNotFound("Product was not found id: " + id)
	}

	return dto.NewSuccessResponseOk(products[0].ToResponse())
}

func (s ServiceImpl) Update(request *dto.UpdateProductRequest) dto.Response {

	productEntity := product.UpdateProductRequestToEntity(request)

	productEntityUpdated, err := s.Repository.Update(&productEntity)

	if err != nil {
		return dto.NewErrorResponseInternalServerError(err.Error())
	}

	return dto.NewSuccessResponseOk(productEntityUpdated.ToResponse())
}

func (s ServiceImpl) Delete(id string) dto.Response {

	_, err := s.Repository.DeleteById(id)

	if err != nil {
		return dto.NewErrorResponseInternalServerError(err.Error())
	}

	response := map[string]interface{}{
		"message": "A Product Product has been deleted id: " + id,
	}

	return dto.NewSuccessResponseOk(response)
}

func (s ServiceImpl) GetAll(pageable *dto.Pageable) dto.Response {

	page, err := s.Repository.FindAll(pageable)

	if err != nil {
		return dto.NewErrorResponseInternalServerError(err.Error())
	}

	pageResponse := product.MapToProductResponseList(page)

	return dto.NewSuccessResponseOk(pageResponse)
}
