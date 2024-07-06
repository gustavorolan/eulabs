package product

//go:generate mockgen -source=ProductService.go -destination=../../../test/mocks/ProductServiceMock.go

import "eulabs/src/main/core/dto"

type Service interface {
	Create(request *dto.NewProductRequest) dto.Response
	GetById(id string) dto.Response
	Update(b *dto.UpdateProductRequest) dto.Response
	Delete(id string) dto.Response
	GetAll(*dto.Pageable) dto.Response
	CreateMany() dto.Response
	CreateManyWithChanel() dto.Response
}
