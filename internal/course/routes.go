package course

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CourseRoutes(router *gin.Engine, db *gorm.DB){

	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router.GET("/course/:id", handler.GetFullCourseHandler)
	router.GET("/course/courses", handler.CoursesHandler)
}