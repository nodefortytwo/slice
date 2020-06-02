package slice

type SliceError struct {
	error string
}

func (err *SliceError) Error() string {
	return err.error
}
