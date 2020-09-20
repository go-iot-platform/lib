package model

// Storage represents company model
type Storage struct {
	Uuid           string      `json:"uuid"`
	Url            string      `json:"url"`
	Path           string      `json:"path"`
	Type           StorageType `json:"type"`
	OwnerID        string      `json:"-"`
	CustomerNumber string      `json:"-"`
	IsUsed         bool        `json:"-"`
}
