package course

import (
	"time"

	"github.com/HoanNghi16/Devall_backend/internal/user"
	"gorm.io/datatypes"
)

type Course struct {
	ID       		 uint   `gorm:"primaryKey;autoIncrement"`
	Name     		 string `gorm:"not null"`
	Avatar 	 		 string
	AuthorID 		 uint   `gorm:"not null"`
	Author   		 user.Profile `gorm:"foreignKey:AuthorID"`
	ShortDescription string
	CreatedAt 		 time.Time
	UpdatedAt  		 time.Time
	IsPublished 	 bool `gorm:"default:false"`
}

type Lesson struct{
	ID 		 	uint `gorm:"primaryKey;autoIncrement"`
	CourseID 	uint `gorm:"not null"`
	Course 	 	Course `gorm:"foreignKey:CourseID"`
	Position 	uint
	Name 	 	string `gorm:"not null"`
}

type ContentBlock struct {
	ID 			uint `gorm:"primaryKey;autoIncrement"`
	Position 	uint
	LessonID 	uint `gorm:"not null"`
	Lesson 		Lesson `gorm:"foreignKey:LessonID"`
	BlockType 	string  // "text" | "video" | "visualizer" | "codeEditor"
	Data 		datatypes.JSON
}


func (c *Course) ToResponseData (course Course) (*ResponseCourse){
	return &ResponseCourse{
		ID: course.ID,
		Avatar: course.Avatar,
		Name: course.Name,
		ShortDescription: course.ShortDescription,
		CreatedAt: course.CreatedAt,
		UpdatedAt: course.UpdatedAt,
		Author: ResponseAuthor{
			Name: course.Author.Name,
			Avatar: course.Author.Avatar,
		},
	}
}

func (c *Course)ToResponseDataList(coursesList []Course)([]ResponseCourse){
	result := make([]ResponseCourse, len(coursesList))
	for i, course := range coursesList{
		result[i] = *c.ToResponseData(course)
	}
	return result
}