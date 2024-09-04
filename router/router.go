package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// init Router with gin default settings
	router := gin.Default()
	// define route
	router.GET("/ping", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "pong",
		})
	})
	// listen and serve on 0.0.0.0:8080
	router.Run()
}
