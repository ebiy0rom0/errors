package errors

import "errors"

// Is the wrapper for the standard errors.Is().
// It's returns same result.
func Is(err, target error) bool { return errors.Is(err, target) }

// As the wrapper for the standard errors.As().
// It's returns same result.
func As(err error, target any) bool { return errors.As(err, target) }
