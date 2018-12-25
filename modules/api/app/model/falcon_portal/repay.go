package falcon_portal

import (
	"time"
)

type Repay struct {
	ID               int64      `gorm:"column:id" json:"id"`
	Number           string     `gorm:"column:number" json:"number"`
	CurrentPrincipal float32    `gorm:"column:current_principal" json:"current_principal"`
	CurrentInterest  float32    `gorm:"column:current_interest" json:"current_interest"`
	ShouldDate       time.Time  `gorm:"column:should_date" json:"should_date"`
	RealPrincipal    float32    `gorm:"column:real_principal" json:"real_principal"`
	RealInterest     float32    `gorm:"column:real_interest" json:"real_interest"`
	RealDate         *time.Time `gorm:"column:real_date" json:"real_date"`
	UserID           int64      `gorm:"column:user_id" json:"user_id"`
}

func (this Repay) TableName() string {
	return "repay"
}
