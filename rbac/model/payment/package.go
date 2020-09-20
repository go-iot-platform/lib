package payment

import (
	"time"

	"github.com/go-iot-platform/rbac/model"
)

type PackageType uint8

const (
	Time = iota + 1
	Quota
	Interval
)

var PackageTypes = []PackageType{
	Time,
	Quota,
	Interval,
}

func (s PackageType) String() string {
	switch s {
	case Time:
		return "Time"
	case Quota:
		return "Quota"
	case Interval:
		return "Interval"
	default:
		return ""
	}
}

// Package represents package table
type Package struct {
	model.Base
	UUID             string                 `pg:"type:uuid" json:"uuid" validate:"omitempty,uuid4"`
	Name             string                 `json:"name" validate:"required,max=50"`
	Description      string                 `json:"description" validate:"omitempty,max=255"`
	Price            uint64                 `json:"price" validate:"required,numeric"`
	StartDate        time.Time              `json:"start_date"`
	EndDate          time.Time              `json:"end_date"`
	Active           bool                   `pg:",use_zero" json:"active"`
	Type             PackageType            `json:"type" validate:"required,min=1,max=3"`
	TypeName         string                 `pg:"-" json:"type_name,omitempty"`
	BusinessName     string                 `json:"business_name" validate:"required,max=50"`
	Image            string                 `json:"image" validate:"omitempty,uuid4"`
	Publish          bool                   `pg:",use_zero" json:"publish"`
	ShortDescription string                 `json:"short_description" validate:"omitempty,max=50"`
	Meta             map[string]interface{} `json:"meta" validate:"omitempty"`
	IntervalTime     uint16                 `json:"interval_time" validate:"omitempty,gt=0"` // Month
	Quota            uint16                 `json:"quota" validate:"omitempty,gte=0"`
	RegionID         int                    `json:"region,omitempty" validate:"required"`
	CustomerNumber   string                 `json:"-"`
	Region           *Region                `json:"region_detail"`
	HaveTrialPackage bool                   `pg:",use_zero" json:"have_trial_package"`
	TrialDuration    string                 `json:"trial_duration"`                       //Could be a specific of time like : 1-m, 2-m (month) or quota
	Duration         uint16                 `json:"duration" validate:"omitempty,gte=24"` //Hour
	ServiceName      string                 `pg:"-" json:"service_name" validate:"required"`
	Currency         string                 `pg:"-" json:"currency"`
	PromotionID      string                 `json:"promotion_id" validate:"omitempty,uuid4"`
	Promotion        *Promotion             `json:"promotion_detail"`
	IsAutoCharge     bool                   `pg:"-" json:"is_auto_charge,omitempty"`
}

type ListPackageWeb struct {
	UUID         string      `json:"uuid,omiempty"`
	Name         string      `json:"name,omitempty"`
	Price        uint64      `json:"price,omitempty"`
	Image        string      `json:"image,omitempty"`
	Type         PackageType `json:"package_type,omitempty"`
	TypeName     string      `json:"type_name,omitempty"`
	Active       bool        `json:"active,omitempty"`
	Currency     string      `json:"currency,omitempty"`
	ServiceName  string      `json:"service_name,omitempty"`
	BusinessName string      `json:"business_name,omitempty"`
	Promotion    *Promotion  `json:"promotion,omitempty"`
	Bought       bool        `json:"bought"`
}

//ListPackageRequest represents API list package request
type ListPackageRequest struct {
	Filter string `json:"filter" `
	Offset int    `json:"offset" `
	Limit  int    `json:"limit" `
	Search string `json:"search" `
}
