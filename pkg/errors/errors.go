package errors

func New(message string, err error) error {
	return &tracedError{
		err: &anyError{
			message: message,
			err:     err,
		},
		location: Locate(),
	}
}

func TracedToStackTrace(err error) []Location {
	if e, ok := err.(*tracedError); ok {
		return e.location
	}

	return nil
}
