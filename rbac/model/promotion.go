package model

import (
	"time"
)

//Promotion represent the promotions table
type Promotion struct {
	UUID        string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Title       string
	Description string
	Type        uint8
	Value       int16
	StartDate   time.Time
	EndDate     time.Time
	IsActive    bool
	Code        string
	Meta        map[string]interface{}
}
