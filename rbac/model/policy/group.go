package policy

import (
	"github.com/go-iot-platform/lib/rbac/model"
)

// Group represents service domain model
type Group struct {
	model.Base
	PolicyID int64  `pg:"policy_id" json:"policy_id,omitempty"`
	GroupID  string `pg:"group_id" json:"group_id,omitempty"`

	Users  []User `pg:"-" json:"users,omitempty"`
	Policy *Policy
}

type ListGroup struct {
	GroupID  string    `json:"group_id"`
	Policies []*Policy `json:"policies"`
}
