package impl

import (
	"eulabs/src/main/core/dto"
	"eulabs/src/main/core/product"
	config "eulabs/src/main/infraestructure/database"
	"gorm.io/gorm"
)

type RepositoryImpl struct{}

func (r RepositoryImpl) Create(productEntity *product.Product) (*product.Product, error) {
	db := config.DB()
	err := db.Create(&productEntity).Error
	return productEntity, err
}
func (r RepositoryImpl) FindById(id string) ([]*product.Product, error) {
	var productEntity []*product.Product
	db := config.DB()
	err := db.Find(&productEntity, id).Error
	return productEntity, err
}
func (r RepositoryImpl) Update(productEntity *product.Product) (*product.Product, error) {
	db := config.DB()
	err := db.Save(&productEntity).Error
	return productEntity, err
}
func (r RepositoryImpl) DeleteById(id string) (*product.Product, error) {
	db := config.DB()
	productEntity := new(product.Product)
	err := db.Delete(&productEntity, id).Error
	return productEntity, err
}

func (r RepositoryImpl) FindAll(pageable *dto.Pageable) (dto.Page, error) {
	db := config.DB()

	var products []*product.Product
	err := db.Scopes(Paginate(pageable)).Find(&products).Error

	var count int64 = 0
	db.Model(&product.Product{}).Count(&count)

	page := dto.NewPage(products, count, pageable)

	return page, err
}

func Paginate(pagination *dto.Pageable) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := (pagination.Page - 1) * pagination.PageSize
		return db.Offset(offset).Limit(pagination.PageSize).Order(pagination.Sort)
	}
}
