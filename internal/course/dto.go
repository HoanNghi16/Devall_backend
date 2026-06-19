package course

import (
	"time"

	"gorm.io/datatypes"
)

// Dùng để lấy đầu vào filter
type CourseFilter struct{
	TopicIDs 	[]uint `form:"topics"`
	Level 		string `form:"level" binding:"omitempty,oneof=easy medium hard advanced"`
	Cursor 		uint `form:"cursor"`
}

type ResponseAuthor struct{
	Name string `json:"name"`
	Avatar string `json:"avatar"`
}

type ResponseCourse struct { //Để json.Marshal() trả về đúng tên fields
	ID               uint  `json:"id"`
	Name             string`json:"name"`
	Avatar           string`json:"avatar"`
	Author           ResponseAuthor `json:"author"`
	ShortDescription string `json:"short_description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
}


//Dùng cho request POST,PUT,PATCH

type RequestContentBlock struct {
	Position 	uint `json:"position"`
	BlockType 	string  `json:"block_type" binding:"oneof=text video visualizer codeEditor codePreview"` // "text" | "video" | "visualizer" | "codeEditor" | "codePreview"
	Data 		datatypes.JSON `json:"data"`
}

type RequestLesson struct{
	Position 	uint `json:"position"`
	Name 	 	string `json:"name" binding:"required"`
	ContentBlocks []RequestContentBlock `json:"content_blocks" binding:"required,min=1,dive"`
}

type RequestCourse struct {
	Name     		 string `json:"name" binding:"required"`
	Avatar 	 		 string	`json:"avatar"`
	ShortDescription string `json:"short_description" binding:"required"`
	IsPublished 	 bool `json:"is_published"`
	Lessons []RequestLesson `json:"lessons" binding:"required,min=1,dive"`
	Topics []uint `json:"topics"`
	Level string `json:"level" binding:"oneof=easy medium hard advanced"` // easy | medium | hard | advanced
}
