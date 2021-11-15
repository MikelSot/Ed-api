package middleware

import (
	"github.com/labstack/echo"
	"log"
)

func Log(fun echo.HandlerFunc) echo.HandlerFunc{
	return func(c echo.Context) error {
		log.Printf("peticion %q, metodo: %q", c.Path())
		return fun(c)
	}
}