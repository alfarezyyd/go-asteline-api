package dto

type CampaignCreateDto struct {
	Title       string `json:"title" validate:"required,min=3,max=100"`
	Description string `json:"description" validate:"required,min=30"`
	GoalAmount  uint64 `json:"goalAmount" validate:"required,gte=0"`
	StartDate   string `json:"startDate" validate:"required,date"`
	EndDate     string `json:"endDate" validate:"required,date"`
}
