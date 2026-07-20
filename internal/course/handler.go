package course

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
	}
}

// GET /course/courses
func (handler *Handler) CoursesHandler(cntx *gin.Context) {
	
	var filter CourseFilter
	
	if err:= cntx.ShouldBindQuery(&filter); err!=nil{
		cntx.JSON(400,gin.H{
			"message": "Bộ lọc không hợp lệ!",
		})
		return
	}

	courses, err := handler.service.ListCourseService(filter.Cursor, filter.TopicIDs, filter.Level)

	if err != nil{
		cntx.JSON(403, gin.H{
			"message": "Bạn không có quyền truy cập!",
		})
		return
	}
	cntx.JSON(200, courses)
}




// GET /course/:id
func (handler *Handler) GetFullCourseHandler(cntx *gin.Context){
	
	id, ok := cntx.Params.Get("id")
	if !ok{
		cntx.JSON(400, gin.H{
			"message": "Lỗi!",
		})
		return
	}

	id1, err1 := strconv.ParseUint(id, 10, 64)

	if err1 != nil{
		cntx.JSON(400, gin.H{
			"message": "ID khóa học không hợp lệ!",
		})
		return
	}

	userID, ok := cntx.Get("userID")
	if !ok{
		userID = uint(0)
	}

	course, err := handler.service.CourseFullService(uint(id1), userID.(uint))
	if err!=nil{
		cntx.JSON(404, gin.H{
			"message": "Không tìm thấy khóa học",
		})
		return
	}
	cntx.JSON(200, course)
}



// POST /course/my
func (handler *Handler) CreateNewCourse(cntx *gin.Context){
	var newCourse RequestCourse

	if err := cntx.ShouldBindJSON(&newCourse); err != nil{
		cntx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	id, ok := cntx.Get("userID")
	
	if !ok{
		cntx.JSON(400, gin.H{
			"message":"ID người dùng không hợp lệ!",
		})
		return
	}


	if err := handler.service.CreateMyCourse(id.(uint), &newCourse); err != nil{
		cntx.JSON(400, gin.H{
			"message": err.Error(),
		})
		return
	}

	cntx.JSON(200, gin.H{
		"message": "Thêm khóa học thành công!",
	})
}


// GET /course/my
func (handler *Handler) MyCourses(cntx * gin.Context){
	userID, ok := cntx.Get("userID")
	if ok{
		courses,err := handler.service.MyCourseService(userID.(uint))
		if err != nil{
			cntx.JSON(400, gin.H{
				"message": err,
			})
			return
		}
		cntx.JSON(200, courses)
	}
}



func (handler *Handler) Topics (cntx *gin.Context){
	topics, err := handler.service.GetTopics()

	if err != nil{
		cntx.JSON(404, gin.H{"message":  err.Error()})
		return
	}

	cntx.JSON(200, topics)
}