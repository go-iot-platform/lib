package payment

import (
	"time"
)

type PromotionType uint8

const (
	SalePromotion = iota + 1
)

var PromotionTypes = []PromotionType{
	SalePromotion,
}

func (p PromotionType) String() string {
	switch p {
	case SalePromotion:
		return "SalePromotion"
	default:
		return ""
	}
}

//Promotion represent the promotions table
type Promotion struct {
	UUID        string                 `pg:",pk" json:"uuid,omitempty"`
	CreatedAt   time.Time              `pg:"created_at" json:"created_at,omitempty"`
	UpdatedAt   time.Time              `pg:"updated_at" json:"updated_at,omitempty"`
	Title       string                 `json:"title,omitempty"`
	Description string                 `json:"description,omitempty"`
	Type        PromotionType          `json:"type,omitempty"`
	Value       int16                  `json:"value,omitempty"`
	StartDate   time.Time              `json:"start_date,omitempty"`
	EndDate     time.Time              `json:"end_date,omitempty"`
	IsActive    bool                   `pg:",use_zero" json:"is_active,omitempty"`
	Code        string                 `json:"code,omitempty"`
	Meta        map[string]interface{} `json:"meta,omitempty"`
}
