package mapper

import (
	"github.com/go-viper/mapstructure/v2"
	"go-asteline-api/campaign/dto"
	"go-asteline-api/model"
)

func MapCampaignCreateDtoIntoCampaignModel(campaignCreateDto *dto.CampaignCreateDto) (*model.Campaign, error) {
	var modelCampaign model.Campaign
	err := mapstructure.Decode(campaignCreateDto, &modelCampaign)
	if err != nil {
		return nil, err
	}
	return &modelCampaign, nil
}
