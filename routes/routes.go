package routes

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

//SetupRouter and API path
func SetupRouter () *gin.Engine{
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{
			"http://localhost:3000",
		},
		AllowMethods: []string {
			"POST", "PUT", "PATCH", "GET", "DELETE",
		},
	}))

	router.GET("/test", func ( cntx *gin.Context) {
		cntx.JSON(200, gin.H{"message": "Test Ngon lành"})
	} )

	return router
}
