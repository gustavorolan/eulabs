package product

import (
	"encoding/json"
	"eulabs/src/main/rest/product"
	"eulabs/src/test/factory"
	mock_product "eulabs/src/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func getTestContext(t *testing.T) (*mock_product.MockService, product.Handler) {
	controller := gomock.NewController(t)
	defer controller.Finish()
	serviceMock := mock_product.NewMockService(controller)
	handler := product.Handler{Service: serviceMock}
	return serviceMock, handler
}

func TestHandler_Create_FieldError(t *testing.T) {
	rec := httptest.NewRecorder()
	body := factory.GetCreateNewProductJsonRequestWithoutSomeParams()
	_, handler := getTestContext(t)
	responseExpected := factory.GetProductResponseWithoutSomeParams()
	echoContext, rec := factory.GetContext(body, rec)

	err := handler.Create(echoContext)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusBadRequest, echoContext.Response().Status)
		assert.Equal(t, responseExpected, unmarshal(rec))
	}
}

func TestHandler_Create_InvalidFieldError(t *testing.T) {
	rec := httptest.NewRecorder()
	body := factory.GetCreateNewProductJsonRequestInvalidParam()
	_, handler := getTestContext(t)
	responseExpected := factory.GetProductResponseWithInvalidParam()
	echoContext, rec := factory.GetContext(body, rec)

	err := handler.Create(echoContext)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusInternalServerError, echoContext.Response().Status)
		assert.Equal(t, responseExpected, unmarshal(rec))
	}
}

func unmarshal(echoResponse *httptest.ResponseRecorder) map[string]interface{} {
	var responseBody map[string]interface{}
	_ = json.Unmarshal(echoResponse.Body.Bytes(), &responseBody)
	return responseBody
}
