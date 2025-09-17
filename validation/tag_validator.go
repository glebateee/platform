package validation

import (
	"reflect"
	"strings"
)

type TagValidator struct {
	validators map[string]ValidatorFunc
}

func NewDefaultValidator(map[string]ValidatorFunc) Validator {
	return &TagValidator{DefaultValidators()}
}

func (tv *TagValidator) Validate(data interface{}) (ok bool, errs []ValidationError) {
	dataValue := reflect.ValueOf(data)
	if dataValue.Kind() == reflect.Ptr {
		dataValue = dataValue.Elem()
	}
	if dataValue.Kind() != reflect.Struct {
		panic("Only structs can be validated")
	}
	for i := range dataValue.Type().NumField() {
		fieldType := dataValue.Type().Field(i)
		validationTag, found := fieldType.Tag.Lookup("validation")
		if found {
			for _, valTag := range strings.Split(validationTag, ",") {
				name, base := "", ""
				if strings.Contains(valTag, ":") {
					nameAndBase := strings.SplitN(valTag, ":", 2)
					name, base = nameAndBase[0], nameAndBase[1]
				} else {
					name = valTag
				}
				if validatorFunc, ok := tv.validators[name]; ok {
					valid, err := validatorFunc(fieldType.Name, dataValue.Field(i).Interface(), base)
					if !valid {
						errs = append(errs, ValidationError{
							FieldName: fieldType.Name,
							Error:     err})
					}
				} else {
					panic("Unknown validator: " + name)
				}
			}
		}
	}
	ok = len(errs) == 0
	return
}
