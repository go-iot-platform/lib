package payment

import (
	"time"
)

//PaymentSetting represent payment setting table
type PaymentSetting struct {
	CustomerNumber string    `pg:",pk" json:"customer_number"`
	CreatedAt      time.Time `pg:"created_at" json:"created_at,omitempty"`
	UpdatedAt      time.Time `pg:"updated_at" json:"updated_at,omitempty"`
	IsAutoCharge   bool      `pg:",use_zero" json:"is_auto_charge"`
	CardToken      string    `json:"card_token"`
}

//! Need validate
type PaymentSettingReq struct {
	IsAutoCharge       bool   `json:"is_auto_charge"`
	PaymentMethodNonce string `json:"payment_method_nonce,omitempty"`
	PaymentGateway     string `json:"payment_gateway,omitempty"`
}

type PaymentSettingRes struct {
	IsAutoCharge   bool   `json:"is_auto_charge"`
	CardType       string `json:"card_type,omitempty"`
	ExpirationDate string `json:"expiration_date,omitempty"`
	CardNumber     string `json:"card_number,omitempty"`
	CardImage      string `json:"card_image,omitempty"`
	Email          string `json:"email,omitempty"`
}
