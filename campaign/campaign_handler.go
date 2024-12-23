package campaign

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/campaign/dto"
	"go-asteline-api/helper"
	"net/http"
)

type Handler struct {
	CampaignService Service
}

func NewHandler(campaignService Service) *Handler {
	return &Handler{CampaignService: campaignService}
}

func (campaignHandler *Handler) Create(ginContext *gin.Context) {
	var campaignCreateDto dto.CampaignCreateDto
	err := ginContext.ShouldBindJSON(&campaignCreateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	campaignHandler.CampaignService.HandleCreate(ginContext, &campaignCreateDto)
}
