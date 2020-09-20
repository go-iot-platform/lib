package iot

import model "github.com/go-iot-platform/rbac/model"

// Datatype of a property of thing
type DataType int8

const (
	Number = iota + 1
	String
	Json
	Image
	Boolean
	ThingObject
	XML
	DateTime
	Location
	Schedule
	HTML
)

func (s DataType) String() string {
	switch s {
	case Number:
		return "Number"
	case String:
		return "String"
	case Json:
		return "Json"
	case Image:
		return "Image"
	case Boolean:
		return "Boolean"
	case ThingObject:
		return "ThingObject"
	case XML:
		return "XML"
	case DateTime:
		return "DateTime"
	case Location:
		return "Location"
	case Schedule:
		return "Schedule"
	case HTML:
		return "HTML"
	default:
		return "Unknown"
	}
}
func GetDataTypeList() []Configure {
	var result = make([]Configure, 11)
	for i := 1; i < 12; i++ {
		id := i
		dataType := DataType(i)
		result[i-1] = Configure{
			&id,
			dataType.String(),
		}
	}

	return result
}

// ThingProperty
type ThingProperty struct {
	tableName struct{} `pg:",discard_unknown_columns"`
	model.Base
	//---
	//--- null when inherit from template
	//--- not null when uninherit from template
	//--- Name, Description, DataType, IsPersistent,IsReadOnly, IsLog is able to override.
	//+++ If not override get value from template
	Name        *string `pg:"type:varchar(150), unique:name_parent_id" json:"name,omitempty"`
	Description *string `pg:"type:varchar(500)" json:"description,omitempty"`

	DataType *DataType `pg:"type:smallint" json:"dataType",omitempty`
	Value    *string   `pg:"type:text" json:"value,omitempty"`

	IsPersistent *bool `json:"isPersistent,omitempty"`
	IsReadOnly   *bool `json:"isReadOnly,omitempty"`
	IsLog        *bool `json:"isLogged,omitempty"`
	//---
	ParentId int `pg:"type:bigint,notnull,unique:name_parent_id" json:"-"`

	Parent *Thing `json:"parent,omitempty"`

	TemplateId int `pg:"type:bigint" json:"templateId,omitempty"`

	Template *PropertyTemplate `json:"template,omitempty"`

	Alerts []PropertyAlert `pg:"fk:parent_id" json:"alerts,omitempty"`

	ProperyDevices []ThingProperty `pg:"fk:control_id" json:"properyDevices,omitempty"`

	ControlId *int `json:"controlId,omitempty"`

	Control *ThingProperty `json:"control,omitempty"`
}

type ThingPropertyUpdateValue struct {
	Name      string                       `json:"name,omitempty"`
	ThingName string                       `json:"thingName,omitempty"`
	Body      ThingPropertyUpdateValueBody `json:"value,omitempty"`
}

type ThingPropertyUpdateValueBody struct {
	Value string `json:"value,omitempty"`
}
