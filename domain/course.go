package domain

import "time"

type Course struct {
	ID            uint      `json:"id"`
	Name          string    `json:"name"`
	Description   string    `json:"description,omitempty"`
	Duration      string    `json:"duration,omitempty"`
	ReleaseDate   time.Time `json:"release_date,omitempty"`
	Category      string    `json:"category,omitempty"`
	Teacher       string    `json:"teacher,omitempty"`
	CoverImage    string    `json:"cover_image,omitempty"`
	Qualification float32   `json:"qualification,omitempty"`
	Syllabus      []string    `json:"syllabus,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
}

type Courses []*Course
