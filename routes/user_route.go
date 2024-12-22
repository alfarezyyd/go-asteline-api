package routes

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/user"
)

func UserRoute(route *gin.Engine, userController user.Controller) {
	route.Group("/user")
	route.POST("/register", userController.Register)
}
