package model

// AuthUser represents data stored in JWT token for user
type AuthUser struct {
	ID             int        `json:"id,omitempty"`
	CustomerID     int        `json:"customerID,omitempty"`
	CustomerNumber string     `json:"customerNumber,omitempty"`
	Clients        []string   `json:"clients,omitempty"`
	Username       string     `json:"username,omitempty"`
	UserUuid       string     `json:"userUuid,omitempty"`
	Role           AccessRole `json:"-"`
}

type AuthToken struct {
	Token        string `json:"token"`
	Expires      string `json:"expires"`
	RefreshToken string `json:"refresh_token"`
}

// RefreshToken holds authentication token details
type RefreshToken struct {
	Token   string `json:"token"`
	Expires string `json:"expires"`
}

type ResetToken struct {
	ResetToken        string `json:"reset_token"`
	ResetTokenExpires string `json:"reset_token_expires"`
}
