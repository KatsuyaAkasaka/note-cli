package io

import "errors"

type ErrNotFound struct {
	Err error
}

func (e *ErrNotFound) Error() string {
	return e.Err.Error()
}

func IsErrNotFound(err error) bool {
	var nerr *ErrNotFound
	return errors.As(err, &nerr)
}
