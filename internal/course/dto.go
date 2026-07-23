package course

import (
	"time"

	"gorm.io/datatypes"
)

// Dùng để lấy đầu vào filter
type CourseFilter struct{
	TopicIDs 	[]uint `form:"topics"`
	Level 		string `form:"level" binding:"omitempty,oneof=easy medium hard advanced all"`
	Cursor 		uint `form:"cursor"`
}

type ResponseAuthor struct{
	Name string `json:"name"`
	Avatar string `json:"avatar"`
}


type RequestCourseUser struct{
	Progress  		*float32 `json:"progress" binding:"omitempty"`
	CreatedAt 		*time.Time `json:"created_at" binding:"omitempty"`
	UpdatedAt 		*time.Time `json:"updated_at" binding:"omitempty"`
	DeletedAt 		*time.Time `json:"deleted_at" binding:"omitempty"`
	IsActive  		*bool `json:"is_active" binding:"omitempty"` //Dùng để hiển thị trong trang lịch sử hoặc ko
	IsMarked  		*bool `json:"is_marked" binding:"omitempty"`
	LastAccessAt	*time.Time `json:"last_access_at" binding:"omitempty"`
}

type ResponseCourse struct { //Để json.Marshal() trả về đúng tên fields
	ID               uint  `json:"id"`
	Name             string`json:"name"`
	Avatar           string`json:"avatar"`
	Author           ResponseAuthor `json:"author"`
	ShortDescription string `json:"short_description"`
	CreatedAt        time.Time `json:"created_at"`
	UpdatedAt        time.Time `json:"updated_at"`
	CourseUser  	*CourseUser `json:"course_user"`
}


//Dùng cho request POST,PUT,PATCH

type RequestContentBlock struct {
	Position 	uint `json:"position"`
	BlockType 	string  `json:"block_type" binding:"oneof=text media visualizer codeEditor codePreview"` // "text" | "video" | "visualizer" | "codeEditor" | "codePreview"
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
	Password string
}

func (request *RequestCourse) ParseContentBlocks(requestBlocks []RequestContentBlock) []ContentBlock{
	contentBlocks := make([]ContentBlock, len(requestBlocks))
	for index,block := range requestBlocks{
		contentBlocks[index] = ContentBlock{
			Position: block.Position,
			BlockType: block.BlockType,
			Data: block.Data,
		}
	}
	return contentBlocks
}


func (request *RequestCourse) ParseLessons(requestLessons []RequestLesson) []Lesson{
	lessons := make([]Lesson, len(requestLessons))

	for index, lesson := range requestLessons{
		lessons[index] = Lesson{
			Position: lesson.Position,
			Name: lesson.Name,
			ContentBlocks: request.ParseContentBlocks(lesson.ContentBlocks),
		}
	}
	return lessons
}

func (request *RequestCourse) ParseTopics(topicIDs []uint) []TopicCourse{
	topics := make([]TopicCourse, len(topicIDs))
	for index, id := range topicIDs{
		topics[index] = TopicCourse{
			TopicID: id,
		}
	}
	return topics
}

func (request *RequestCourse) ParseCourse() Course{
	return Course{
		Name: request.Name,
		ShortDescription: request.ShortDescription,
		IsPublished: request.IsPublished,
		Level: request.Level,
		Lessons: request.ParseLessons(request.Lessons),
		Avatar: request.Avatar,
		Topics: request.ParseTopics(request.Topics),
	}
}

func (request *RequestCourseUser)ParseCourseUser() (CourseUser, []string){
	now := time.Now()
	var columns []string

	createdAt := now
	if request.CreatedAt != nil{
		createdAt = *request.CreatedAt
		columns = append(columns, "created_at")
	}

	deletedAt := request.DeletedAt

	isMarked := false
	if request.IsMarked != nil{
		columns = append(columns, "is_marked")
		isMarked = *request.IsMarked
	}

	progress := float32(0) 

	if request.Progress != nil{
		progress = *request.Progress
		columns = append(columns, "progress")
	}

	updatedAt := now

	isActive := true
	if request.IsActive != nil{
		isActive = *request.IsActive
		columns = append(columns, "is_active")
	}

	lastAcccessAt := now
	if request.LastAccessAt != nil {
		lastAcccessAt = *request.LastAccessAt
		columns = append(columns, "progress")
	}

	return CourseUser{
		DeletedAt: deletedAt,
		CreatedAt: createdAt,
		IsMarked: isMarked,
		Progress: progress,
		UpdatedAt: updatedAt,
		IsActive: isActive,
		LastAccessAt: lastAcccessAt,
	},columns
}


