package errors

import (
	"fmt"
)

type anyError struct {
	message string
	err     error
}

func (e *anyError) Error() string {
	return fmt.Sprintf("%s-%s", e.message, e.err.Error())
}

type tracedError struct {
	err      error
	location []Location
}

func (e *tracedError) Error() string {
	return e.err.Error()
}

type baseError struct {
	code    int
	message string
}
