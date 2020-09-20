package policy

// SystemPolicy sytem policies
type SystemPolicy struct {
	Name          string         `yaml:"name,omitempty"`
	Description   string         `yaml:"description,omitempty"`
	ResourceTypes []ResourceType `yaml:"resource_types,omitempty"`
}
