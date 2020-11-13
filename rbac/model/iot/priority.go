package iot

import model "github.com/go-iot-platform/lib/rbac/model"

// Priority
type Priority struct {
	tableName struct{} `pg:",discard_unknown_columns"`
	model.Base
	Name        string `pg:"type:varchar(150),notnull,unique:name_customer_number" json:"name"`
	Description string `pg:"type:varchar(500)" json:"description"`
	Color       string `pg:"type:varchar(100)" json:"description"`

	ProjectId int      `pg:"type:bigint, null" json:"projectId,omitempty"`
	Project   *Project `json:"project,omitempty"`

	OwnerId        string `pg:"type:uuid,notnull" json:"-"`
	CustomerNumber string `pg:"type:varchar(50),notnull,unique:name_customer_number" json:"-"`
}
