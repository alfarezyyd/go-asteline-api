package model

import "time"

type Campaign struct {
	ID            uint64    `gorm:"primary_key;auto_increment;->:"`
	Title         string    `gorm:"column:title;"`
	Description   string    `gorm:"column:description"`
	GoalAmount    float64   `gorm:"column:goal_amount"`
	CurrentAmount uint64    `gorm:"column:current_amount"`
	StartDate     string    `gorm:"column:start_date"`
	EndDate       string    `gorm:"column:end_date"`
	ImageUrl      string    `gorm:"column:image_url"`
	status        string    `gorm:"column:status"`
	UserId        uint64    `gorm:"column:user_id"`
	User          User      `gorm:"foreignKey:user_id;references:id"`
	CreatedAt     time.Time `gorm:"column:created_at;autoCreateTime" mapstructure:"created_at"`
	UpdatedAt     time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" mapstructure:"updated_at"`
}
