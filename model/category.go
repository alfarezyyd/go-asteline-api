package model

import "time"

type Category struct {
	Id          uint32    `gorm:"column:id;primary_key;auto_increment"`
	Name        string    `gorm:"column:name" mapstructure:"Name"`
	Description string    `gorm:"column:description;type:text" mapstructure:"Description"`
	CreatedAt   time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt   time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}
