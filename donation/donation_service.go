package donation

import (
	"github.com/gin-gonic/gin"
	"go-asteline-api/donation/dto"
)

type Service interface {
	HandleCreate(ginContext *gin.Context, createDonationDto *dto.DonationCreateDto)
}
