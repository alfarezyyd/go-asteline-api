package routes

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/campaign"
)

func UserRoute(routerGroup *gin.RouterGroup, campaignController campaign.Controller) {
	campaignGroup := routerGroup.Group("/campaigns")
	campaignGroup.POST("", campaignController.Create)
	campaignGroup.PUT("/:id", campaignController.Update)
	campaignGroup.DELETE("/:id", campaignController.Delete)
}
