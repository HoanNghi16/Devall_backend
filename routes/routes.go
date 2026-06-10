package routes

import (
	"github.com/HoanNghi16/Devall_backend/internal/user"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//SetupRouter and API path
func SetupRouter (db *gorm.DB) *gin.Engine{
	router := gin.Default()

	//Cho phép port 3000 gửi request qua cors
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string {
			"POST", "PUT", "PATCH", "GET", "DELETE",
		},
	}))

	//Setup API

	user.UserRoutes(router, db)
	// router.GET("/test", func ( cntx *gin.Context) {
	// 	cntx.JSON(200, gin.H{"message": "Test Ngon lành"})
	// } )


	return router
}
