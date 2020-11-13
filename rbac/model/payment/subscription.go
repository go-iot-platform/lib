package payment

import (
	"time"

	"github.com/go-iot-platform/lib/rbac/model"
)

type SubscriptionStatus uint8

const (
	Pending = iota + 1 //! If modify the config below please update the subscription proto file, status field
	Active
	OnHold
	Cancelled
	Expired
	PaymentFailure
)

func (s SubscriptionStatus) String() string {
	switch s {
	case Pending:
		return "Pending" //wait for transaction
	case Active:
		return "Active" //transaction succeeded
	case OnHold:
		return "OnHold" //customer's account have problem
	case Cancelled:
		return "Cancelled" //support in case customer wants to cancel package
	case Expired:
		return "Expired" // customer's package expired
	case PaymentFailure:
		return "PaymentFailure" // transaction failed
	default:
		return ""
	}
}

// Subscription represents package table
type Subscription struct {
	model.Base
	UUID             string                 `pg:"type:uuid" json:"uuid" validate:"omitempty,uuid4"` // Tu generate
	PackageUUID      string                 `json:"package_uuid" validate:"omitempty,uuid4"`        // Truyen
	StartDate        time.Time              `json:"start_date" validate:"gte"`                      // Now
	EndDate          time.Time              `json:"end_date" validate:"gtfield=StartDate"`          // Now + package's duration (day)
	Type             PackageType            `json:"type" validate:"required,min=1,max=3"`           //
	Meta             map[string]interface{} `json:"meta" validate:"omitempty"`                      //
	Duration         uint16                 `json:"duration" validate:"omitempty,gte=0"`            // month
	IntervalTime     uint16                 `json:"interval_time" validate:"omitempty,gte=0"`       // month
	Quota            uint16                 `json:"quota" validate:"omitempty,gte=0"`               // uint16
	Price            uint64                 `json:"price" validate:"required,numeric"`              // uint
	CustomerNumber   string                 `json:"customer_number" validate:"required,alphanum"`   // Truyen
	Service          string                 `json:"service" validate:"required,alpha"`              // Truyen
	Status           SubscriptionStatus     `json:"status" validate:"required,min=1,max=5"`
	HaveTrialPackage bool                   `pg:",use_zero" json:"have_trial_package" validate:"oneof=true false"`
	TrialDuration    string                 `json:"trial_duration"` // hour or quota
	ChargeDate       time.Time              `json:"charge_date"`
	IsManualCharge   bool                   `pg:"-" json:"is_manual_charge"`
	IsAutoCharge     bool                   `pg:"-" json:"is_auto_charge"`
	// CardToken        string                 `json:"card_token,omitempty"`
	// CardCustomerID   string                 `json:"card_customer_id,omitempty"`
}

type SubScriptionQuery struct {
	UUID           string
	CustomerNumber string
	Service        string
	PackageUUID    string
}

type SubscriptionResponse struct {
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	PackageUUID string    `json:"package_uuid"`
	Price       uint64    `json:"price"`
	Service     string    `json:"service"`
	Status      uint32    `json:"status"`
	StatusName  string    `json:"status_name"`
}

type SubscriptionUpdateData struct {
	Status     uint32    `json:"status,omitempty"`
	Duration   uint32    `json:"duration,omitempty"`
	ChargeDate time.Time `json:"charge_date,omitempty"`
}
