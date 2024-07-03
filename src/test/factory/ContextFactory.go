package factory

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
	"net/http/httptest"
	"strings"
)

func GetContext(request map[string]interface{}) echo.Context {
	requestBody, _ := json.Marshal(request)
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/products", strings.NewReader(string(requestBody)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	return e.NewContext(req, rec)
}
