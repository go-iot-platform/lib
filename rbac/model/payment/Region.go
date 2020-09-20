package payment

import (
	"github.com/go-iot-platform/rbac/model"
)

// Region represents region table.
type Region struct {
	model.Base
	Name        string `json:"name,omitempty" validate:"min=1,max=255,alpha"`
	Description string `json:"description,omitempty" validate:"min=1,max=255"`
	Country     string `json:"country,omitempty" validate:"min=2,max=3"`
	Currency    string `json:"currency,omitempty" validate:"len=3"`
}

//DeleteRegionRequest represent delete region request
type DeleteRegionRequest struct {
	ID int `json:"id" validate:"min=1"`
}
