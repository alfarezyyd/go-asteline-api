package donation

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"go-asteline-api/donation/dto"
	"go-asteline-api/helper"
	"go-asteline-api/mapper"
	"go-asteline-api/model"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ServiceImpl struct {
	DonationRepository Repository
	dbConnection       *gorm.DB
	validationInstance *validator.Validate
	midtransCoreClient *coreapi.Client
}

func NewService(donationRepository Repository, dbConnection *gorm.DB, validationInstance *validator.Validate, midtransCoreClient *coreapi.Client) *ServiceImpl {
	return &ServiceImpl{
		DonationRepository: donationRepository,
		validationInstance: validationInstance,
		dbConnection:       dbConnection,
		midtransCoreClient: midtransCoreClient,
	}
}

func (donationHandler *ServiceImpl) HandleCreate(ginContext *gin.Context, donationCreateDto *dto.DonationCreateDto) {
	err := donationHandler.validationInstance.Struct(donationCreateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	//userClaims, isExists := ginContext.Get("claims")
	//if !isExists {
	//	ginContext.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
	//		"message": "User not found",
	//	})
	//	return
	//}
	//claimsMap, _ := userClaims.(jwt.MapClaims)
	//parsedClaimsMap, _ := mapper.MapJwtClaimIntoUserClaim(claimsMap)
	gormTransaction := donationHandler.dbConnection.Begin()
	defer helper.TransactionOperation(gormTransaction, ginContext)
	var campaignModel model.Campaign
	//gormTransaction.Where("email = ?", parsedClaimsMap.Email).First(&userModel)
	gormTransaction.Where("id = ?", donationCreateDto.CampaignId).First(&campaignModel)
	generatedUUID := uuid.New().String()
	donationModel, err := mapper.MapDonationDtoIntoDonationModel(donationCreateDto)
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
	donationModel.PaymentStatus = "Pending"
	donationModel.ID = generatedUUID
	gormTransaction.Create(&donationModel)
	chargeResponse, midtransError := donationHandler.midtransCoreClient.ChargeTransaction(&coreapi.ChargeReq{
		PaymentType: coreapi.PaymentTypeGopay,
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  generatedUUID,
			GrossAmt: donationCreateDto.Amount,
		},
		Items: &[]midtrans.ItemDetails{
			{
				ID:    strconv.FormatUint(campaignModel.ID, 10),
				Price: donationCreateDto.Amount,
				Name:  campaignModel.Title,
				Qty:   1,
			},
		},
		CustomerDetails: &midtrans.CustomerDetails{
			FName: donationCreateDto.Name,
			LName: "",
		},
	})
	if midtransError != nil && helper.CheckErrorOperation(midtransError.GetRawError(), ginContext, http.StatusBadRequest) {
		return
	}
	mapper.MapMidtransResponseIntoDonationModel(donationModel, chargeResponse)
	fmt.Println(donationModel)
	err = gormTransaction.Where("id = ?", generatedUUID).Updates(donationModel).Error
	if helper.CheckErrorOperation(err, ginContext, http.StatusBadRequest) {
		return
	}
}
