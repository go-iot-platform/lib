package params

import (
	"errors"
	"strconv"
	"strings"
)

func ConvertToInt(str string, sep string) ([]int, error) {
	temps := strings.Split(str, sep)
	var ids []int
	for i := 0; i < len(temps); i++ {
		id, err := strconv.Atoi(temps[i])
		if err != nil {
			continue
		}
		if containsInt(ids, id) {
			continue
		}
		ids = append(ids, id)
	}
	if len(ids) != len(temps) {
		return nil, errors.New("invalid:params")
	}
	return ids, nil
}
func ConvertToString(str string, sep string) ([]string, error) {
	temps := strings.Split(str, sep)
	var ids []string
	for i := 0; i < len(temps); i++ {
		if contains(ids, temps[i]) {
			continue
		}
		ids = append(ids, temps[i])
	}

	if len(ids) != len(temps) {
		return nil, errors.New("invalid:params")
	}
	return ids, nil
}
func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
func containsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ParseMicroName parse micro name to support call micro with namespace
func ParseMicroName(microName string, namespace string) (string, string) {

	f := func(c rune) bool {
		return c == 46 // .
	}
	datas := strings.FieldsFunc(microName, f)
	if len(datas) < 2 {
		return microName, namespace
	}
	if namespace != "" {
		return strings.Join(datas, "/"), namespace

	}
	return strings.Join(datas[0:len(datas)-1], "/"), datas[len(datas)-1]
}
