package bundle

import (
	"errors"
	"fmt"
)

// Errors represents a map of errors that occurred during a bundle load indexed by bundle name.
type Errors []Error

func (e Errors) Unwrap() []error {
	output := make([]error, len(e))
	for i := range e {
		output[i] = e[i]
	}
	return output
}
func (e Errors) Error() string {
	err := errors.Join(e.Unwrap()...)
	return err.Error()
}

type Error struct {
	name string
	err  error
}

func (e Error) Error() string {
	return fmt.Sprintf("'%s': %v", e.name, e.err)
}

func (e Error) Unwrap() error {
	return e.err
}
