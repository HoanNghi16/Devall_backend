package media

import (
	"github.com/HoanNghi16/Devall_backend/internal/auth"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func MediaRoutes(router *gin.Engine, db *gorm.DB, cld *cloudinary.Cloudinary) {
	repository := NewRepository(db)
	service := NewService(repository, cld)
	handler := NewHandler(service)

	private_group := router.Group("/media", auth.AuthRequired())
	{
		private_group.POST("/upload", handler.MediaUpload)
	}
}