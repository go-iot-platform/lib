package iot

import model "github.com/go-iot-platform/lib/rbac/model"

// PropertyAlert
type PropertyAlert struct {
	model.Base
	/**-------------------------------------------------------------**/
	//null when inherit from template
	//not null when uninherit from template
	Name        *string `pg:"type:varchar(150),unique:name_parent_id" json:"name,omitempty"`
	Description *string `pg:"type:varchar(500)" json:"description,omitempty"`
	AlertModule *string `pg:"type:varchar(50)" json:"alertModule,omitempty"`
	Value       *string `pg:"type:varchar(50)" json:"value,omitempty"`

	PriorityId *int      `pg:"type:bigint, null" json:"priorityId,omitempty"`
	Priority   *Priority `json:"priority,omitempty"`

	/**-------------------------------------------------------------**/
	Enabled bool `pg:", default:'true'" json:"enabled"`

	ParentId int            `pg:"type:bigint, notnull,unique:name_parent_id" json:"parentId,omitempty"`
	Parent   *ThingProperty `json:"parent,omitempty"`

	TemplateId int `pg:"type:bigint, null" json:"templateId,omitempty"`
	// Template   *AlertTemplate `json:"template,omitempty"`

	// Triggers []Trigger `json:"triggers,omitempty"`
}
