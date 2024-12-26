package mapper

import (
	"github.com/go-viper/mapstructure/v2"
	"github.com/midtrans/midtrans-go/coreapi"
	"go-asteline-api/donation/dto"
	"go-asteline-api/model"
)

func MapDonationDtoIntoDonationModel[T *dto.DonationCreateDto](donationDto T) (*model.Donation, error) {
	var donationModel model.Donation
	err := mapstructure.Decode(donationDto, &donationModel)
	if err != nil {
		return nil, err
	}
	return &donationModel, nil
}

func MapMidtransResponseIntoDonationModel(donationModel *model.Donation, midtransResponse *coreapi.ChargeResponse) {
	donationModel.PaymentFraudStatus = midtransResponse.FraudStatus
	donationModel.TransactionId = midtransResponse.TransactionID
	donationModel.PaymentType = midtransResponse.PaymentType
}
