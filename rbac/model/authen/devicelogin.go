package authen

import (
	"time"
)

// Location represents company location model
type DeviceLogin struct {
	ID           int       `json:"id"`
	User         *User     `json:"user,omitempty"`
	UserID       int       `sql:"type:bigint,null" json:"-"`
	IP           string    `sql:"type:varchar(50),null" json:"-"`
	Location     string    `json:"location"`
	Lat          float32   `json:"lat"`
	Long         float32   `json:"long"`
	LastLogin    time.Time `json:"last_login"`
	Active       bool      `json:"active"`
	IsMainDevice bool      `json:"is_main_device"`
}
