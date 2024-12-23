package campaign

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ServiceImpl struct {
	campaignRepository Repository
	dbConnection       *gorm.DB
	structValidator    *validator.Validate
}

func NewService(dbConnection *gorm.DB, campaignRepository Repository, structValidator *validator.Validate) *ServiceImpl {
	return &ServiceImpl{
		dbConnection:       dbConnection,
		campaignRepository: campaignRepository,
		structValidator:    structValidator,
	}
}

func (serviceImpl *ServiceImpl) HandleCreate(ginContext *gin.Context) {

}
