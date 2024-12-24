package campaign

import (
	"fmt"
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
	err := ginContext.Bind(&campaignCreateDto)
	imageFile, _ := ginContext.FormFile("image")
	fmt.Println(campaignCreateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	campaignHandler.CampaignService.HandleCreate(ginContext, &campaignCreateDto, imageFile)
}
