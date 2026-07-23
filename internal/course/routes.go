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

	router.GET("/course/courses", handler.CoursesHandler)
	router.GET("/course/topics", handler.Topics)

	notPrivate := router.Group("/course", auth.OptionalAuth())
	{
		notPrivate.GET("/:id", handler.GetFullCourseHandler)
	}

	private := router.Group("/course", auth.AuthRequired())
	{
		private.GET("/history", handler.GetHistory)
		private.PATCH("/:id", handler.UpdateCourseUser) // update quan hệ của khóa học và người dùng
		private.POST("/my", handler.CreateNewCourse)
		private.GET("/my", handler.MyCourses)
	}
}