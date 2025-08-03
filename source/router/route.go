package router

import (
	"demo_challenges/source/module/authentication"
	"demo_challenges/source/module/file"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.Default()

	// Initialize routes
	authentication.New().Router(router)
	file.New().Router(router)
	return router
}
