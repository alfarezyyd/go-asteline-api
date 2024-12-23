package routes

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/user"
)

func UserRoute(route *gin.Engine, userController user.Controller) {
	userRouteGroup := route.Group("/user")
	userRouteGroup.POST("/register", userController.Register)
	userRouteGroup.POST("/login", userController.Login)
}
