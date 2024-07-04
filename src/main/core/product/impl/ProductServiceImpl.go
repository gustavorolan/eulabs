package impl

import (
	"eulabs/src/main/core/dto"
	"eulabs/src/main/core/product"
	"log"
	"os"
	"strconv"
)

type ServiceImpl struct {
	Repository product.Repository
}

var loggerInfo = log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile)

var loggerError = log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile)

func (s ServiceImpl) Create(request *dto.NewProductRequest) dto.Response {

	productEntity := product.NewProductRequestToEntity(request)

	productEntitySaved, err := s.Repository.Create(&productEntity)

	if err != nil {
		return dto.NewErrorResponseInternalServerError(err.Error())
	}

	loggerInfo.Println(`Product created successfully id: ` + strconv.Itoa(productEntitySaved.ID))

	return dto.NewSuccessResponseCreated(productEntitySaved.ToResponse())
}

func (s ServiceImpl) GetById(id string) dto.Response {
	products, err := s.Repository.FindById(id)

	loggerInfo.Println(`Consulting product by id: ` + id)

	if err != nil {
		loggerError.Println(err.Error())
		return dto.NewErrorResponseInternalServerError(err.Error())
	} else if len(products) > 1 {
		message := "More than one entity has returned for this id: " + id
		loggerError.Println(message)
		return dto.NewErrorResponseInternalServerError(message)
	} else if len(products) == 0 {
		message := "Product was not found id: " + id
		loggerError.Println(message)
		return dto.NewErrorResponseNotFound(message)
	}

	return dto.NewSuccessResponseOk(products[0].ToResponse())
}

func (s ServiceImpl) Update(request *dto.UpdateProductRequest) dto.Response {

	productEntity := product.UpdateProductRequestToEntity(request)

	productEntityUpdated, err := s.Repository.Update(&productEntity)

	if err != nil {
		loggerError.Println(err.Error())
		return dto.NewErrorResponseInternalServerError(err.Error())
	}

	loggerInfo.Println(`Product updated successfully id: ` + strconv.Itoa(productEntityUpdated.ID))

	return dto.NewSuccessResponseOk(productEntityUpdated.ToResponse())
}

func (s ServiceImpl) Delete(id string) dto.Response {

	_, err := s.Repository.DeleteById(id)

	if err != nil {
		loggerError.Println(err.Error())
		return dto.NewErrorResponseInternalServerError(err.Error())
	}

	response := map[string]interface{}{
		"message": "A Product Product has been deleted id: " + id,
	}

	loggerInfo.Println(response)

	return dto.NewSuccessResponseOk(response)
}

func (s ServiceImpl) GetAll(pageable *dto.Pageable) dto.Response {

	loggerInfo.Println("Consulting all products")

	page, err := s.Repository.FindAll(pageable)

	if err != nil {
		loggerError.Println(err.Error())
		return dto.NewErrorResponseInternalServerError(err.Error())
	}

	pageResponse := product.MapToProductResponseList(page)

	return dto.NewSuccessResponseOk(pageResponse)
}
