package campaign

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/campaign/dto"
)

type Service interface {
	HandleCreate(ginContext *gin.Context, campaignCreateDto *dto.CampaignCreateDto)
}
