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
			}

			if currentValue.IsValid() && currentValue.Kind() == reflect.Slice {
				currentValue = currentValue.Index(index)
			} else {
				return errors.New("field type not slice")
			}
		} else {
			if currentValue.Kind() == reflect.Ptr {
				currentValue = currentValue.Elem()
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

func (a *Attribute) ToString() string {
	return fmt.Sprintf("%v", a.obj)
}

func (a *Attribute) ToInt() (int, error) {
	var i int
	err := CopyInt(&i, a.obj)
	return i, err
}

func (a *Attribute) ToInt8() (int8, error) {
	var i int8
	err := CopyInt(&i, a.obj)
	return i, err
}

func (a *Attribute) ToInt16() (int16, error) {
	var i int16
	err := CopyInt(&i, a.obj)
	return i, err
}

func (a *Attribute) ToInt32() (int32, error) {
	var i int32
	err := CopyInt(&i, a.obj)
	return i, err
}

func (a *Attribute) ToInt64() (int64, error) {
	var i int64
	err := CopyInt(&i, a.obj)
	return i, err
}

func (a *Attribute) ToUInt() (uint, error) {
	var i uint
	err := CopyInt(&i, a.obj)
	return i, err
}

func (a *Attribute) ToUInt8() (uint8, error) {
	var i uint8
	err := CopyInt(&i, a.obj)
	return i, err
}

func (a *Attribute) ToUInt16() (uint16, error) {
	var i uint16
	err := CopyInt(&i, a.obj)
	return i, err
}

func (a *Attribute) ToUInt32() (uint32, error) {
	var i uint32
	err := CopyInt(&i, a.obj)
	return i, err
}

func (a *Attribute) ToUInt64() (uint64, error) {
	var i uint64
	err := CopyInt(&i, a.obj)
	return i, err
}
