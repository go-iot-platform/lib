package payment

import (
	"github.com/go-iot-platform/rbac/model"
)

type InvoiceType uint

const (
	TopupInvoice = iota + 1
	BuyService
)

func (s InvoiceType) String() string {
	switch s {
	case Topup:
		return "Topup"
	case BuyService:
		return "BuyService"
	default:
		return ""
	}
}

// Invoice represents package table
type Invoice struct {
	model.Base
	ServiceName          string      `json:"service_name" validate:"omitempty"`
	PackageBussinessName string      `json:"package_business_name" validate:"omitempty"`
	UnitPrice            uint64      `json:"unit_price" validate:"omitempty"`
	Quantity             uint16      `json:"quantity" validate:"omitempty"`
	Amount               uint64      `json:"amount" validate:"omitempty"`
	TaxRate              uint8       `json:"tax_rate" validate:"omitempty"`
	Total                uint64      `json:"total" validate:"omitempty"`
	InvoiceType          InvoiceType `json:"invoice_type" validate:"omitempty"`
}
