package goattribute

import (
	"errors"
	"reflect"
	"strconv"
)

func CopyInt(target interface{}, input interface{}) error {
	if input == nil {
		return errors.New("input is nil")
	}

	targetValue := reflect.ValueOf(target)
	if targetValue.Kind() != reflect.Ptr || !targetValue.Elem().CanSet() {
		return errors.New("target is not a pointer or not settable")
	}

	var intValue int64
	switch v := input.(type) {
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
		intValue = reflect.ValueOf(v).Convert(reflect.TypeOf(int64(0))).Int()
	case string:
		floatValue, parseErr := strconv.ParseFloat(v, 64)
		if parseErr != nil {
			return errors.New("cannot convert string to number")
		}
		intValue = int64(floatValue)
	default:
		return errors.New("unsupported input type")
	}

	targetValue.Elem().Set(reflect.ValueOf(intValue).Convert(targetValue.Elem().Type()))
	return nil
}
