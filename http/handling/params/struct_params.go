package params

import (
	"encoding/json"
	"io"
	"reflect"
	"strings"
)

func populateStructFromForm(structVal reflect.Value, formVals map[string][]string) (err error) {
	for i := 0; i < structVal.Elem().Type().NumField(); i++ {
		field := structVal.Elem().Type().Field(i)
		for key, vals := range formVals {
			if strings.EqualFold(key, field.Name) && len(vals) > 0 {
				fieldVal := structVal.Elem().Field(i)
				if fieldVal.CanSet() {
					valToSet, convErr := parseValueToType(fieldVal.Type(), vals[0])
					if convErr == nil {
						fieldVal.Set(valToSet)
					} else {
						err = convErr
					}
				}
			}
		}
	}
	return
}

func populateStructFromJSON(target reflect.Value, reader io.ReadCloser) error {
	return json.NewDecoder(reader).Decode(target.Interface())
}
