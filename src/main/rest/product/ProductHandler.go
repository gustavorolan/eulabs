package product

import (
	"eulabs/src/main/core/dto"
	"eulabs/src/main/core/product"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	Service product.Service
}

func (handler Handler) Create(c echo.Context) error {
	productRequest := new(dto.NewProductRequest)

	if err := c.Bind(productRequest); err != nil {
		response := dto.NewErrorResponseInternalServerError(err.Error())
		return toJson(response, c)
	}

	if err := c.Validate(productRequest); err != nil {
		return toJson(dto.NewErrorResponseBadRequest(err.Error()), c)
	}

	response := handler.Service.Create(productRequest)

	return c.JSON(response.HTTPStatus, response)
}

func (handler Handler) GetById(c echo.Context) error {
	id := c.Param("id")

	response := handler.Service.GetById(id)

	return toJson(response, c)
}

func (handler Handler) Update(c echo.Context) error {

	request := new(dto.UpdateProductRequest)

	if err := c.Bind(request); err != nil {
		response := dto.NewErrorResponseInternalServerError(err.Error())
		return toJson(response, c)
	}

	if err := c.Validate(request); err != nil {
		return toJson(dto.NewErrorResponseBadRequest(err.Error()), c)
	}

	response := handler.Service.Update(request)

	return toJson(response, c)
}

func (handler Handler) Delete(c echo.Context) error {

	id := c.Param("id")

	response := handler.Service.Delete(id)

	return toJson(response, c)
}

func (handler Handler) GetAll(c echo.Context) error {

	pageable := new(dto.Pageable)

	if err := c.Bind(pageable); err != nil {
		response := dto.NewErrorResponseInternalServerError(err.Error())
		return toJson(response, c)
	}

	if err := c.Validate(pageable); err != nil {
		return toJson(dto.NewErrorResponseBadRequest(err.Error()), c)
	}

	response := handler.Service.GetAll(pageable)

	return toJson(response, c)
}

func toJson(r dto.Response, c echo.Context) error {
	return c.JSON(r.HTTPStatus, r)
}
