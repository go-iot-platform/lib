package iot

import model "github.com/go-iot-platform/rbac/model"

// PropertyTemplate
type PropertyTemplate struct {
	tableName struct{} `pg:",discard_unknown_columns"`
	model.Base
	Name        string `pg:"type:varchar(150),notnull, unique:name_parent_id" json:"name,omitempty"`
	Description string `pg:"type:varchar(500)" json:"description,omitempty"`

	DataType DataType `pg:"type:smallint" json:"dataType"`

	DefaultValue string `pg:"type:text" json:"defaultValue,omitempty"`

	IsPersistent bool `pg:",default:'false' " json:"isPersistent"`
	IsReadOnly   bool `pg:",default:'false' " json:"isReadOnly"`
	IsLog        bool `pg:",default:'false' " json:"isLogged"`

	ParentId int            `pg:"type:bigint,notnull, unique:name_parent_id" json:"-"`
	Parent   *ThingTemplate `json:"parent,omitempty"`

	// Alerts []AlertTemplate `pg:"fk:parent_id" json:"alerts,omitempty"`
}
