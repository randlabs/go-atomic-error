package exterror

import (
	"fmt"
	"sort"
)

// -----------------------------------------------------------------------------

// AugmentedError is just an error with extended data.
type AugmentedError struct {
	Message string
	Fields  map[string]interface{}
	Err     error // Underlying error that occurred during the operation.
}

// -----------------------------------------------------------------------------

// NewAugmentedError creates a new AugmentedError.
func NewAugmentedError(wrappedErr error, text string, fields map[string]interface{}) *AugmentedError {
	e := AugmentedError{}
	e.Message = text
	e.Fields = fields
	e.Err = wrappedErr
	return &e
}

// Unwrap returns the underlying error.
func (e *AugmentedError) Unwrap() error {
	return e.Err
}

// Error returns a string representation of the error.
func (e *AugmentedError) Error() string {
	if e == nil {
		return ""
	}
	s := e.Message
	if e.Fields != nil {
		keys := make([]string, 0, len(e.Fields))
		for k := range e.Fields {
			keys = append(keys, k)
		}
		sort.Strings(keys)

		for _, k := range keys {
			s += fmt.Sprintf(" [%s=%v]", k, e.Fields[k])
		}
	}
	if e.Err != nil {
		s += " [err=" + e.Err.Error() + "]"
	}
	return s
}
