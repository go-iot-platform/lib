package rbac

import (
	"testing"
)

// TestContainer test
func TestContainer(t *testing.T) {

	result := Foo()
	if result != "Bar" {
		t.Fatal("reject")
	}
}
