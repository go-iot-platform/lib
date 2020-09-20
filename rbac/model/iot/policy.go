package iot

import model "github.com/go-iot-platform/rbac/model"

// Certificate
type Policy struct {
	tableName struct{} `pg:",discard_unknown_columns"`
	model.Base
	Name        string `pg:"type:varchar(150),notnull,unique:name_customer_number" json:"name"`
	Description string `pg:"type:varchar(500)" json:"description"`
	Effect      bool   `pg:", default:'true'" json:"effect"`
	Action      string `pg:", notnull" json:"action"`

	Resources []string `pg:", notnull" json:"resources"`

	ProjectId int      `pg:"type:bigint, null" json:"-"`
	Project   *Project `json:"project,omitempty"`

	OwnerId        string `pg:"type:uuid,notnull" json:"-"`
	CustomerNumber string `pg:"type:varchar(50),notnull,unique:name_customer_number" json:"-"`

	Certificates []Certificate `pg:"many2many:policy_certificates,joinFK:certificate_id" json:"certificates,omitempty"` // many to many relation
}
type PolicyCertificate struct {
	PolicyId      int          `pg:",pk"`
	Policy        *Policy      `json:"policy,omitempty"`
	CertificateId int          `pg:",pk"` // pk tag is used to mark field as primary key
	Certificate   *Certificate `json:"certificate,omitempty"`
}
