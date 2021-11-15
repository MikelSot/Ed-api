package main

import (
	"github.com/MikelSot/Ed-api/data"
	"github.com/MikelSot/Ed-api/infrastructure"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
)

func main() {
	data := data.NewMemory()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	route := infrastructure.NewRoute(e, data)
	route.Routes()
	log.Println("Servidor iniciado")
	err := e.Start(":3000")
	if err != nil {
		log.Printf("error en el servidor %v\n", err)
	}
}
