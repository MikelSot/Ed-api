package infrastructure

import (
	"github.com/MikelSot/Ed-api/interfaces"
	"github.com/labstack/echo"
)

type route struct {
	echo *echo.Echo
	data interfaces.Data
}

func NewRoute(e *echo.Echo, d interfaces.Data) *route {
	return &route{e, d}
}

func (r route) Routes()  {
	handler := newCourse(r.data)
	course := r.echo.Group("/v1/course")

	course.POST("/create", handler.create)
	course.PUT("/create", handler.update)
	course.GET("/get-by-id/:id", handler.getById)
	course.GET("/get-all", handler.getAll)
}

