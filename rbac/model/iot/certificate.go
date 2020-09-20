package iot

import (
	model "github.com/go-iot-platform/rbac/model"
)

// Certificate
type Certificate struct {
	model.Base
	CertId      string `pg:"type:varchar(500), notnull, unique" json:"certId"`
	PublicKey   string `pg:"type:varchar(5000), notnull, unique" json:"publicKey"`
	PrivateKey  string `pg:"type:varchar(5000), notnull, unique" json:"privateKey"`
	Certificate string `pg:"type:varchar(5000), notnull, unique" json:"certificate"`
	RootCA      string `pg:"-" json:"ca"`

	ProjectId int      `pg:"type:bigint, null" json:"-"`
	Project   *Project `json:"project,omitempty"`

	OwnerId        string `pg:"type:uuid,notnull" json:"-"`
	CustomerNumber string `pg:"type:varchar(50),notnull" json:"-"`

	Things []Thing `pg:"many2many:certificate_things,joinFK:thing_id" json:"things,omitempty"` // many to many relation

	Policies []Policy `pg:"many2many:policy_certificates,joinFK:policy_id" json:"policies,omitempty"` // many to many relation
}
type CertificateThing struct {
	CertificateId int          `pg:",pk"` // pk tag is used to mark field as primary key
	Certificate   *Certificate `json:"certificate,omitempty"`
	ThingId       int          `pg:",pk"`
	Thing         *Thing       `json:"thing,omitempty"`
}
