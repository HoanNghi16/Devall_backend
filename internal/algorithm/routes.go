package algorithm

import (
	"github.com/HoanNghi16/Devall_backend/internal/auth"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AlgorithmRoutes(router *gin.Engine, db *gorm.DB){
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	notProtected := router.Group("/algorithm", auth.OptionalAuth())
	{
		notProtected.GET("/algorithms", handler.AlgorithmList)
	}

	router.GET("/algorithm/tags", handler.TagsHandler)
	router.GET("/algorithm/:id", handler.GetAlgorithm)
}