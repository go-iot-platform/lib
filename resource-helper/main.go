package resourcehelper

import (
	"errors"
	"fmt"
	"strings"
)

// ItemActionCache item action in cache
type ItemActionCache struct {
	Name         string `json:"name,omitempty"`
	ResourceType string `json:"resourceType,omitempty"`
	Type         string `json:"type,omitempty"`
	Service      string `json:"service,omitempty"`
}

// ResourceItem response data
type ResourceItem struct {
	Patrition      string
	Service        string
	Region         string
	CustomerNumber string
	ResourceType   string
	Resource       string
}

// ToString convert resource item to resource orn
func (r *ResourceItem) ToString() string {
	resourceFormatStr := `orn:%s:%s:%s:%s:%s/%s`
	return fmt.Sprintf(resourceFormatStr, r.Patrition, r.Service, r.Region, r.CustomerNumber, r.ResourceType, r.Resource)
}

// GetResourceIds get list resources id from list resource string
func GetResourceIds(resourceStr []string) ([]string, error) {
	var resourceIds []string

	for i := 0; i < len(resourceStr); i++ {
		r := GenResourceItem(resourceStr[i])
		if r == nil {
			continue
		}
		if Contains(resourceIds, r.Resource) {
			continue
		}
		resourceIds = append(resourceIds, r.Resource)
	}
	if resourceIds == nil || len(resourceIds) < 0 {
		return nil, errors.New("Resource Empty")

	}
	return resourceIds, nil
}

// GetResourceFormat get resource format
// name of policy           info.Service      info.ResourceType
// iot:all                     iot               policy
func GetResourceFormat(resourceOrn string, info *ItemActionCache, customerNumber string) *ResourceItem {
	var result *ResourceItem
	patrition := ""
	region := ""
	//                   [orn::iot::1969541697209631746:policy/iot]
	resourceFormatStr := `orn:%s:%s:%s:%s:%s/%s`
	//check formart of resourceOrn must follow `orn:%s:%s:%s:%s:%s/%s`
	if resourceOrn == "" {
		return nil
	}
	if resourceOrn == "*" {
		resourceOrn = fmt.Sprintf(resourceFormatStr, patrition, info.Service, region, customerNumber, info.ResourceType, resourceOrn)
	}
	if resourceOrn != "*" {
		s1 := strings.Split(resourceOrn, ":")
		if s1 == nil || len(s1) < 6 {
			resourceOrn = fmt.Sprintf(resourceFormatStr, patrition, info.Service, region, customerNumber, info.ResourceType, resourceOrn)
		}
	}
	//by pass check format if resourceOrn  "*"
	if resourceOrn != "" {
		result = GenResourceItem(resourceOrn)

	}
	return result
}

// ParseResource return values
//- bool : false if source, origin resource not match format or source is not match origin
//- source in ResourceItem struct
//- origin in ResourceItem struct
func ParseResource(resourceOrn string, originalResource string, info *ItemActionCache, customerNumber string) (bool, *ResourceItem, *ResourceItem) {
	source := GetResourceFormat(resourceOrn, info, customerNumber)
	if source == nil {
		return false, nil, nil
	}
	original := GetResourceFormat(originalResource, info, customerNumber)
	if original == nil {
		return false, source, original
	}
	result := checkResourceSubItem(source.Patrition, original.Patrition, true) &&
		checkResourceSubItem(source.Service, original.Service, true) &&
		checkResourceSubItem(source.Region, original.Region, true) &&
		checkResourceSubItem(source.CustomerNumber, original.CustomerNumber, true) &&
		checkResourceSubItem(source.ResourceType, original.ResourceType, true) &&
		checkResourceSubItem(source.Resource, original.Resource, false)
	// fmt.Printf("ParseResource original:%+v\n", original)
	// fmt.Printf("ParseResource source:%+v\n", source)
	return result, source, original

}

// GenResourceItem parser resource from resource orn
func GenResourceItem(resourceOrn string) *ResourceItem {
	var result *ResourceItem
	// resourceFormatStr := `orn:%s:%s:%s:%s:%s/%s`
	//check formart of resourceOrn must follow `orn:%s:%s:%s:%s:%s/%s`
	//by pass check format if resourceOrn  "*"
	if resourceOrn == "*" {
		result = &ResourceItem{
			Patrition:      "*",
			Service:        "*",
			Region:         "*",
			CustomerNumber: "*",
			ResourceType:   "*",
			Resource:       "*",
		}
		return result
	}
	//`orn:%s:%s:%s:%s:%s/%s`
	orn := strings.Split(resourceOrn, "/")
	if orn == nil || len(orn) < 2 {
		return nil
	}
	prefixOrn := strings.Split(orn[0], ":")
	if prefixOrn == nil || len(prefixOrn) < 6 || prefixOrn[0] != "orn" {
		return nil
	}
	result = &ResourceItem{
		Patrition:      prefixOrn[1],
		Service:        prefixOrn[2],
		Region:         prefixOrn[3],
		CustomerNumber: prefixOrn[4],
		ResourceType:   prefixOrn[5],
		Resource:       orn[1],
	}
	return result
}

func checkResourceSubItem(source, original string, isExact bool) bool {
	if source == "" {
		source = "*"
	}
	if original == "" {
		original = "*"
	}
	if original == "*" {
		return true
	}
	if !isExact {
		return strings.HasPrefix(source, original) // true
	}
	return source == original
}

// Contains contain an string item in string slice
func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
