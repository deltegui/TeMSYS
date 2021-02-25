package temsys

import "fmt"

// UseCaseError is an error that can return a UseCase
type UseCaseError struct {
	Code   uint16
	Reason string
}

func (caseErr UseCaseError) Error() string {
	return fmt.Sprintf("UseCaseError -> [%d] %s", caseErr.Code, caseErr.Reason)
}

// Common errors
var (
	MalformedRequestErr = UseCaseError{Code: 000, Reason: "Bad request"}
	InternalErr         = UseCaseError{Code: 001, Reason: "Internal Error"}
	UpdateErr           = UseCaseError{Code: 002, Reason: "Error while updating your data"}
)

// Users errors
var (
	UserNotFoundErr     = UseCaseError{Code: 100, Reason: "User not found"}
	InvalidPasswordErr  = UseCaseError{Code: 101, Reason: "Invalid password"}
	UserAlreadyExitsErr = UseCaseError{Code: 102, Reason: "User already exits"}
)

// Token errors
var (
	InvalidTokenErr = UseCaseError{Code: 301, Reason: "Invalid token"}
	NotAuthErr      = UseCaseError{Code: 300, Reason: "There is no 'Authorization' header"}
	ExpiredTokenErr = UseCaseError{Code: 302, Reason: "Token is expired"}
	OnlyAdminErr    = UseCaseError{Code: 303, Reason: "Endpoint is only available for users with 'admin' role"}
)

// Sensor errors
var (
	SensorNotFoundErr   = UseCaseError{Code: 200, Reason: "Sensor not found"}
	SensorNotRespondErr = UseCaseError{Code: 201, Reason: "Sensor does not respond"}
	SensorAlreadyExist  = UseCaseError{Code: 202, Reason: "Sensor already exists"}
)

// Report errors
var (
	ReportTypeDoesNotExists = UseCaseError{Code: 300, Reason: "ReportType does not exists"}
	ReportTypeAlreadyExists = UseCaseError{Code: 301, Reason: "ReportType already exists"}
)
