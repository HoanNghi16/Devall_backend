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
	Lessons []Lesson `gorm:"foreignKey:CourseId" json:"lessons"`
	Topics []TopicCourse `gorm:"foreignKey:CourseID" json:"topics"`
	Password string
	Level string `gorm:"default:easy" json:"level"`
}

type Lesson struct{
	ID 		 	uint `gorm:"primaryKey;autoIncrement" json:"id"`
	CourseID 	uint `gorm:"not null" json:"course_id"`
	Position 	uint `json:"position"`
	Name 	 	string `gorm:"not null" json:"name"`
	ContentBlocks []ContentBlock `gorm:"foreignKey:LessonID" json:"content_blocks"`
}

type ContentBlock struct {
	ID 			uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Position 	uint `json:"position"`
	LessonID 	uint `gorm:"not null" json:"lesson_id"`
	BlockType 	string  `json:"block_type"` // "text" | "video" | "visualizer" | "codeEditor" | "codePreview"
	Data 		datatypes.JSON `json:"data"`
}


type Topic struct{
	ID uint `gorm:"primaryKey;autoIncrement" json:"id"`
	Name string `gorm:"not null" json:"name"`
}

type TopicCourse struct{
	CourseID uint `gorm:"primaryKey"`
	TopicID uint `gorm:"primaryKey"`
	Topic Topic `gorm:"foreignKey:TopicID"`
}

type CourseUser struct{
	CourseID uint `gorm:"primaryKey"`
	UserID uint	`gorm:"primaryKey"`
	Course Course `gorm:"foreignKey:CourseID" json:"course"`
	User user.User `gorm:"foreignKey:UserID" json:"user"`
	Progress float32 `gorm:"not null; check: progress >= 0 and progress <= 1" json:"progress"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`
	UpdatedAt time.Time `gorm:"not null" json:"updated_at"`
	DeletedAt time.Time `gorm:"null" json:"deleted_at"`
	IsActive bool `gorm:"default:true" json:"is_active"` //Dùng để hiển thị trong trang lịch sử hoặc ko
	IsMarked bool `gorm:"default:false" json:"is_marked"`
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