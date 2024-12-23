package routes

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/user"
)

func PublicRoute(ginEngine *gin.Engine, userController user.Controller) {
	ginEngine.POST("/login", userController.Login)
	ginEngine.POST("/register", userController.Register)
}
