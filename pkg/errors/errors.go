package errors

import (
	"errors"
	"fmt"
	"net/http"
)

type Code string

const (
	// 4xx — client errors
	CodeNotFound        Code = "NOT_FOUND"
	CodeUnauthorized    Code = "UNAUTHORIZED"
	CodeForbidden       Code = "FORBIDDEN"
	CodeBadRequest      Code = "BAD_REQUEST"
	CodeValidation      Code = "VALIDATION_ERROR"
	CodeConflict        Code = "CONFLICT"
	CodeTooManyRequests Code = "TOO_MANY_REQUESTS"
	CodeUnprocessable   Code = "UNPROCESSABLE_ENTITY"

	// 5xx — server errors
	CodeInternal    Code = "INTERNAL_ERROR"
	CodeUnavailable Code = "SERVICE_UNAVAILABLE"
	CodeTimeout     Code = "TIMEOUT"
	CodeDatabase    Code = "DATABASE_ERROR"
)

var httpStatusMap = map[Code]int{
	CodeNotFound:        http.StatusNotFound,
	CodeUnauthorized:    http.StatusUnauthorized,
	CodeForbidden:       http.StatusForbidden,
	CodeBadRequest:      http.StatusBadRequest,
	CodeValidation:      http.StatusUnprocessableEntity,
	CodeConflict:        http.StatusConflict,
	CodeTooManyRequests: http.StatusTooManyRequests,
	CodeUnprocessable:   http.StatusUnprocessableEntity,
	CodeInternal:        http.StatusInternalServerError,
	CodeUnavailable:     http.StatusServiceUnavailable,
	CodeTimeout:         http.StatusGatewayTimeout,
	CodeDatabase:        http.StatusInternalServerError,
}

type AppError struct {
	Code    Code   // machine-readable code
	Message string // human-readable message (safe to expose to client)
	Detail  string // optional internal detail (for logging only)
	Err     error  // original/wrapped error
}

func (e *AppError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("[%s] %s: %v", e.Code, e.Message, e.Err)
	}
	return fmt.Sprintf("[%s] %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
	return e.Err
}

// HTTPStatus returns the HTTP status code for this error.
func (e *AppError) HTTPStatus() int {
	if status, ok := httpStatusMap[e.Code]; ok {
		return status
	}
	return http.StatusInternalServerError
}

// Is enables errors.Is() matching by Code.
func (e *AppError) Is(target error) bool {
	var t *AppError
	if errors.As(target, &t) {
		return e.Code == t.Code
	}
	return false
}

// — Constructors —

func New(code Code, message string) *AppError {
	return &AppError{Code: code, Message: message}
}

func Wrap(code Code, message string, err error) *AppError {
	return &AppError{Code: code, Message: message, Err: err}
}

func WrapWithDetail(code Code, message, detail string, err error) *AppError {
	return &AppError{Code: code, Message: message, Detail: detail, Err: err}
}

// — Shorthand constructors —

func NotFound(entity string) *AppError {
	return New(CodeNotFound, fmt.Sprintf("%s not found", entity))
}

func Unauthorized(msg string) *AppError {
	return New(CodeUnauthorized, msg)
}

func Forbidden(msg string) *AppError {
	return New(CodeForbidden, msg)
}

func BadRequest(msg string) *AppError {
	return New(CodeBadRequest, msg)
}

func Validation(msg string) *AppError {
	return New(CodeValidation, msg)
}

func Conflict(msg string) *AppError {
	return New(CodeConflict, msg)
}

func Internal(err error) *AppError {
	return Wrap(CodeInternal, "internal server error", err)
}

func Timeout(err error) *AppError {
	return Wrap(CodeTimeout, "request timed out", err)
}
