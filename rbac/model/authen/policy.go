package authen

type Policy struct {
	ID            int            `json:"id"`
	PolicyId      string         `json:"policyId,omitempty"`
	Name          string         `json:"name,omitempty"`
	Description   string         `json:"description,omitempty"`
	Type          string         `json:"type,omitempty"`
	Version       string         `json:"version,omitempty"`
	ResourceTypes []ResourceType `json:"resourceTypes,omitempty"`
}

type ResourceType struct {
	Name         string   `json:"name,omitempty"`
	Effect       string   `json:"effect,omitempty"`
	ResourceType string   `json:"resourceType,omitempty"`
	Actions      []string `json:"actions,omitempty"`
	Resources    []string `json:"resources,omitempty"`
}
