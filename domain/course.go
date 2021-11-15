package domain

import "time"

type Course struct {
	Name          string    `json:"name"`
	Description   string    `json:"description"`
	Duration      string    `json:"duration"`
	Category      string    `json:"category"`
	Teacher       string    `json:"teacher"`
	CoverImage    string    `json:"cover_image"`
	Qualification float32   `json:"qualification"`
	Syllabus      []string  `json:"syllabus"`
	CreatedAt     time.Time `json:"created_at"`
}

type Courses []*Course
