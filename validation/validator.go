package validation

type ValidationError struct {
	FieldName string
	Error     error
}

type Validator interface {
	Validate(data interface{}) (bool, []ValidationError)
}

type ValidatorFunc func(fieldName string, value interface{}, base string) (bool, error)

func DefaultValidators() map[string]ValidatorFunc {
	return map[string]ValidatorFunc{
		"required": required,
		"min":      min,
	}
}
