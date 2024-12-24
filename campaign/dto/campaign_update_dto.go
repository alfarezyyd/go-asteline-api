package dto

type CampaignUpdateDto struct {
	Title       string `form:"title" json:"title" validate:"required,min=3,max=100"`
	Description string `form:"description" json:"description" validate:"required,min=30"`
	GoalAmount  uint64 `form:"goalAmount" json:"goalAmount" validate:"required,gte=0"`
	StartDate   string `form:"startDate" json:"startDate" validate:"required,date"`
	EndDate     string `form:"endDate" json:"endDate" validate:"required,date"`
}
