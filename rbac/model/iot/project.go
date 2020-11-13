package iot

import (
	model "github.com/go-iot-platform/lib/rbac/model"
)

// User represents user domain model
type Project struct {
	tableName struct{} `pg:",discard_unknown_columns"`
	model.Base
	// DisplayName    string `pg:"type:varchar(150),  json:"displayName,omitempty"`
	Name           string `pg:"type:varchar(150),notnull,unique:name_customer_number" json:"name"`
	Description    string `pg:"type:varchar(500)" json:"description"`
	OwnerId        string `pg:"type:uuid,notnull" json:"-"`
	CustomerNumber string `pg:"type:varchar(50),notnull,unique:name_customer_number" json:"-"`
	ImageId        string `pg:"type:uuid, null" json:"imageId"`
	ImageUrl       string `pg:"-" json:"imageUrl"`
}
