package errors

import (
	"context"
	"database/sql"
	"errors"
	"net"
	"strings"

	"github.com/lib/pq"
)

// pqErrorCodes maps PostgreSQL error codes to AppError.
// Reference: https://www.postgresql.org/docs/current/errcodes-appendix.html
var pqErrorCodes = map[string]func() *AppError{
	"23505": func() *AppError { return New(CodeConflict, "duplicate entry") },
	"23503": func() *AppError { return New(CodeBadRequest, "referenced record does not exist") },
	"23502": func() *AppError { return New(CodeBadRequest, "required field is null") },
	"23514": func() *AppError { return New(CodeBadRequest, "value violates check constraint") },
	"42P01": func() *AppError { return New(CodeDatabase, "table does not exist") },
	"53300": func() *AppError { return New(CodeUnavailable, "too many database connections") },
	"57014": func() *AppError { return New(CodeTimeout, "database query cancelled") },
}

// FromRuntime inspects a raw error and converts it to an AppError.
// Use this at the boundary between infrastructure and service layers.
//
// Example:
//
//	row, err := db.QueryContext(ctx, query)
//	if err != nil {
//	    return errors.FromRuntime(err)
//	}
func FromRuntime(err error) *AppError {
	if err == nil {
		return nil
	}

	// Already an AppError — pass through.
	var appErr *AppError
	if errors.As(err, &appErr) {
		return appErr
	}

	// Context errors.
	if errors.Is(err, context.DeadlineExceeded) || errors.Is(err, context.Canceled) {
		return Timeout(err)
	}

	// sql.ErrNoRows — common "not found" from database/sql.
	if errors.Is(err, sql.ErrNoRows) {
		return NotFound("record")
	}

	// PostgreSQL errors via lib/pq.
	var pqErr *pq.Error
	if errors.As(err, &pqErr) {
		if fn, ok := pqErrorCodes[string(pqErr.Code)]; ok {
			return fn()
		}
		return WrapWithDetail(CodeDatabase, "database error", pqErr.Message, err)
	}

	// Network errors.
	var netErr *net.OpError
	if errors.As(err, &netErr) {
		if netErr.Timeout() {
			return Timeout(err)
		}
		return Wrap(CodeUnavailable, "network error", err)
	}

	// Fallback: check error message for common patterns.
	msg := strings.ToLower(err.Error())
	switch {
	case strings.Contains(msg, "timeout") || strings.Contains(msg, "deadline"):
		return Timeout(err)
	case strings.Contains(msg, "connection refused") || strings.Contains(msg, "no such host"):
		return Wrap(CodeUnavailable, "service unavailable", err)
	case strings.Contains(msg, "duplicate") || strings.Contains(msg, "unique"):
		return Wrap(CodeConflict, "duplicate entry", err)
	}

	return Internal(err)
}

// IsNotFound reports whether err is a CodeNotFound AppError.
func IsNotFound(err error) bool {
	var e *AppError
	return errors.As(err, &e) && e.Code == CodeNotFound
}

// IsConflict reports whether err is a CodeConflict AppError.
func IsConflict(err error) bool {
	var e *AppError
	return errors.As(err, &e) && e.Code == CodeConflict
}

// IsUnauthorized reports whether err is a CodeUnauthorized AppError.
func IsUnauthorized(err error) bool {
	var e *AppError
	return errors.As(err, &e) && e.Code == CodeUnauthorized
}

// IsInternal reports whether err is a 5xx-class AppError.
func IsInternal(err error) bool {
	var e *AppError
	if !errors.As(err, &e) {
		return false
	}
	return e.HTTPStatus() >= 500
}
