package validator

import (
	"sensorapi/src/domain"

	playgroundValidator "gopkg.in/go-playground/validator.v9"
)

// PlaygroundValidator is a validator implemented using GoPlaygroundValidator v9
type PlaygroundValidator struct {
	validator *playgroundValidator.Validate
}

// NewPlaygroundValidator builds a new PlaygroundValidator
func NewPlaygroundValidator() domain.Validator {
	return PlaygroundValidator{validator: playgroundValidator.New()}
}

// Validate a struct using annotations
func (val PlaygroundValidator) Validate(target interface{}) error {
	return val.validator.Struct(target)
}
