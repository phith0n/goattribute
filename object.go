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
var truePattern = regexp.MustCompile(`^y|Y|yes|Yes|YES|true|True|TRUE|on|On|ON|1$`)
var falsePattern = regexp.MustCompile(`^n|N|no|No|NO|false|False|FALSE|off|Off|OFF|0$`)

func New(obj interface{}) *Attribute {
	return &Attribute{
		obj: obj,
	}
}

func NewWithTag(obj interface{}, tagName string) *Attribute {
	return &Attribute{
		obj:     obj,
		tagName: tagName,
	}
}

type Attribute struct {
	tagName string
	obj     interface{}
}

func (a *Attribute) SetAttr(path string, value interface{}) error {
	keys := strings.Split(path, ".")
	currentValue := reflect.ValueOf(a.obj).Elem()

	for _, key := range keys {
		if !currentValue.IsValid() {
			return errors.New("filed not found in path")
		}

		tagToName := a.getTagMap(currentValue)
		matches := arrayPattern.FindStringSubmatch(key)
		if len(matches) == 3 {
			name := matches[1]
			index, _ := strconv.Atoi(matches[2])
			currentValue = currentValue.FieldByName(tagToName[name])
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

			currentValue = currentValue.FieldByName(tagToName[key])
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
		case reflect.Bool:
			var v = strings.TrimSpace(fmt.Sprintf("%v", value))
			if truePattern.MatchString(v) {
				currentValue.SetBool(true)
			} else if falsePattern.MatchString(v) {
				currentValue.SetBool(false)
			} else {
				return fmt.Errorf("cannot set %s to a boolean field", v)
			}
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

		tagToName := a.getTagMap(currentValue)
		matches := arrayPattern.FindStringSubmatch(key)
		if len(matches) == 3 {
			name := matches[1]
			index, _ := strconv.Atoi(matches[2])
			currentValue = currentValue.FieldByName(tagToName[name])
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
			currentValue = currentValue.FieldByName(tagToName[key])
		}
	}

	if currentValue.IsValid() && currentValue.CanInterface() {
		return currentValue.Interface(), nil
	} else {
		return nil, fmt.Errorf("invalid path")
	}
}

func (a *Attribute) GetObject() interface{} {
	return a.obj
}

func (a *Attribute) getTagMap(val reflect.Value) map[string]string {
	var m = make(map[string]string)
	var t reflect.Type

	if val.Kind() == reflect.Ptr && !val.IsNil() {
		t = val.Elem().Type()
	} else if val.Kind() == reflect.Struct {
		t = val.Type()
	} else {
		return m
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if a.tagName == "" {
			m[field.Name] = field.Name
			continue
		}

		tagValue := field.Tag.Get(a.tagName)
		if tagValue != "" {
			m[tagValue] = field.Name
		} else {
			m[field.Name] = field.Name
		}
	}

	return m
}
