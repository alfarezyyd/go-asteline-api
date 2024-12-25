package campaign

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/campaign/dto"
	"go-asteline-api/model"
	"mime/multipart"
)

type Service interface {
	HandleGetAll(context *gin.Context) []model.Campaign
	HandleCreate(ginContext *gin.Context, campaignCreateDto *dto.CampaignCreateDto, multipartFile *multipart.FileHeader)
	HandleUpdate(ginContext *gin.Context, campaignUpdateDto *dto.CampaignUpdateDto, multipartFile *multipart.FileHeader)
	HandleDelete(ginContext *gin.Context)
}
