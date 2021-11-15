package interfaces

import "github.com/MikelSot/Ed-api/domain"

type Data interface {
	Create(course *domain.Course) error
	Update(id int, course *domain.Course) error
	Delete(id int) error
	GetByID(id int) (*domain.Course, error)
	GetAll() (*domain.Courses, error)
}
