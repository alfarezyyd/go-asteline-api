package donation

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/donation/dto"
	"go-asteline-api/helper"
	"net/http"
)

type Handler struct {
	donationService Service
}

func NewHandler() *Handler {
	return &Handler{}

}

func (donationHandler *Handler) Create(ginContext *gin.Context) {
	var donationCreateDto dto.DonationCreateDto
	err := ginContext.ShouldBindBodyWithJSON(&donationCreateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	donationHandler.donationService.HandleCreate(ginContext, &donationCreateDto)
	ginContext.JSON(http.StatusCreated, gin.H{"status": "created"})
}
