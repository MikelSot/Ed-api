package data

import (
	"github.com/MikelSot/Ed-api/domain"
	"log"
	"time"
)

var (
	id = 0
)

type memory struct {
	courses map[int]domain.Course
}

func NewMemory() *memory {
	return &memory{
		make(map[int]domain.Course),
	}
}

func (m memory) Create(course *domain.Course) error {
	if course == nil {
		return nullCourseError
	}
	if course.Qualification > 5 || course.Qualification < 0 {
		return qualificationError
	}
	id++
	log.Println("ID -->", id)
	course.CreatedAt = time.Now()
	m.courses[id] = *course
	return nil
}

func (m memory) Update(id int, course *domain.Course) error {
	if course == nil {
		return nullCourseError
	}
	if course.Qualification > 5 || course.Qualification < 0 {
		return qualificationError
	}
	if _, ok := m.courses[id]; !ok {
		return errorCourseDoesNotExist
	}
	m.courses[id] = *course
	return nil
}

func (m memory) Delete(id int) error {
	if _, ok := m.courses[id]; !ok {
		return errorCourseDoesNotExist
	}
	delete(m.courses, id)
	return nil
}

func (m memory) GetByID(id int) (*domain.Course, error) {
	course, ok := m.courses[id]
	if !ok {
		return &course, errorCourseDoesNotExist
	}
	return &course, nil
}

func (m memory) GetAll() (*domain.Courses, error) {
	courses := domain.Courses{}
	for _, course := range m.courses {
		courses = append(courses, &course)
	}
	return &courses, nil
}
