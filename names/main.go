package names

import (
	"fmt"
	"strings"
)

// GetFirstName get first name
func GetFirstName(name string) string {
	if name == "" {
		return ""
	}
	wordArr := strings.Fields(name)
	return wordArr[0]
}

// GetLastName get last name
func GetLastName(name string) string {
	if name == "" {
		return ""
	}
	wordArr := strings.Fields(name)
	if len(wordArr) == 1 {
		return wordArr[0]
	}
	return strings.Join(wordArr[1:len(wordArr)], " ")

}

// GetFullName get full name
func GetFullName(firstName string, lastName string) string {
	if firstName == "" && lastName == "" {
		return ""
	}
	if firstName == lastName {
		return firstName
	}
	return fmt.Sprintf("%s %s", firstName, lastName)
}
