package domain

import "fmt"

// UseCaseError is an error that can return a UseCase
type UseCaseError struct {
	Code   uint16
	Reason string
	Fix    string
}

func (caseErr UseCaseError) Error() string {
	return fmt.Sprintf("UseCaseError -> [%d] %s: %s", caseErr.Code, caseErr.Reason, caseErr.Fix)
}

var (
	MalformedRequestErr = UseCaseError{Code: 000, Reason: "Bad request", Fix: "See documentation and try again"}
	InternalError       = UseCaseError{Code: 001, Reason: "Internal Error", Fix: ""}
)
