package errors

// ToError Returns an application error from input. If The type
// is not of type Error, nil will be returned.
func ToError(err any) *Error {
	switch v := err.(type) {
	case *Error:
		return v
	case Error:
		return &v
	default:
		return nil
	}
}
