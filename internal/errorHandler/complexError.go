package errorHandler

import "errors"

type ComplexError struct {
	err         error
	descriptors []string
}

func (ce *ComplexError) Error() string {
	return ce.err.Error()
}

func newComplexError(_err string, _desc ...string) *ComplexError {
	return &ComplexError{err: errors.New(_err), descriptors: _desc}
}

func (ce *ComplexError) GetDescriptors() []string {
	return ce.descriptors
}

func (ce *ComplexError) SetDesc(newDesc ...string) *ComplexError {
	newErr := *ce
	newErr.descriptors = append(newErr.descriptors, newDesc...)
	return &newErr
}
