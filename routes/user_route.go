package routes

import "github.com/gin-gonic/gin"

func User(route *gin.Engine) {
	userRoutes := route.Group("/api/user")
	userRoutes.POST("", func(c *gin.Context) {

	})
}
