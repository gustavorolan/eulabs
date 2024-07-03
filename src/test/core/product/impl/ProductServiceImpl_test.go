package impl

import (
	"eulabs/src/main/core/product"
	"eulabs/src/main/core/product/impl"
	"eulabs/src/test/factory"
	mock_product "eulabs/src/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func getTestContext(t *testing.T) (*mock_product.MockRepository, impl.ServiceImpl) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	repositoryMock := mock_product.NewMockRepository(controller)
	serviceImpl := impl.ServiceImpl{Repository: repositoryMock}
	return repositoryMock, serviceImpl
}

func TestServiceImpl_GetById_Status_Ok(t *testing.T) {
	repositoryMock, serviceImpl := getTestContext(t)

	entity := factory.GetEntityAsList()
	id := entity[0].ID
	responseExpected := factory.GetResponseProductResponseOk()

	repositoryMock.EXPECT().FindById(strconv.Itoa(id)).Return(entity, nil).Times(1)

	response := serviceImpl.GetById(strconv.Itoa(id))

	assert.Equal(t, responseExpected, response)
}

func TestServiceImpl_GetById_Status_InternalServerError_MoreThanOneReturnRepository(t *testing.T) {
	repositoryMock, serviceImpl := getTestContext(t)

	entity := factory.GetTwoEntityAsList()
	id := entity[0].ID
	responseExpected := factory.GetResponseProductResponseInternalServerErrorMoreThanOneReturn()

	repositoryMock.EXPECT().FindById(strconv.Itoa(id)).Return(entity, nil).Times(1)

	response := serviceImpl.GetById(strconv.Itoa(id))

	assert.Equal(t, responseExpected, response)
}

func TestServiceImpl_GetById_Status_NotFound(t *testing.T) {
	repositoryMock, serviceImpl := getTestContext(t)

	entity := make([]*product.Product, 0)
	id := "1"
	responseExpected := factory.GetResponseProductResponseNotFound()

	repositoryMock.EXPECT().FindById(id).Return(entity, nil).Times(1)

	response := serviceImpl.GetById(id)

	assert.Equal(t, responseExpected, response)
}

func TestServiceImpl_GetById_InternalServerError(t *testing.T) {
	repositoryMock, serviceImpl := getTestContext(t)

	entity := factory.GetEntityAsList()
	id := entity[0].ID
	responseExpected := factory.GetResponseInternalServerError()
	serverError := factory.GetServerError()

	repositoryMock.EXPECT().FindById(strconv.Itoa(id)).Return(nil, serverError).Times(1)

	response := serviceImpl.GetById(strconv.Itoa(id))

	assert.Equal(t, responseExpected, response)
}

func TestServiceImpl_Create_InternalServerError(t *testing.T) {

	repositoryMock, serviceImpl := getTestContext(t)

	request := factory.GetNewProductRequest()
	responseExpected := factory.GetResponseInternalServerError()
	serverError := factory.GetServerError()

	repositoryMock.EXPECT().Create(gomock.Any()).Return(nil, serverError).Times(1)

	response := serviceImpl.Create(request)

	assert.Equal(t, responseExpected, response)
}

func TestServiceImpl_Create_Created(t *testing.T) {

	repositoryMock, serviceImpl := getTestContext(t)

	request := factory.GetNewProductRequest()
	entity := factory.GetProductEntity()
	responseExpected := factory.GetResponseProductResponseCreated()

	repositoryMock.EXPECT().Create(gomock.Any()).Return(entity, nil).Times(1)

	response := serviceImpl.Create(request)

	assert.Equal(t, responseExpected, response)
}

func TestServiceImpl_Update_Created(t *testing.T) {

	repositoryMock, serviceImpl := getTestContext(t)

	request := factory.GetUpdateProductRequest()
	entity := factory.GetProductEntity()
	responseExpected := factory.GetResponseProductResponseOk()

	repositoryMock.EXPECT().Update(gomock.Any()).Return(entity, nil).Times(1)

	response := serviceImpl.Update(request)

	assert.Equal(t, responseExpected, response)
}

func TestServiceImpl_Update_InternalServerError(t *testing.T) {

	repositoryMock, serviceImpl := getTestContext(t)

	request := factory.GetUpdateProductRequest()
	responseExpected := factory.GetResponseInternalServerError()
	serverError := factory.GetServerError()

	repositoryMock.EXPECT().Update(gomock.Any()).Return(nil, serverError).Times(1)

	response := serviceImpl.Update(request)

	assert.Equal(t, responseExpected, response)
}

func TestServiceImpl_Delete_InternalServerError(t *testing.T) {

	repositoryMock, serviceImpl := getTestContext(t)

	entity := factory.GetProductEntity()
	responseExpected := factory.GetResponseInternalServerError()
	serverError := factory.GetServerError()

	repositoryMock.EXPECT().DeleteById(strconv.Itoa(entity.ID)).Return(nil, serverError).Times(1)

	response := serviceImpl.Delete(strconv.Itoa(entity.ID))

	assert.Equal(t, responseExpected, response)
}

func TestServiceImpl_Delete_Ok(t *testing.T) {

	repositoryMock, serviceImpl := getTestContext(t)

	entity := factory.GetProductEntity()
	responseExpected := factory.GetResponseDeleteProductResponseOk()

	repositoryMock.EXPECT().DeleteById(strconv.Itoa(entity.ID)).Return(entity, nil).Times(1)

	response := serviceImpl.Delete(strconv.Itoa(entity.ID))

	assert.Equal(t, responseExpected, response)
}

func TestServiceImpl_GetAll_Ok(t *testing.T) {

	repositoryMock, serviceImpl := getTestContext(t)

	entity := factory.GetPageProductEntity()
	responseExpected := factory.GetResponseProductResponse()
	pageable := factory.GetPageable()

	repositoryMock.EXPECT().FindAll(pageable).Return(entity, nil).Times(1)

	response := serviceImpl.GetAll(pageable)

	assert.Equal(t, responseExpected, response)
}

func TestServiceImpl_GetAll_InternalServerError(t *testing.T) {

	repositoryMock, serviceImpl := getTestContext(t)

	responseExpected := factory.GetResponseInternalServerError()
	pageable := factory.GetPageable()
	entity := factory.GetPageProductEntity()
	err := factory.GetServerError()

	repositoryMock.EXPECT().FindAll(gomock.Any()).Return(entity, err).Times(1)

	response := serviceImpl.GetAll(pageable)

	assert.Equal(t, responseExpected, response)
}
