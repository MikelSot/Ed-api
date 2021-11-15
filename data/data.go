package data

import "errors"

var (
	nullCourseError = errors.New("Error curso nulo")
	errorCourseDoesNotExist = errors.New("El curso no existe")
	qualificationError = errors.New("Error de calificacion")
)

