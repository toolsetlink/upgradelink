package strx

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// StructToString converts a struct to a string. The string is snake case.
func StructToString(v any) (string, error) {
	val := reflect.ValueOf(v)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	if val.Kind() != reflect.Struct {
		return "", fmt.Errorf("the input value must be structure or structure pointer")
	}

	var result strings.Builder
	for i := 0; i < val.NumField(); i++ {
		fieldVal := val.Field(i)
		if i > 0 {
			result.WriteString("_")
		}
		switch fieldVal.Kind() {
		case reflect.Slice, reflect.Array:
			result.WriteString(fmt.Sprint(fieldVal))
		case reflect.Struct:
			tmp, err := StructToString(fieldVal.Interface())
			if err != nil {
				return "", err
			}
			result.WriteString(tmp)
		case reflect.String:
			result.WriteString(fieldVal.String())
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			result.WriteString(strconv.FormatInt(fieldVal.Int(), 10))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			result.WriteString(strconv.FormatUint(fieldVal.Uint(), 10))
		case reflect.Float32, reflect.Float64:
			result.WriteString(strconv.FormatFloat(fieldVal.Float(), 'f', -1, 64))
		case reflect.Bool:
			result.WriteString(strconv.FormatBool(fieldVal.Bool()))
		default:
			continue
		}
	}
	return result.String(), nil
}
