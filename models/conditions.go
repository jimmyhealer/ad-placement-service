package models

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Conditions struct {
	gorm.Model
	AdvertisementID uint
	AgeStart        int            `gorm:"index" form:"age_start" binding:"omitempty,min=1,max=100"`
	AgeEnd          int            `gorm:"index" form:"age_end" binding:"omitempty,min=1,max=100"`
	Gender          pq.StringArray `gorm:"type:text[];index" form:"gender" binding:"omitempty,dive,oneof=M F"`
	Country         pq.StringArray `gorm:"type:text[];index" form:"country" binding:"omitempty,dive,oneof=TW JP"`
	Platform        pq.StringArray `gorm:"type:text[];index" form:"platform" binding:"omitempty,dive,oneof=android ios web"`
}
