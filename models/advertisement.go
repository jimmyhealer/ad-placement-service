package models

import (
	"time"

	"gorm.io/gorm"
)

type Advertisement struct {
	gorm.Model
	ID         uint         `gorm:"primaryKey" `
	Title      string       `gorm:"type:varchar(255);not null" form:"title" binding:"required"`
	StartAt    time.Time    `gorm:"type:timestamp with time zone;not null;index:idx_member,priority:2" form:"start_at" binding:"required"`
	EndAt      time.Time    `gorm:"type:timestamp with time zone;not null;index:idx_member,priority:1" form:"end_at" binding:"required"`
	Conditions []Conditions `form:"conditions" binding:"required"`
}

type GetAdsRequest struct {
	Offset   int            `form:"offset,default=0" binding:"omitempty,min=1,max=100"`
	Limit    int            `form:"limit,default=5" binding:"omitempty,min=1,max=100"`
	Age      int            `form:"age" binding:"omitempty,min=1,max=100"`
	Gender   []GenderType   `form:"gender" binding:"omitempty,dive,oneof=M F"`
	Country  []CountryCode  `form:"country" binding:"omitempty,dive,oneof=TW JP"`
	Platform []PlatformType `form:"platform" binding:"omitempty,dive,oneof=android ios web"`
}
