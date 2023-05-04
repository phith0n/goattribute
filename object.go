package goattribute

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

var arrayPattern = regexp.MustCompile(`(\w+)\[(\d+)\]`)

func New(obj interface{}) *Attribute {
	return &Attribute{
		obj: obj,
	}
}

type Attribute struct {
	obj interface{}
}

func (a *Attribute) SetAttr(path string, value interface{}) error {
	keys := strings.Split(path, ".")
	currentValue := reflect.ValueOf(a.obj).Elem()

	for _, key := range keys {
		if !currentValue.IsValid() {
			return errors.New("filed not found in path")
		}

		matches := arrayPattern.FindStringSubmatch(key)
		if len(matches) == 3 {
			field := matches[1]
			index, _ := strconv.Atoi(matches[2])
			currentValue = currentValue.FieldByName(field)
			if currentValue.Kind() == reflect.Ptr {
				currentValue = currentValue.Elem()
				if !currentValue.IsValid() {
					return errors.New("cannot set field on a nil object")
				}
			}

			if currentValue.Kind() == reflect.Slice && currentValue.Len() > index {
				currentValue = currentValue.Index(index)
			} else {
				return errors.New("field type not slice")
			}
		} else {
			if currentValue.Kind() == reflect.Ptr {
				currentValue = currentValue.Elem()
				if !currentValue.IsValid() {
					return errors.New("cannot set field on a nil object")
				}
			}

			currentValue = currentValue.FieldByName(key)
		}
	}

	var err error
	if currentValue.IsValid() && currentValue.CanSet() {
		switch currentValue.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8,
			reflect.Uint16, reflect.Uint32, reflect.Uint64:
			var i int64
			err = CopyInt(&i, value)
			if err == nil {
				currentValue.SetInt(i)
			}
		case reflect.String:
			currentValue.SetString(fmt.Sprintf("%v", value))
		default:
			return errors.New("unsupported type")
		}
	} else {
		return errors.New("invalid path or not settable")
	}

	return err
}

func (a *Attribute) GetAttr(path string) (interface{}, error) {
	keys := strings.Split(path, ".")
	currentValue := reflect.ValueOf(a.obj).Elem()

	for _, key := range keys {
		if !currentValue.IsValid() {
			return nil, errors.New("filed not found in path")
		}

		matches := arrayPattern.FindStringSubmatch(key)
		if len(matches) == 3 {
			field := matches[1]
			index, _ := strconv.Atoi(matches[2])
			currentValue = currentValue.FieldByName(field)
			if currentValue.Kind() == reflect.Ptr {
				currentValue = currentValue.Elem()
			}

			if currentValue.Kind() == reflect.Slice && currentValue.Len() > index {
				currentValue = currentValue.Index(index)
			} else {
				return nil, errors.New("unsupported type")
			}
		} else {
			if currentValue.Kind() == reflect.Ptr {
				currentValue = currentValue.Elem()
			}
			currentValue = currentValue.FieldByName(key)
		}
	}

	if currentValue.IsValid() {
		return currentValue.Interface(), nil
	} else {
		return nil, fmt.Errorf("invalid path")
	}
}

func (a *Attribute) GetObject() interface{} {
	return a.obj
}
