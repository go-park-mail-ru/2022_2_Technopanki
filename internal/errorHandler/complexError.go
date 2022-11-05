package errorHandler

import "fmt"

type ComplexError struct {
	err         string
	descriptors map[string]any
}

func (ce *ComplexError) Error() string {
	return ce.err
}

func newComplexError(_err string, _desc map[string]any) error {
	return &ComplexError{err: _err, descriptors: _desc}
}

func newNonDescError(_err string) error {
	return newComplexError(_err, map[string]any{})
}

func newSimpleDescError(_err string, key string, value any) error {
	return newComplexError(_err, map[string]any{key: value})
}

func (ce *ComplexError) GetDescriptors(key string) (any, error) {
	desc, ok := ce.descriptors[key]
	if !ok {
		return nil, fmt.Errorf("descriptor not fount")
	}
	return desc, nil
}

func (ce *ComplexError) SetDesc(key string, value any) error {
	newErr := ce
	newErr.descriptors[key] = value
	return ce
}
