package router

import "github.com/gin-gonic/gin"

func Initialize() {
	// init Router with gin default settings
	router := gin.Default()
	// define routes
	initializeRoutes(router)
	// listen and serve on 0.0.0.0:8080
	router.Run()
}
