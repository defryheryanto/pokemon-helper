package errors

import (
	"net/http"
	"strings"
)

func newBaseError(code int, msg string) baseError {
	return baseError{
		code:    code,
		message: msg,
	}
}

func (err baseError) Error() string {
	return strings.ToLower(err.message)
}

func (err baseError) Message() string {
	return err.message
}

type UnauthorizedError struct {
	baseError
}

func NewUnauthorizedError(message string) UnauthorizedError {
	return UnauthorizedError{
		baseError: newBaseError(http.StatusUnauthorized, message),
	}
}

type NotFoundError struct {
	baseError
}

func NewNotFoundError(message string) NotFoundError {
	return NotFoundError{
		baseError: newBaseError(http.StatusUnauthorized, message),
	}
}

type ForbiddenError struct {
	baseError
}

func NewForbiddenError(message string) ForbiddenError {
	return ForbiddenError{
		baseError: newBaseError(http.StatusUnauthorized, message),
	}
}

type BadRequestError struct {
	baseError
}

func NewBadRequestError(message string) BadRequestError {
	return BadRequestError{
		baseError: newBaseError(http.StatusUnauthorized, message),
	}
}
