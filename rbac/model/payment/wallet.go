package payment

import (
	"github.com/go-iot-platform/lib/rbac/model"
)

type Wallet struct {
	model.Base
	CustomerNumber string `pg:"customer_number,type:varchar(50),unique,notnull" json:"customer_number"`
	Currency       string `pg:"currency,type:varchar(5),notnull" json:"currency"`
	Coin           int64  `pg:"coin,type:bigint,notnull" json:"coin"`
}

type GetWalletInfo struct {
	CustomerNumber string `json:"customer_number" validate:"max=50"`
	Currency       string `json:"currency" validate:"max=3"`
}
