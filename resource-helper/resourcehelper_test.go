package resourcehelper

import (
	"testing"
)

// TestContainer test for resource helper
func TestContainer(t *testing.T) {
	resources := []string{"*"}

	result, err := GetResourceIds(resources)
	if err != nil || result == nil {
		t.Fatal("Get resources id fail!")
	}
}
