package user

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB){
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	user := router.Group("/user")
	{
		user.POST("/register", handler.RegisterHandler)
		user.POST("/login", handler.LoginHandler)
	}
}