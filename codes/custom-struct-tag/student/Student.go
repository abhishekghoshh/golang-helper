package student

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type Student struct {
	Age  int    `validate:"min=18"`
	Name string `validate:"required"`
}

func (s *Student) Validate() error {
	// get the value of interface{}/ pointer point to
	val := reflect.Indirect(reflect.ValueOf(s))
	for i := 0; i < val.NumField(); i++ {
		typeField := val.Type().Field(i) // get field i-th of type(val)
		tag := typeField.Tag.Get("validate")
		if tag == "" {
			continue
		}
		fmt.Println(tag)
		// get the value of field like 17 or "Aiden"
		valueField := val.Field(i)
		fmt.Println(valueField)
		// split the tag so we can use like this: `required:"limit=20"
		rules := strings.Split(tag, ",")
		fmt.Println(rules)
		for _, rule := range rules {
			parts := strings.Split(rule, "=")
			key := parts[0]
			var value string
			if len(parts) > 1 {
				value = parts[1]
			}
			switch key {
			case "required":
				if err := s.isRequired(valueField); err != nil {
					return err
				}
			case "min":
				if err := s.isMin(valueField, value); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func (s *Student) isMin(valueField reflect.Value, minStr string) error {
	typeField := valueField.Type()
	if minStr == "" {
		return nil
	}
	min, err := strconv.ParseFloat(minStr, 64)
	if err != nil {
		return fmt.Errorf("min value %f is not a number", min)
	}
	switch valueField.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if float64(valueField.Int()) < min {
			return fmt.Errorf("field %s must be greater or equal %d", typeField.Name(), int(min))
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if float64(valueField.Uint()) < min {
			return fmt.Errorf("field %s must be greater or equal than %d", typeField.Name(), uint(min))
		}
	case reflect.Float32, reflect.Float64:
		if valueField.Float() < min {
			return fmt.Errorf("field %s must be greater or equal than %f", typeField.Name(), min)
		}
	}
	return nil
}
func (s *Student) isRequired(v reflect.Value) error {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		if v.Len() != 0 {
			return nil
		}
	case reflect.Bool:
		if v.Bool() {
			return nil
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		if v.Int() != 0 {
			return nil
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		if v.Uint() != 0 {
			return nil
		}
	}
	return fmt.Errorf("field %s is required", v.Type().Name())
}
