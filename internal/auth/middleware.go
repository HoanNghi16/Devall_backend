package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc{
	return func(cntx *gin.Context){
		access, err := cntx.Cookie("access")

		if err != nil{
			cntx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Bạn chưa đăng nhập!",
			} )
			cntx.Abort()
			return 
		}

		claims, err := VerifyToken(access, "SECRET_KEY")
		
		if err != nil{
			cntx.JSON(http.StatusUnauthorized, gin.H{
				"message": "Token không hợp lệ!",
			})
			cntx.Abort()
			return
		}

		cntx.Set("userID", claims.UserID)
		cntx.Set("role",claims.Role)

		cntx.Next()
	}
}