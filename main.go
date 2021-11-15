package main

import (
	"github.com/MikelSot/Ed-api/data"
	"github.com/MikelSot/Ed-api/infrastructure"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"log"
)

func main() {
	data := data.NewMemory()
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	infrastructure.NewRoute(e, data)
	log.Println("Servidor iniciado")
	err := e.Start(":8080")
	if err != nil {
		log.Printf("error en el servidor %v\n", err)
	}
}
