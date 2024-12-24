package campaign

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/campaign/dto"
	"mime/multipart"
)

type Service interface {
	HandleCreate(ginContext *gin.Context, campaignCreateDto *dto.CampaignCreateDto, multipartFile *multipart.FileHeader)
}
