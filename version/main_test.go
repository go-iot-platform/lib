package version

import (
	"testing"
)

func TestContainer(t *testing.T) {
	version := "1.1.9.9.8.9"

	testUpdateVersion(version, t)
}
func testUpdateVersion(version string, t *testing.T) string {
	result, err := UpdateVersion(version)
	if err != nil || result == "" {
		t.Fatal("update version error")
	}
	return result
}
