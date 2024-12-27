package campaign

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go-asteline-api/campaign/dto"
	"go-asteline-api/exception"
	"go-asteline-api/helper"
	"go-asteline-api/mapper"
	"go-asteline-api/model"
	"gorm.io/gorm"
	"mime/multipart"
	"net/http"
	"os"
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

func (serviceImpl *ServiceImpl) HandleGetAll(ginContext *gin.Context) []model.Campaign {
	var allCampaignsModel []model.Campaign
	serviceImpl.dbConnection.Find(&allCampaignsModel)
	return allCampaignsModel
}

func (serviceImpl *ServiceImpl) HandleCreate(ginContext *gin.Context, campaignCreateDto *dto.CampaignCreateDto, multipartFile *multipart.FileHeader) {
	err := serviceImpl.structValidator.Struct(campaignCreateDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	campaignModel, err := mapper.MapCampaignCreateDtoIntoCampaignModel(campaignCreateDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
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
	operationResult := gormTransaction.Where("email = ?", parsedClaimsMap.Email).First(&userModel).Error
	helper.CheckErrorOperation(operationResult, exception.ParseGormError(operationResult))
	err = ginContext.SaveUploadedFile(multipartFile, fmt.Sprintf("public/assets/%d/%s", campaignModel.ID, multipartFile.Filename))
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	campaignModel.UserId = userModel.ID
	campaignModel.ImageUrl = multipartFile.Filename
	err = gormTransaction.Create(&campaignModel).Error
	helper.CheckErrorOperation(err, exception.ParseGormError(err))
	helper.TransactionOperation(gormTransaction, ginContext)
	ginContext.JSON(http.StatusCreated, campaignModel)
}

func (serviceImpl *ServiceImpl) HandleUpdate(ginContext *gin.Context, campaignUpdateDto *dto.CampaignUpdateDto, multipartFile *multipart.FileHeader) {
	err := serviceImpl.structValidator.Struct(campaignUpdateDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	campaignModel, err := mapper.MapCampaignCreateDtoIntoCampaignModel(campaignUpdateDto)
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	userClaims, _ := ginContext.Get("claims")
	claimsMap, _ := userClaims.(jwt.MapClaims)
	parsedClaimsMap, _ := mapper.MapJwtClaimIntoUserClaim(claimsMap)
	gormTransaction := serviceImpl.dbConnection.Begin()

	var userModel model.User
	var existingCampaignModel model.Campaign
	gormTransaction.Where("email = ?", parsedClaimsMap.Email).First(&userModel)

	if multipartFile != nil {
		err := gormTransaction.First(&existingCampaignModel, "id = ? AND user_id = ?", ginContext.Param("id"), userModel.ID).Error
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))

		err = os.Remove(fmt.Sprintf("public/assets/%d/%s", existingCampaignModel.ID, existingCampaignModel.ImageUrl))
		helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))

	}
	gormTransaction.Where("id = ? AND user_id = ?", ginContext.Param("id"), userModel.ID).Updates(campaignModel)
	helper.TransactionOperation(gormTransaction, ginContext)
	ginContext.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success",
	})
}

func (serviceImpl *ServiceImpl) HandleDelete(ginContext *gin.Context) {
	userClaims, _ := ginContext.Get("claims")
	claimsMap, _ := userClaims.(jwt.MapClaims)
	parsedClaimsMap, _ := mapper.MapJwtClaimIntoUserClaim(claimsMap)
	gormTransaction := serviceImpl.dbConnection.Begin()
	var userModel model.User
	var existingCampaignModel model.Campaign
	gormTransaction.Where("email = ?", parsedClaimsMap.Email).First(&userModel)
	err := gormTransaction.Where("id = ? AND user_id = ?", ginContext.Param("id"), userModel.ID).First(&existingCampaignModel).Error
	helper.CheckErrorOperation(err, exception.NewClientError(http.StatusBadRequest, exception.ErrInvalidRequestBody))
	{
		gormTransaction.Delete(&existingCampaignModel)
		helper.TransactionOperation(gormTransaction, ginContext)
	}
}
