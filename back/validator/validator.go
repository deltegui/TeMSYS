package validator

import (
	"temsys"

	playgroundValidator "github.com/go-playground/validator/v10"
)

// PlaygroundValidator is a validator implemented using GoPlaygroundValidator v9
type PlaygroundValidator struct {
	validator *playgroundValidator.Validate
}

// NewPlaygroundValidator builds a new PlaygroundValidator
func New() temsys.Validator {
	return PlaygroundValidator{validator: playgroundValidator.New()}
}

// Validate a struct using annotations
func (val PlaygroundValidator) Validate(target interface{}) ([]string, error) {
	err := val.validator.Struct(target)
	if err != nil {
		switch err.(type) {
		case playgroundValidator.ValidationErrors:
			{
				valErr := err.(playgroundValidator.ValidationErrors)
				return errorsToString(valErr), nil
			}
		default:
			return nil, err
		}
	}
	return make([]string, 0, 0), nil
}

func errorsToString(ee playgroundValidator.ValidationErrors) []string {
	result := make([]string, len(ee), len(ee))
	for i, e := range ee {
		result[i] = e.Error()
	}
	return result
}
