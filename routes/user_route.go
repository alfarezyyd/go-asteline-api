package routes

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/user"
)

func UserRoute(route *gin.Engine, userController *user.Handler) {
	userRoutes := route.Group("/user")
	userRoutes.POST("/register")
}
