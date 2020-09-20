package model

// Customer
type Customer struct {
	ID            int            `json:"id"`
	AccNumber   string             `sql:"type:bigint, notnull, unique, default: public.account_generator()" json:"accountNumber"`
	FirstName   string             `sql:"type:varchar(50)" json:"firstName,omitempty"`
	LastName    string             `sql:"type:varchar(50)" json:"lastName,omitempty"`
	Email       string             `sql:"type:varchar(150),notnull,unique:email_parent_id" json:"email"`
	Mobile      string             `sql:"type:varchar(15)" json:"mobile,omitempty"`
	Address1    string             `sql:"type:varchar(150)" json:"address1,omitempty"`
	Address2    string             `sql:"type:varchar(150)" json:"address2,omitempty"`
	Domain      string             `sql:"type:varchar(100), null, unique" json:"domain,omitempty"`
	Alias       string             `sql:"type:varchar(50), null" json:"alias,omitempty"`
	Confirmed   bool               `sql:", default:'false'" json:"confirmed"`
	ImageID     string             `sql:"type:uuid, null" json:"imageId"`
	ImageUrl    string             `sql:"-" json:"imageUrl"`
	Active      bool               `sql:", default:'true'" json:"active"`
	Clients     []Customer         `pg:"fk:parent_id" json:"clients,omitempty"`
	ParentID    *int               `sql:",unique:email_parent_id" json:"parentId,omitempty"`
	Parent      *Customer          `json:"parent,omitempty"`
	Permissions []ClientPermission `json:"permissions,omitempty"`
	Setting     *CustomerSetting   `json:"setting,omitempty"`
	SettingID   *int               `json:"settingId,omitempty"`
	Region      int                `json:"region,omitempty"`
}

// CustomerDB represents customer related database interface (repository)
type CustomerSetting struct {
	IsAllowRegister bool             `json:"isAllowRegister"`
	NeedConfirm     bool             `json:"needConfirm"`
	SMTP            *SMTPConfig      `json:"smtp"`
	Payment         *PaymentConfig   `json:"payment"`
	Other           *OtherConfig     `json:"other"`
	ParentID        *int             `json:"parentId,omitempty"`
	Parent          *CustomerSetting `json:"parent,omitempty"`
}
type SMTPConfig struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	Account  string `json:"account"`
}
type PaymentConfig struct {
}
type OtherConfig struct {
	Value string `json:"value"`
}

// CustomerDB represents customer related database interface (repository)
type ClientPermission struct {
	CustomerID int       `sql:"type:bigint,notnull,unique:customer_id_policy_id" json:"-"`
	Customer   *Customer `json:"customer,omitempty"`
	PolicyID   string    `sql:"type:uuid,notnull,unique:customer_id_policy_id" json:"-"`
}
