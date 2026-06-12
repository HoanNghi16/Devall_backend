package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var blackList = map[string]bool{}

func LogoutHandler(cntx *gin.Context) {
	refresh,err := cntx.Cookie("refresh")
	if err == nil{
		BlackListToken(refresh)
		cntx.SetCookie("access", "", -1, "/","",false, true ) //Để client xóa tokens ngay lập tức
		cntx.SetCookie("refresh", "", -1, "/","",false, true )
		cntx.JSON(http.StatusAccepted, gin.H{
			"message": "Đăng xuất thành công!",
		})
		return
	}
	cntx.JSON(http.StatusUnauthorized, gin.H{
		"message": "Bạn chưa đăng nhập!",
	})
}
