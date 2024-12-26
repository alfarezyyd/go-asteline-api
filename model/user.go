package model

import "time"

type User struct {
	ID             uint64     `gorm:"primary_key,auto_increment"`
	Email          string     `mapstructure:"Email"`
	Password       string     `mapstructure:"Password"`
	FullName       string     `gorm:"column:full_name" mapstructure:"FullName"`
	BirthDate      time.Time  `gorm:"column:birth_date" mapstructure:"BirthDate"`
	Gender         string     `mapstructure:"gender"`
	PhoneNumber    string     `gorm:"column:phone_number" mapstructure:"PhoneNumber"`
	ProfilePicture string     `gorm:"column:profile_picture" mapstructure:"ProfilePicture"`
	TotalDonations float64    `gorm:"column:total_donations" mapstructure:"TotalDonations"`
	Campaign       []Campaign `gorm:"foreignKey:user_id;references:id"`
	Donation       []Donation `gorm:"foreignKey:user_id;references:id"`
	CreatedAt      time.Time  `gorm:"column:created_at;autoCreateTime" mapstructure:"created_at"`
	UpdatedAt      time.Time  `gorm:"column:updated_at;autoCreateTime;autoUpdateTime" mapstructure:"updated_at"`
}
