package main

import (
	config "eulabs/src/main/infraestructure/database"
	"eulabs/src/main/rest"
	valdiatorconfig "eulabs/src/main/rest/config"
	"github.com/go-playground/validator/v10"
	_ "github.com/gustavorolan/eulabs/docs"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title eulabs
// @version 0.0.1
// @description
// @BasePath /api/v0
func main() {

	config.DatabaseInit()

	e := echo.New()

	e.Validator = &valdiatorconfig.CustomValidator{Validator: validator.New()}

	rest.Routes(e)

	e.GET("/swagger/*", echoSwagger.WrapHandler)

	e.Logger.Fatal(e.Start(":8080"))
}
