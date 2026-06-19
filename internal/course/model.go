package course

import (
	"time"

	"github.com/HoanNghi16/Devall_backend/internal/user"
	"gorm.io/datatypes"
)

type Course struct {
	ID       		 uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name     		 string `gorm:"not null;unique" json:"name"`
	Avatar 	 		 string	`json:"avatar"`
	AuthorID 		 uint   `gorm:"not null" json:"author_id"`
	Author   		 user.Profile `gorm:"foreignKey:AuthorID" json:"author"`
	ShortDescription string `json:"short_description"`
	CreatedAt 		 time.Time 	`json:"created_at"`
	UpdatedAt  		 time.Time	`json:"updated_at"`
	IsPublished 	 bool `gorm:"default:false" json:"is_published"`
	Lessons []Lesson `json:"lessons"`
	Topics []TopicCourse `json:"topics"`
	Level string `gorm:"default:easy" json:"level"`
}

type Lesson struct{
	ID 		 	uint `gorm:"primaryKey;autoIncrement" json:"id"`
	CourseID 	uint `gorm:"not null" json:"course_id"`
	Course 	 	Course `gorm:"foreignKey:CourseID"`
	Position 	uint `json:"position"`
	Name 	 	string `gorm:"not null" json:"name"`
	ContentBlocks []ContentBlock `json:"content_blocks"`
}

type ContentBlock struct {
	ID 			uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Position 	uint `json:"position"`
	LessonID 	uint `gorm:"not null" json:"lesson_id"`
	Lesson 		Lesson `gorm:"foreignKey:LessonID"`
	BlockType 	string  `json:"block_type"` // "text" | "video" | "visualizer" | "codeEditor" | "codePreview"
	Data 		datatypes.JSON `json:"data"`
}


type Topic struct{
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null"`
}

type TopicCourse struct{
	CourseID uint `gorm:"primaryKey"`
	TopicID uint `gorm:"primaryKey"`
	Topic Topic `gorm:"foreignKey:TopicID"`
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