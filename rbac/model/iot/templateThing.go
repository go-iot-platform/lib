package iot

import model "github.com/go-iot-platform/rbac/model"

// "fmt"

// AccessRole represents access role type
type TemplateType int8
type ThingType int8

const (
	// template defined by system
	System_Template = iota + 1

	// template define by user
	User_Defined_Template
)

const (
	Device_Thing = iota + 1
	Remote_Thing
	Gateway_Thing
	Generic_Thing
	Zone_Thing
	Schedule_Thing
	Scence_Thing
)

func (s TemplateType) String() string {
	switch s {
	case System_Template:
		return "System"
	case User_Defined_Template:
		return "User_Defined"
	default:
		return "Unknown"
	}
}

func (s ThingType) String() string {
	switch s {
	case Device_Thing:
		return "Device"
	case Remote_Thing:
		return "Remote"
	case Gateway_Thing:
		return "Gateway"
	case Generic_Thing:
		return "Generic"
	case Zone_Thing:
		return "Zone"
	case Schedule_Thing:
		return "Schedule"
	case Scence_Thing:
		return "Scence"
	default:
		return "Unknown"
	}
}
func GetThingTypeList() []Configure {
	var result = make([]Configure, 7)
	for i := 1; i < 8; i++ {
		id := i
		thingType := ThingType(i)
		result[i-1] = Configure{
			&id,
			thingType.String(),
		}
	}

	return result
}

// ThingTemplate
type ThingTemplate struct {
	tableName struct{} `pg:",discard_unknown_columns"`
	model.Base
	Name        string `pg:"type:varchar(150),notnull,unique:name_customer_number" json:"name,omitempty"`
	Description string `pg:"type:varchar(500)" json:"description,omitempty"`

	TemplateType TemplateType `json:"templateType"`

	Type ThingType `json:"type"`

	ProjectId int      `pg:"type:bigint, null" json:"-"`
	Project   *Project `json:"project,omitempty"`

	ImageId  string `pg:"type:uuid, null" json:"imageId"`
	ImageUrl string `pg:"-" json:"imageUrl"`

	OwnerId        string `pg:"type:uuid,null" json:"-"`
	CustomerNumber string `pg:"type:varchar(50),notnull,unique:name_customer_number" json:"-"`

	Parent   *ThingTemplate  `json:"parent,omitempty"`
	ParentId *int            `json:"parentId,omitempty"`
	Items    []ThingTemplate `pg:"fk:parent_id" json:"subTemplates,omitempty"`

	Properties []PropertyTemplate `pg:"fk:parent_id" json:"properties,omitempty"`
}
