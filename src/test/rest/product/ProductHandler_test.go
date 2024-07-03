package product

import (
	"eulabs/src/main/rest/product"
	"eulabs/src/test/factory"
	mock_product "eulabs/src/test/mocks"
	"github.com/golang/mock/gomock"
	"testing"
)

func TestHandler_Create(t *testing.T) {
	context := factory.GetContext(factory.GetCreateNewProductJsonRequest())

	controller := gomock.NewController(t)
	defer controller.Finish()
	serviceMock := mock_product.NewMockService(controller)

	handler := product.Handler{Service: serviceMock}

}
