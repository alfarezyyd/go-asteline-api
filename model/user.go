package model

import "time"

type User struct {
	id             uint64 `gorm:"primary_key,auto_increment"`
	email          string
	password       string
	fullName       string    `gorm:"column:full_name"`
	birthDate      time.Time `gorm:"column:birth_date"`
	gender         string
	phoneNumber    string    `gorm:"column:phone_number"`
	profilePicture string    `gorm:"column:profile_picture"`
	totalDonations float64   `gorm:"column:total_donations"`
	createdAt      time.Time `gorm:"column:created_at;autoCreateTime"`
	updatedAt      time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
