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

func (handler *Handler) CoursesHandler(cntx *gin.Context) {
	id, err := strconv.ParseUint(cntx.Query("cursor"), 10, 64)
	
	if err!=nil {
		courses,_ := handler.service.ListCourseService(uint(0))
		cntx.JSON(200, courses)
		return
	}
	courses, err := handler.service.ListCourseService(uint(id))

	if err != nil{
		cntx.JSON(403, gin.H{
			"message": "Bạn không có quyền truy cập!",
		})
	}
	cntx.JSON(200, courses)
}

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

	course, err := handler.service.CourseFullService(uint(id1))
	if err!=nil{
		cntx.JSON(404, gin.H{
			"message": "Không tìm thấy khóa học",
		})
		return
	}
	cntx.JSON(200, course)
}


func (handler *Handler) MyCourses(cntx * gin.Context){
	userID, ok := cntx.Get("userID")
	if ok{
		courses,err := handler.service.repository.GetMyCourses(userID.(uint))
		if err != nil{
			cntx.JSON(400, gin.H{
				"message": err,
			})
			return
		}
		cntx.JSON(200, courses)
	}
}