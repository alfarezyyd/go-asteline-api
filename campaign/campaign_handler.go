package campaign

import "github.com/gin-gonic/gin"

type Handler struct {
	CampaignService Service
}

func NewHandler(campaignService Service) *Handler {
	return &Handler{CampaignService: campaignService}
}

func (campaignHandler *Handler) Create(ginContext *gin.Context) {

}
