package defaults

import (
	"errors"
	"reflect"

	"gopkg.in/yaml.v3"
)

var (
	errInvalidType = errors.New("invalid type")
)

const (
	fieldName = "default"
)

func Set(p any) (err error) {
	typ := reflect.TypeOf(p)
	if typ.Kind() != reflect.Ptr {
		return errInvalidType
	}
	val := reflect.ValueOf(p).Elem()
	if val.Kind() != reflect.Struct {
		return errInvalidType
	}
	typ = val.Type()
	for i := 0; i < typ.NumField(); i++ {
		fieldVal := val.Field(i)
		defaultVal := typ.Field(i).Tag.Get(fieldName)
		if defaultVal == "-" {
			continue
		}
		if err = setField(fieldVal, defaultVal); err != nil {
			return err
		}
	}
	return nil
}

func setField(field reflect.Value, defaultVal string) (err error) {
	if !field.CanSet() {
		return nil
	}

	if defaultVal != "" && isInitialValue(field) {
		if err = yaml.Unmarshal([]byte(defaultVal), field.Addr().Interface()); err != nil {
			return err
		}
	}

	switch field.Kind() {
	case reflect.Ptr:
		if field.Elem().Kind() == reflect.Struct {
			return Set(field.Elem().Addr().Interface())
		}
	case reflect.Struct:
		return Set(field.Addr().Interface())
	case reflect.Slice:
		for j := 0; j < field.Len(); j++ {
			if err := setField(field.Index(j), ""); err != nil {
				return err
			}
		}
	case reflect.Map:
		for _, e := range field.MapKeys() {
			var v = field.MapIndex(e)
			switch v.Kind() {
			case reflect.Ptr:
				switch v.Elem().Kind() {
				case reflect.Struct, reflect.Slice, reflect.Map:
					if err := setField(v.Elem(), ""); err != nil {
						return err
					}
				}
			case reflect.Struct, reflect.Slice, reflect.Map:
				ref := reflect.New(v.Type())
				ref.Elem().Set(v)
				if err := setField(ref.Elem(), ""); err != nil {
					return err
				}
				field.SetMapIndex(e, ref.Elem().Convert(v.Type()))
			}
		}
	}
	return nil
}

func isInitialValue(field reflect.Value) bool {
	return reflect.DeepEqual(reflect.Zero(field.Type()).Interface(), field.Interface())
}
