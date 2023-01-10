package restlib

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

func FillStruct(data map[string]interface{}, result interface{}) error {
	for k, v := range data {
		err := SetField(result, k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func SetField(obj interface{}, name string, value interface{}) error {
	structValue := reflect.ValueOf(obj).Elem()
	structFieldValue := structValue.FieldByNameFunc(func(s string) bool { return strings.ToLower(s) == strings.ToLower(name) })
	if !structFieldValue.IsValid() {
		return fmt.Errorf("No such field: %s in obj", name)
	}
	if !structFieldValue.CanSet() {
		return fmt.Errorf("Cannot set %s field value", name)
	}
	structFieldType := structFieldValue.Type()
	val := reflect.ValueOf(value)
	if structFieldType != val.Type() {
		return errors.New("Provided value type didn't match obj field type")
	}
	structFieldValue.Set(val)
	return nil
}
