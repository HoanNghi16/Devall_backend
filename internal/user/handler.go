package user

import (
	"net/http"

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

func (h *Handler) RegisterHandler(cntx *gin.Context) {
	var input RegisterRequest

	if err := cntx.ShouldBindJSON(&input); err != nil{
		cntx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	err := h.service.Register(&input)
	if err != nil {
		cntx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		} )
		return
	}
	cntx.JSON(http.StatusAccepted, gin.H{
		"message": "Đăng ký thành công",
	} )
}


func (h *Handler) LoginHandler(cntx *gin.Context){
	var input LoginRequest
	if err := cntx.ShouldBindJSON(&input); err != nil{
		cntx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if err := h.service.Login(&input); err != nil{
		cntx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	cntx.JSON(200, gin.H{
		"message": "Đăng nhập thành công",
	})
}