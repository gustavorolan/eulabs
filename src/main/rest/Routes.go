package rest

import (
	"eulabs/src/main/rest/product"
	"github.com/labstack/echo/v4"
)

func Routes(e *echo.Echo) {
	apiRoutes := e.Group("/api/v0")
	product.Resource(apiRoutes)
}
