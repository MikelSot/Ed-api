package infrastructure

import (
	"github.com/MikelSot/Ed-api/domain"
	"github.com/MikelSot/Ed-api/interfaces"
	"github.com/labstack/echo"
	"net/http"
	"strconv"
)

type course struct {
	data interfaces.Data
}

func newCourse(data interfaces.Data) *course{
	return &course{data}
}

func (c course) create(e echo.Context) error {
	data := domain.Course{}
	err := e.Bind(&data)
	if err != nil{
		res := newResponse(Error, "Error en la estructura", nil)
		return e.JSON(http.StatusBadRequest,res)
	}

	err = c.data.Create(&data)
	if err != nil{
		res := newResponse(Error, "Hubo un problema", nil)
		return e.JSON(http.StatusInternalServerError,res)
	}

	res := newResponse(Error, "ok", nil)
	return  e.JSON(http.StatusCreated, res)
}

func (c course) update (e echo.Context) error{
	data := domain.Course{}
	err := e.Bind(&data)
	if err != nil{
		res := newResponse(Error, "Error en la estructura", nil)
		return e.JSON(http.StatusBadRequest,res)
	}

	err = c.data.Update(&data)
	if err != nil{
		res := newResponse(Error, "Hubo un problema al actualizar este registro", nil)
		return e.JSON(http.StatusInternalServerError,res)
	}

	res := newResponse(Message, "se actualizo el curso", nil)
	return e.JSON(http.StatusOK, res)
}

func (c course) delete(e echo.Context) error{
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil{
		res := newResponse(Error, "Id no valido", nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	err = c.data.Delete(uint(id))
	if err != nil {
		res := newResponse(Error, err.Error(), nil)
		return e.JSON(http.StatusBadRequest, res)
	}
	res := newResponse(Message, "ok", nil)
	return e.JSON(http.StatusOK, res)
}


func (c course) getById(e echo.Context) error{
	id, err := strconv.Atoi(e.Param("id"))
	if err != nil{
		res := newResponse(Error, "Id no valido", nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	course, err := c.data.GetByID(uint(id))
	if err != nil {
		res := newResponse(Error, err.Error(), nil)
		return e.JSON(http.StatusBadRequest, res)
	}

	res := newResponse(Message, "ok", course)
	return e.JSON(http.StatusOK, res)
}

func (c course) getAll(e echo.Context) error{
	courses, err := c.data.GetAll()
	if err != nil {
		res := newResponse(Error, err.Error(), nil)
		return e.JSON(http.StatusBadRequest, res)
	}
	res := newResponse(Message, "ok", courses)
	return e.JSON(http.StatusOK, res)
}