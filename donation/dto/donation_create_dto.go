package dto

type DonationCreateDto struct {
	Id         string
	CampaignId string `json:"campaign_id" validate:"required;min=0"`
	Name       string `json:"name" validate:"required;min=3"`
	Amount     int64  `json:"amount" validate:"required;min=0"`
}
