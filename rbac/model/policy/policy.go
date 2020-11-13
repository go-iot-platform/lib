package policy

import (
	"github.com/go-iot-platform/lib/rbac/model"
)

type PolicyType int8

const (
	List = iota + 1
	Read
	Create
	Update
	Delete
	System
	CustomerDefine
	PreSystem
	// ResourceType
)

var PolicyTypes = []PolicyType{
	List,
	Read,
	Create,
	Update,
	Delete,
	System,
	CustomerDefine,
	PreSystem,
	// ResourceType,
}

func (s PolicyType) String() string {
	switch s {
	case List:
		return "pol:listPolicy"
	case Read:
		return "pol:readPolicy"
	case Create:
		return "pol:addPolicy"
	case Update:
		return "pol:editPolicy"
	case Delete:
		return "pol:deletePolicy"
	case System:
		return "system"
	case CustomerDefine:
		return "customer-define"
	case PreSystem:
		return "trial"
	// case ResourceType:
	// 	return "resourceType"
	default:
		return ""
	}
}

// Policy represents company model
type Policy struct {
	model.Base
	PolicyID       string         `pg:"policy_id,unique" json:"policy_id,omitempty"`
	Name           string         `pg:"name,type:varchar(255)" json:"name,omitempty"`
	Description    string         `pg:"description,type:varchar(255)" json:"description,omitempty"`
	CustomerNumber string         `pg:"customer_number,type:varchar(32)" json:"customer_number,omitempty"`
	Version        string         `pg:"version,type:varchar(255)" json:"version,omitempty"`
	Type           PolicyType     `pg:"type" json:"type"`
	ResourceTypes  []ResourceType `pg:"resource_types" json:"resourceTypes,omitempty"`
	IsActive       bool           `pg:"is_active"`
	User           *User
	Group          *Group
}

type ResourceType struct {
	Name         string   `json:"name,omitempty" yaml:"name,omitempty"`
	Effect       string   `json:"effect,omitempty" yaml:"effect,omitempty"`
	Actions      []string `json:"actions,omitempty" yaml:"actions,omitempty"`
	ResourceType string   `json:"resourceType,omitempty"`
	Resources    []string `json:"resources,omitempty" yaml:"resources,omitempty"`
}
type UserInterface struct {
	ID             int
	CustomerID     int
	CustomerNumber string
	Username       string
	UserUUID       string
}
type AuthorizeResponse struct {
	Resources []string
	User      UserInterface
	Clients   []UserInterface
}

type ListPoliciesQuery struct {
	Limit     uint64
	Skip      uint64
	Keyword   string
	Keysort   string
	IsAsc     bool
	IsActive  string
	Type      string
	Resources []string
	User      UserInterface
}

type GetPolicyByResouceQuery struct {
	ResourceTypeName string
	ResourceOrn      []string
	Effect           string
	AuthResp         *AuthorizeResponse
}

type GetPolicyByGroupQuery struct {
	GroupIDs []string
}

type SystemPoliciesQuery struct {
	Limit    uint64
	Skip     uint64
	Keyword  string
	Keysort  string
	IsAsc    bool
	IsActive bool
}
