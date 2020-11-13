package policy

import (
	"github.com/go-iot-platform/lib/rbac/model"
)

// User represents service domain model
type User struct {
	model.Base
	PolicyID int64  `pg:"policy_id" json:"policyId,omitempty"`
	UserID   string `pg:"user_id" json:"userId,omitempty"`
	UUID     string `pg:"-" json:"uuid,omitempty"`
	Policy   *Policy
}

type ListUser struct {
	UserUUID string    `json:"policyId"`
	Policies []*Policy `json:"policies"`
}
