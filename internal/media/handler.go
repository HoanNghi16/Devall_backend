package media

import (
	"log"

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

func (handler *Handler)MediaUpload(cntx *gin.Context) {
	var req MediaRequest

	if err := cntx.ShouldBind(&req); err != nil{
		log.Print(err.Error())
		cntx.JSON(400, gin.H{
			"message": "Dữ liệu không hợp lệ!",
		})
		return
	}

	userID,ok := cntx.Get("userID")
	if !ok{
		cntx.JSON(401, gin.H{"message":"Người dùng chưa đăng nhập!"})
		return
	}

	url, err :=handler.service.UploadMedia(cntx.Request.Context(),&req, userID.(uint))
	if err != nil{
		cntx.JSON(500, gin.H{"message": err.Error()})
		return
	}

	cntx.JSON(201, gin.H{"message": "Upload thành công!", "url": url})

}