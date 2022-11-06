package errorHandler

type ComplexError struct {
	err         string
	descriptors []string
}

func (ce *ComplexError) Error() string {
	return ce.err
}

func newComplexError(_err string, _desc ...string) *ComplexError {
	return &ComplexError{err: _err, descriptors: _desc}
}

func Complex(_err error) *ComplexError {
	return &ComplexError{err: _err.Error()}
}

func (ce *ComplexError) GetDescriptors() []string {
	return ce.descriptors
}

func (ce *ComplexError) SetDesc(newDesc ...string) *ComplexError {
	newErr := *ce
	newErr.descriptors = append(newErr.descriptors, newDesc...)
	return &newErr
}
