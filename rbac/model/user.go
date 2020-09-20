package model

// User represents user domain model
type User struct {
	ID            int            `json:"id"`
	Uuid     string `json:"uuid"`
	Username string `json:"username"`
	Active   bool   `json:"active"`
	Token    string `json:"-"`

	Role       *Role `json:"role,omitempty"`
	CustomerID int   `json:"-"`

	Groups     []GroupCache `json:"groups,omitempty"` // many to many relation
	ResetToken string       `json:"-"`
	Policies   []Policy     `json:"policies,omitempty"`
}
