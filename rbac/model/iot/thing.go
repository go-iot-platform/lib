package iot

import model "github.com/go-iot-platform/rbac/model"

// Thing
type Thing struct {
	tableName struct{} `pg:",discard_unknown_columns"`
	model.Base
	// DeletedAt   *time.Time `json:"deleted_at,omitempty" pg:",soft_delete"`
	Serial      string `pg:"type:varchar(50), notnull, unique, default: public.serial_generator()" json:"serial"`
	Name        string `pg:"type:varchar(150),notnull, unique:name_owner_id" json:"name,omitempty"`
	DisplayName string `pg:"type:varchar(150)" json:"displayName,omitempty"`
	Description string `pg:"type:varchar(500)" json:"description,omitempty"`

	// Type     ThingType `json:"type"`
	IsActive   bool `pg:", default:'false'" json:"isActive"`
	IsRegister bool `pg:", default:'false'" json:"isRegister"`

	ProjectId int      `pg:"type:bigint, null" json:"projectId,omitempty"`
	Project   *Project `json:"project,omitempty"`

	TemplateId int            `pg:"type:bigint, notnull" json:"templateId,omitempty"`
	Template   *ThingTemplate `json:"template,omitempty"`

	ImageId  string `pg:"type:uuid, null" json:"imageId"`
	ImageUrl string `pg:"-" json:"imageUrl"`

	OwnerId        string          `pg:"type:uuid,notnull,unique:name_owner_id" json:"-"`
	CustomerNumber string          `pg:"type:varchar(50),notnull" json:"customerNumber,omitempty"`
	Properties     []ThingProperty `pg:"fk:parent_id" json:"properties,omitempty"`

	Certificates []Certificate `pg:"many2many:certificate_things,joinFK:certificate_id" json:"certificates,omitempty"`

	Items    []Thing `pg:"fk:parent_id" json:"subThings,omitempty"`
	ParentId *int    `json:"parentId,omitempty"`
	Parent   *Thing  `json:"parent,omitempty"`
}
