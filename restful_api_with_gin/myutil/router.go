package myutil

import "errors"

type StatusError struct {
	status int
	err    error
}
type StatusErrorer interface {
	Error() string
	Status() int
}

func (se *StatusError) Error() string {
	return se.err.Error()
}

func (se *StatusError) AttachError(status int, err error) {
	se.status = status
	se.err = err
}

func (se *StatusError) Status() int {
	return se.status
}

func NewStatusError(status int, err error) StatusErrorer {
	return &StatusError{
		status: status,
		err:    err,
	}
}
func NewStatusErrorString(status int, str string) StatusErrorer {
	return &StatusError{
		status: status,
		err:    errors.New(str),
	}
}
