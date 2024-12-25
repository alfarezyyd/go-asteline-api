package routes

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/campaign"
	"go-asteline-api/user"
)

func PublicRoute(ginEngine *gin.Engine, userController user.Controller, campaignController campaign.Controller) {
	ginEngine.POST("/login", userController.Login)
	ginEngine.POST("/register", userController.Register)

	ginEngine.GET("/campaigns", campaignController.GetAll)
}
