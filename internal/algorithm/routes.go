package algorithm

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func AlgorithmRoutes(router *gin.Engine, db *gorm.DB){
	repository := NewRepository(db)
	service := NewService(repository)
	handler := NewHandler(service)

	router.GET("/algorithm/algorithms", handler.AlgorithmList)
	router.GET("/algorithm/:id", handler.GetAlgorithm)
}