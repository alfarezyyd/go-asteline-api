package dto

type DonationCreateDto struct {
	CampaignId uint64 `json:"campaign_id" validate:"required,gte=0"`
	Name       string `json:"name" validate:"required,min=3"`
	Amount     int64  `json:"amount" validate:"required,gte=0"`
}
