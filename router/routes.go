package router

import (
	"github.com/Reydner96/gopportunities/handler"
	"github.com/gin-gonic/gin"
)

func initializerRoutes(router *gin.Engine) {
	v1 := router.Group("/api/v1")
	{
		// Show Opening
		v1.GET("/opening", handler.ShowOpeningHandler)
		// Post Opening
		v1.POST("/opening", handler.CreateOpeningHandler)
		// Delete Opening
		v1.DELETE("/opening", handler.DeleteOpeningHandler)
		// Update Opening
		v1.PUT("/opening", handler.UpdateOpeningHandler)
		// List Openings
		v1.GET("/openings", handler.ListOpeningHandler)
	}
}
