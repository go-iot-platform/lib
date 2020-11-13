package structs

import (
	"encoding/json"
	"reflect"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/go-iot-platform/lib/marshaller"
	structs "github.com/go-iot-platform/structs-fatih"
)

// Merge receives two structs, and merges them excluding fields with tag name: `structs`, value "-"
func Merge(dst, src interface{}) {
	s := reflect.ValueOf(src)
	d := reflect.ValueOf(dst)
	if s.Kind() != reflect.Ptr || d.Kind() != reflect.Ptr {
		return
	}
	for i := 0; i < s.Elem().NumField(); i++ {
		v := s.Elem().Field(i)
		fieldName := s.Elem().Type().Field(i).Name
		skip := s.Elem().Type().Field(i).Tag.Get("structs")
		if skip == "-" {
			continue
		}
		if v.Kind() > reflect.Float64 &&
			v.Kind() != reflect.String &&
			v.Kind() != reflect.Struct &&
			v.Kind() != reflect.Ptr &&
			v.Kind() != reflect.Slice {
			continue
		}
		if v.Kind() == reflect.Ptr {
			// Field is pointer check if it's nil or set
			if !v.IsNil() {
				// Field is set assign it to dest

				if d.Elem().FieldByName(fieldName).Kind() == reflect.Ptr {
					d.Elem().FieldByName(fieldName).Set(v)
					continue
				}
				f := d.Elem().FieldByName(fieldName)
				if f.IsValid() {
					f.Set(v.Elem())
				}
			}
			continue
		}
		d.Elem().FieldByName(fieldName).Set(v)
	}
}

// MergeOverwrite overwrite map
func MergeOverwrite(to, from, dst interface{}) error {
	var toMap map[string]interface{}
	var result map[string]interface{}
	var fromMap map[string]interface{}
	kindTo := reflect.ValueOf(to)
	kindFrom := reflect.ValueOf(from)
	if kindTo.Kind() == reflect.Map {
		toMap = to.(map[string]interface{})
	} else {
		toMap = structs.Map(to)
	}
	if kindFrom.Kind() == reflect.Map {
		fromMap = from.(map[string]interface{})
	} else {
		fromMap = structs.Map(from)
	}
	for k, v := range fromMap {
		mapBase(k, v, &fromMap)
		toMap[k] = fromMap[k]
	}
	_, ok := toMap["Base"]
	if ok {
		delete(toMap, "Base")
	}
	conv := marshaller.ConventionalMarshaller{Value: toMap}
	b, err := conv.MarshalJSON()
	if err != nil {
		return err
	}
	json.Unmarshal(b, &result)
	if err := mapstructure.Decode(toMap, dst); err != nil {
		return err
	}
	return nil
}

// MergeOverwriteCamel overwrite map to cammel case
func MergeOverwriteCamel(to, from, dst interface{}) error {
	var toMap map[string]interface{}
	var result map[string]interface{}
	var fromMap map[string]interface{}
	kindTo := reflect.ValueOf(to)
	kindFrom := reflect.ValueOf(from)
	if kindTo.Kind() == reflect.Map {
		toMap = to.(map[string]interface{})
	} else {
		toMap = structs.Map(to)
	}
	if kindFrom.Kind() == reflect.Map {
		fromMap = from.(map[string]interface{})
	} else {
		fromMap = structs.Map(from)
	}
	for k, v := range fromMap {
		mapBase(k, v, &fromMap)
		toMap[k] = fromMap[k]
	}
	_, ok := toMap["Base"]
	if ok {
		delete(toMap, "Base")
	}
	conv := marshaller.ConventionalMarshaller{Value: toMap}
	b, err := conv.MarshalJSON()
	if err != nil {
		return err
	}
	json.Unmarshal(b, &result)
	if err := mapstructure.Decode(result, dst); err != nil {
		return err
	}
	return nil
}

// MergeToMap merge struct  to map
func MergeToMap(to, from interface{}) (map[string]interface{}, error) {
	var toMap map[string]interface{}
	// var result map[string]interface{}
	var fromMap map[string]interface{}
	kindTo := reflect.ValueOf(to)
	kindFrom := reflect.ValueOf(from)
	if kindTo.Kind() == reflect.Map {
		toMap = to.(map[string]interface{})
	} else {
		toMap = structs.Map(to)
	}
	if kindFrom.Kind() == reflect.Map {
		fromMap = from.(map[string]interface{})
	} else {
		fromMap = structs.Map(from)
	}
	for k, v := range fromMap {
		mapBase(k, v, &fromMap)
		toMap[k] = fromMap[k]
	}
	_, ok := toMap["Base"]
	if ok {
		delete(toMap, "Base")
	}
	return toMap, nil
}

// ConvertStringToMap convert string json to map
func ConvertStringToMap(in string) map[string]interface{} {
	var result map[string]interface{}
	d := json.NewDecoder(strings.NewReader(in))
	d.UseNumber()

	if err := d.Decode(&result); err != nil {
		return nil
	}
	return result
}
func mapBase(key string, value interface{}, org *(map[string]interface{})) {
	result := *org
	if key == "Base" {
		temp := value.(map[string]interface{})
		for k, v := range temp {
			result[k] = v
		}
		delete(result, "Base")
		return
	}
	valueRf := reflect.ValueOf(value)
	kind := valueRf.Kind()
	switch kind {
	case reflect.Map:
		temp := value.(map[string]interface{})
		for k, v := range temp {
			mapBase(k, v, &temp)
		}
	case reflect.Slice:
		if reflect.TypeOf(value).String() == "[]string" {
			result[key] = value
			return
		}
		temp := value.([]interface{})
		for _, v := range temp {
			if reflect.TypeOf(v).String() == "string" {
				// fmt.Println(temp)
				result[key] = value
				return
			}
			temp2 := v.(map[string]interface{})
			for k1, v1 := range temp2 {
				mapBase(k1, v1, &temp2)
			}
		}
	case reflect.Ptr:
		if !valueRf.IsNil() {
			// temp := valueRf.Elem()
			// fmt.Println(reflect.TypeOf(value))
			// fmt.Printf("item:%v\n", temp)
			// temp = temp.(map[string]interface{})
			// for k, v := range temp {
			// 	mapBase(k, v, &temp)
			// }
			// value = &temp
		}
	case reflect.String:
		result[key] = value
	}
	org = &result
}
