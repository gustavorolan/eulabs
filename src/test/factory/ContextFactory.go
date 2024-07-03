package factory

import (
	"encoding/json"
	valdiatorconfig "eulabs/src/main/rest/config"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
)

func GetContext(request map[string]interface{}, rec *httptest.ResponseRecorder) (echo.Context, *httptest.ResponseRecorder) {
	requestBody, _ := json.Marshal(request)
	e := echo.New()
	e.Validator = &valdiatorconfig.CustomValidator{Validator: validator.New()}
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(string(requestBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec = httptest.NewRecorder()
	return e.NewContext(req, rec), rec
}
