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
	imageFile, _ := ginContext.FormFile("imageFile")
	fmt.Println(campaignCreateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	campaignCreateDto.ImageFile = imageFile

	campaignHandler.CampaignService.HandleCreate(ginContext, &campaignCreateDto, imageFile)
}

func (campaignHandler *Handler) Update(ginContext *gin.Context) {
	var campaignUpdateDto dto.CampaignUpdateDto
	err := ginContext.Bind(&campaignUpdateDto)
	imageFile, _ := ginContext.FormFile("image")
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	campaignHandler.CampaignService.HandleUpdate(ginContext, &campaignUpdateDto, imageFile)
}

func (campaignHandler *Handler) Delete(ginContext *gin.Context) {
	campaignHandler.CampaignService.HandleDelete(ginContext)
}
