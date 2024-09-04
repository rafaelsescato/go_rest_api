package router

import (
	"github.com/gin-gonic/gin"

	"github.com/rafaelsescato/go_rest_api/handler"
)

func initializeRoutes(router *gin.Engine) {
	router_group_v1 := router.Group("/api/v1")
	{
		router_group_v1.POST("/opening", handler.CreateOpeningHandler)
		router_group_v1.GET("/opening", handler.ShowOpeningHandler)
		router_group_v1.GET("/openings", handler.ListOpeningsHandler)
		router_group_v1.PUT("/opening", handler.UpdateOpeningHandler)
		router_group_v1.DELETE("/opening", handler.DeleteOpeningHandler)
	}
}
