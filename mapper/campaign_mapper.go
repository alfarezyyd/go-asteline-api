package mapper

import (
	"github.com/go-viper/mapstructure/v2"
	"go-asteline-api/campaign/dto"
	"go-asteline-api/model"
)

func MapCampaignCreateDtoIntoCampaignModel[T *dto.CampaignCreateDto | *dto.CampaignUpdateDto](campaignSaveDto T) (*model.Campaign, error) {
	var modelCampaign model.Campaign
	err := mapstructure.Decode(campaignSaveDto, &modelCampaign)
	if err != nil {
		return nil, err
	}
	return &modelCampaign, nil
}
