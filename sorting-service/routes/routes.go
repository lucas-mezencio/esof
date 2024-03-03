package routes

import (
	"sorting-service/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/sort", handlers.SortHandler)
	return r
}
