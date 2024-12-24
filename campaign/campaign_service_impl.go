package campaign

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go-asteline-api/campaign/dto"
	"go-asteline-api/helper"
	"go-asteline-api/mapper"
	"go-asteline-api/model"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
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

func (serviceImpl *ServiceImpl) HandleCreate(ginContext *gin.Context, campaignCreateDto *dto.CampaignCreateDto, multipartFile *multipart.FileHeader) {
	err := serviceImpl.structValidator.Struct(campaignCreateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	campaignModel, err := mapper.MapCampaignCreateDtoIntoCampaignModel(campaignCreateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	userClaims, isExists := ginContext.Get("claims")
	if !isExists {
		ginContext.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "User not found",
		})
		return
	}
	claimsMap, _ := userClaims.(jwt.MapClaims)
	parsedClaimsMap, _ := mapper.MapJwtClaimIntoUserClaim(claimsMap)
	gormTransaction := serviceImpl.dbConnection.Begin()
	var userModel model.User
	dbConn := gormTransaction.Where("email = ?", parsedClaimsMap.Email).First(&userModel)
	fmt.Println(userModel)
	if helper.CheckErrorOperation(dbConn.Error, ginContext, http.StatusBadRequest) {
		return
	}
	campaignModel.UserId = userModel.ID
	dbConn = gormTransaction.Create(&campaignModel)
	if helper.CheckErrorOperation(dbConn.Error, ginContext, http.StatusBadRequest) {
		return
	}
	err = ginContext.SaveUploadedFile(multipartFile, fmt.Sprintf("public/assets/%d/%s", campaignModel.ID, multipartFile.Filename))
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	helper.TransactionOperation(gormTransaction)
	ginContext.JSON(http.StatusCreated, campaignModel)
}
