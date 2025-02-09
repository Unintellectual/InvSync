package routes

import (
	"github.com/Unintellectual/handlers"

	"github.com/gin-gonic/gin"
)

// RegisterRoutes sets up the API endpoints
func RegisterRoutes(r *gin.Engine) {
	r.GET("/items", handlers.GetFoodItems)
}
