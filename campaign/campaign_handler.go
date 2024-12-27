package campaign

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-asteline-api/campaign/dto"
	"go-asteline-api/exception"
	"go-asteline-api/helper"
	"net/http"
)

type Handler struct {
	CampaignService Service
}

func NewHandler(campaignService Service) *Handler {
	return &Handler{CampaignService: campaignService}
}

func (campaignHandler *Handler) GetAll(ginContext *gin.Context) {
	allCampaignModels := campaignHandler.CampaignService.HandleGetAll(ginContext)
	ginContext.JSON(http.StatusOK, allCampaignModels)
}

func (campaignHandler *Handler) Create(ginContext *gin.Context) {
	var campaignCreateDto dto.CampaignCreateDto
	err := ginContext.Bind(&campaignCreateDto)
	helper.CheckErrorOperation(err, http.StatusBadRequest, exception.ErrInvalidRequestBody)
	imageFile, _ := ginContext.FormFile("imageFile")
	fmt.Println(campaignCreateDto)
	campaignCreateDto.ImageFile = imageFile

	campaignHandler.CampaignService.HandleCreate(ginContext, &campaignCreateDto, imageFile)
}

func (campaignHandler *Handler) Update(ginContext *gin.Context) {
	var campaignUpdateDto dto.CampaignUpdateDto
	err := ginContext.Bind(&campaignUpdateDto)
	helper.CheckErrorOperation(err, http.StatusBadRequest, exception.ErrInvalidRequestBody)
	imageFile, _ := ginContext.FormFile("image")
	campaignHandler.CampaignService.HandleUpdate(ginContext, &campaignUpdateDto, imageFile)
}

func (campaignHandler *Handler) Delete(ginContext *gin.Context) {
	campaignHandler.CampaignService.HandleDelete(ginContext)
}
