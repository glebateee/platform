package params

import (
	"errors"
	"reflect"
)

func getParametersFromURLValues(funcType reflect.Type, urlVals []string) (params []reflect.Value, err error) {
	if len(urlVals) == funcType.NumIn()-1 {
		params = make([]reflect.Value, len(urlVals))
		for i := range urlVals {
			params[i], err = parseValueToType(funcType.In(i+1), urlVals[i])
			if err != nil {
				return
			}
		}
	} else {
		err = errors.New("parameter number mismatch")
	}
	return
}
