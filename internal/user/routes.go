package user

import (
	"github.com/HoanNghi16/Devall_backend/internal/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRoutes(router *gin.Engine, db *gorm.DB){
	repo := NewRepository(db)
	service := NewService(repo)
	handler := NewHandler(service)

	//Public handler
	router.POST("/user/register", handler.RegisterHandler)
	router.POST("/user/login", handler.LoginHandler)
	router.GET("/user/refresh", handler.RefreshTokenHandler)

	protected := router.Group("/user", auth.AuthRequired())
	{
		protected.GET("/profile", handler.ProfileHandler)
	}
}