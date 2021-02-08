package temsys

import "log"

// Presenter is anything that have de hability to represent data or errors to the user.
type Presenter interface {
	Present(data interface{})
	PresentError(data error)
}

// UseCaseRequest is anything that stores information needed by his UseCase.
type UseCaseRequest interface{}

// UseCaseResponse is anything returned by the computation of a UseCase.
type UseCaseResponse interface{}

// EmptyRequest is the UseCaseRequest to use if you dont want to pass anything to a UseCase.
// Use this instead calling it with nil. Improves readability.
var EmptyRequest UseCaseRequest = struct{}{}

// UseCase is anything that have application domain code.
type UseCase interface {
	Exec(Presenter, UseCaseRequest)
}

// Validator is anything that can validate a struct, returning
// information about invalid fields or an error.
type Validator interface {
	Validate(interface{}) ([]string, error)
}

// RequestValidator is a wrapper over a UseCase to perform request validation before
// the use case can process it.
type RequestValidator struct {
	inner     UseCase
	validator Validator
}

// Validate a UseCase creating a wrapper over it.
func Validate(useCase UseCase, validator Validator) UseCase {
	return RequestValidator{
		inner:     useCase,
		validator: validator,
	}
}

// Exec request validation.
func (reqVal RequestValidator) Exec(p Presenter, req UseCaseRequest) {
	valErrs, err := reqVal.validator.Validate(req)
	if err != nil {
		log.Printf("Error validating request: %s\n", err)
		return
	}
	if len(valErrs) != 0 {
		p.PresentError(UseCaseError{
			Code:   0,
			Reason: valErrs[0],
		})
		return
	}
	reqVal.inner.Exec(p, req)
}
