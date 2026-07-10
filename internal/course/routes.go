package course

import (
	"github.com/HoanNghi16/Devall_backend/internal/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CourseRoutes(router *gin.Engine, db *gorm.DB){

	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router.GET("/course/:id", handler.GetFullCourseHandler)
	router.GET("/course/courses", handler.CoursesHandler)
	router.GET("/course/topics", handler.Topics)

	private := router.Group("/course", auth.AuthRequired())
	{
		private.POST("/my", handler.CreateNewCourse)
		private.GET("/my", handler.MyCourses)
	}
}