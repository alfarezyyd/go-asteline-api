package donation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-asteline-api/donation/dto"
	"go-asteline-api/exception"
	"go-asteline-api/helper"
	"net/http"
)

type Handler struct {
	donationService Service
}

func NewHandler(donationService Service) *Handler {
	return &Handler{
		donationService: donationService,
	}
}

func (donationHandler *Handler) Create(ginContext *gin.Context) {
	var donationCreateDto dto.DonationCreateDto
	err := ginContext.ShouldBindBodyWithJSON(&donationCreateDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	fmt.Println(ginContext)
	donationHandler.donationService.HandleCreate(ginContext, &donationCreateDto)
	ginContext.JSON(http.StatusCreated, gin.H{"status": "created"})
}

func (donationHandler *Handler) Notification(ginContext *gin.Context) {
	var donationNotificationDto dto.DonationNotificationDto
	err := ginContext.ShouldBindBodyWithJSON(&donationNotificationDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
}
