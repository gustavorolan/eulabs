package product

//go:generate mockgen -source=ProductRepository.go -destination=../../../test/mocks/ProductRepositoryMock.go

import "eulabs/src/main/core/dto"

type Repository interface {
	Create(b *Product) (*Product, error)
	FindById(id string) ([]*Product, error)
	Update(b *Product) (*Product, error)
	DeleteById(id string) (*Product, error)
	FindAll(pageable *dto.Pageable) (dto.Page, error)
}
