package error

import "errors"

// App Error Definition
var (
	ErrNotFound            = errors.New("not found")
	ErrConflictResource    = errors.New("conflict resource")
	ErrForbidden           = errors.New("forbidden")
	ErrLockedResource      = errors.New("locked resource")
	ErrContextCancelled    = errors.New("context is canceled")
	ErrUnprocessableEntity = errors.New("unprocessable entity")
	ErrThroughputExceeded  = errors.New("resource throughput exceeded")
)
