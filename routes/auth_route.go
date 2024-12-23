package routes

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/user"
)

func AuthRoute(ginEngine *gin.Engine, userController user.Controller) {
	ginEngine.POST("/register", userController.Register)
	ginEngine.POST("/login", userController.Login)
}
