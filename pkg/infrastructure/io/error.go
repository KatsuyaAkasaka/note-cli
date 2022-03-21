package io

import "errors"

type NotFoundError struct {
	Err error
}

func (e *NotFoundError) Error() string {
	return e.Err.Error()
}

func IsErrNotFound(err error) bool {
	var nerr *NotFoundError

	return errors.As(err, &nerr)
}
