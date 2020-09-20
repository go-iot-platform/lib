package names

import (
	"testing"
)

func TestContainer(t *testing.T) {
	fullname := "Nguyen Cat Pham"

	firstName := testGetFirstName(fullname, t)
	lastName := testGetLastName(fullname, t)
	testGetFullName(firstName, lastName, t)
}
func testGetFirstName(name string, t *testing.T) string {
	result := GetFirstName(name)
	if result == "" {
		t.Fatal("get first name error")
	}
	return result
}
func testGetLastName(name string, t *testing.T) string {
	result := GetLastName(name)
	if result == "" {
		t.Fatal("get last name error")
	}
	return result
}
func testGetFullName(firstname string, lastName string, t *testing.T) {
	result := GetFullName(firstname, lastName)
	if result == "" {
		t.Fatal("get full name error")
	}
}
