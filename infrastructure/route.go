package infrastructure

import (
	"github.com/MikelSot/Ed-api/interfaces"
	"github.com/labstack/echo/v4"
)

type route struct {
	echo *echo.Echo
	data interfaces.Data
}

func NewRoute(e *echo.Echo, d interfaces.Data) *route {
	return &route{e, d}
}

func (r route) routes()  {
	handler := newCourse(r.data)
	course := r.echo.Group("/api/v1/course")

	course.POST("", handler.create)
	course.PUT("/:id", handler.update)
	course.GET("/:id", handler.getById)
	course.GET("", handler.getAll)
	course.DELETE("", handler.delete)
}

func (r route) Routes()  {
	r.routes()
}

