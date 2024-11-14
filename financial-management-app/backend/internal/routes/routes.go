package routes

import "github.com/gin-gonic/gin"

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.POST("/register", handlers.RegisterUser)
	router.POST("login", handlers.LoginUser)
	autorized := router.Group("/")
	autorized.Use(handlers.AuthMiddleware())

	return router
}
