package authen

type GroupCache struct {
	ID         int    `json:"id"`
	Name       string ` json:"name"`
	Uuid       string ` json:"uuid"`
	CustomerID int    ` json:"-"`

	Users    []string `json:"users,omitempty"` // many to many relation
	Policies []Policy `json:"policies,omitempty"`
}
