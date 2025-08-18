package params

import (
	"fmt"
	"reflect"
	"strconv"
)

func parseValueToType(target reflect.Type, requestParam string) (result reflect.Value, err error) {
	switch target.Kind() {
	case reflect.String:
		result = reflect.ValueOf(requestParam)
	case reflect.Int:
		iVal, convErr := strconv.Atoi(requestParam)
		if convErr == nil {
			result = reflect.ValueOf(iVal)
		} else {
			return reflect.Value{}, convErr
		}
	case reflect.Float64:
		fVal, convErr := strconv.ParseFloat(requestParam, 64)
		if convErr == nil {
			result = reflect.ValueOf(fVal)
		} else {
			return reflect.Value{}, convErr
		}
	case reflect.Bool:
		bVal, convErr := strconv.ParseBool(requestParam)
		if convErr == nil {
			result = reflect.ValueOf(bVal)
		} else {
			return reflect.Value{}, convErr
		}
	default:
		err = fmt.Errorf("Cannot use type %v as handler method parameter", target.Name())
	}
	return
}
