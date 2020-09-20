package version

import (
	"errors"
	"strconv"
	"strings"
)

//MaxNumber Max number
const MaxNumber = 9

// SeperateChar Seperate Char
const SeperateChar = "."

// DefaultVersion default version
const DefaultVersion = "1.0.0"

// UpdateVersion updates version of a service when update
func UpdateVersion(version string) (string, error) {
	if version == "" {
		return DefaultVersion, nil
	}
	versions := strings.Split(version, SeperateChar)

	for i := len(versions) - 1; i >= 0; i-- {
		number, err := strconv.Atoi(versions[i])
		if err != nil {
			return "", errors.New("version:invalid")
		}
		if number < MaxNumber {
			number++
			versions[i] = strconv.Itoa(number)
			return strings.Join(versions, SeperateChar), nil
		}
		if i == 0 {
			number++
			versions[i] = strconv.Itoa(number)
			break
		}
		number = 0
		versions[i] = strconv.Itoa(number)
	}
	return strings.Join(versions, SeperateChar), nil
}
