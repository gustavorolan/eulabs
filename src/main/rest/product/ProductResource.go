package product

import (
	"eulabs/src/main/core/product/impl"
	repository "eulabs/src/main/infraestructure/database/product/impl"
	"github.com/labstack/echo/v4"
)

func Resource(e *echo.Group) {

	bookRoute := e.Group("/product")

	handler := Handler{
		Service: impl.ServiceImpl{
			Repository: repository.RepositoryImpl{},
		},
	}

	bookRoute.POST("", handler.Create)
	bookRoute.GET("", handler.GetAll)
	bookRoute.GET("/:id", handler.GetById)
	bookRoute.PUT("/:id", handler.Update)
	bookRoute.DELETE("/:id", handler.Delete)
}
