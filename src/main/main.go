package main

import (
	config "eulabs/src/main/infraestructure/database"
	"eulabs/src/main/rest"
	valdiatorconfig "eulabs/src/main/rest/config"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

func main() {

	config.DatabaseInit()

	e := echo.New()

	e.Validator = &valdiatorconfig.CustomValidator{Validator: validator.New()}

	rest.Routes(e)

	e.Logger.Fatal(e.Start(":3000"))
}
