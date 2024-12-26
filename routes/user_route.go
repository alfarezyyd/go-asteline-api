package routes

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/campaign"
	"go-asteline-api/category"
)

func UserRoute(routerGroup *gin.RouterGroup, campaignController campaign.Controller, categoryController category.Controller) {
	campaignGroup := routerGroup.Group("/campaigns")
	campaignGroup.POST("", campaignController.Create)
	campaignGroup.PUT("/:id", campaignController.Update)
	campaignGroup.DELETE("/:id", campaignController.Delete)

	categoryGroup := routerGroup.Group("/categories")
	categoryGroup.GET("", categoryController.GetAll)
	categoryGroup.POST("", categoryController.Create)
	categoryGroup.PUT("/:id", categoryController.Update)
	categoryGroup.DELETE("/:id", categoryController.Delete)

}
