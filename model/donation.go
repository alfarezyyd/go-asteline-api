package model

import "time"

type Donation struct {
	ID                 string    `gorm:"column:id;primary_key" mapstructure:"Id"`
	UserId             uint64    `gorm:"column:user_id;foreignKey"`
	CampaignId         uint64    `gorm:"column:campaign_id" mapstructure:"CampaignId"`
	User               *User     `gorm:"foreignKey:user_id;references:id"`
	Campaign           *Campaign `gorm:"foreignKey:campaign_id;references:campaign_id"`
	TransactionId      string    `gorm:"column:transaction_id"`
	Name               string    `gorm:"column:name" mapstructure:"Name"`
	Amount             float64   `gorm:"column:amount" mapstructure:"Amount"`
	PaymentStatus      string    `gorm:"column:payment_status"`
	PaymentType        string    `gorm:"column:payment_type"`
	PaymentFraudStatus string    `gorm:"column:payment_fraud_status"`
	CreatedAt          time.Time `gorm:"column:created_at;autoCreateTime'"`
	UpdatedAt          time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime'"`
}
