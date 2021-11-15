package data

import (
	"github.com/MikelSot/Ed-api/domain"
	"time"
)

type memory struct {
	courses map[uint]domain.Course
}

func NewMemory() *memory {
	return &memory{
		make(map[uint]domain.Course),
	}
}

func (m memory) Create(course *domain.Course) error {
	if course == nil {
		return nullCourseError
	}
	if course.Qualification > 5 || course.Qualification < 0{
		return qualificationError
	}
	course.ID++
	course.CreatedAt = time.Now()
	m.courses[course.ID] = *course
	return nil
}

func (m memory) Update(course *domain.Course) error {
	if course == nil {
		return nullCourseError
	}
	if course.Qualification > 5 || course.Qualification < 0{
		return qualificationError
	}
	if _, ok := m.courses[course.ID]; !ok{
		return errorCourseDoesNotExist
	}
	m.courses[course.ID] = *course
	return nil
}

func (m memory) Delete(id uint) error {
	if _, ok := m.courses[id]; !ok{
		return errorCourseDoesNotExist
	}
	delete(m.courses, id)
	return nil
}

func (m memory) GetByID(id uint) (*domain.Course, error) {
	 course, ok := m.courses[id]
	 if !ok{
		return &course,errorCourseDoesNotExist
	}
	return &course, nil
}

func (m memory) GetAll() (*domain.Courses, error) {
	courses := domain.Courses{}
	for _, course := range m.courses{
		courses = append(courses, &course)
	}
	return &courses, nil
}
