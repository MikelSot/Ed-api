package interfaces

import "github.com/MikelSot/Ed-api/domain"

type Data interface {
	Create(course *domain.Course) error
	Update(course *domain.Course) error
	Delete(id uint) error
	GetByID(id uint) (*domain.Course, error)
	GetAll() (*domain.Courses, error)
}
