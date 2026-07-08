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

func (h *Handler) ProfileHandler(cntx *gin.Context){
 	userID, ok := cntx.Get("userID") ;
	if !ok{
		cntx.JSON(http.StatusNoContent, gin.H{
			"message": "Bạn chưa đăng nhập!",
		})
		return
	}
	profile,err := h.service.GetProfile(userID.(uint))

	if err != nil{
		cntx.JSON(http.StatusForbidden, gin.H{
			"message": err.Error(),
		})
		return
	}
	cntx.JSON(http.StatusOK, profile)
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
	tokens,err := h.service.Login(&input); 
	if err != nil{
		cntx.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	cntx.SetCookie("access", tokens.Access, 15*60, "/", "", false, true)
	cntx.SetCookie("refresh", tokens.Refresh, 60*60*24*7, "/","", false, true)

	cntx.JSON(200, gin.H{
		"message": "Đăng nhập thành công",
	})
}