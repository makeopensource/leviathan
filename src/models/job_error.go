package models

type JobError interface {
	// Reason will be displayed to the end user, providing a user-friendly message.
	Reason() string
	// Err err parameter holds the underlying error, used for debugging purposes.
	Err() error
	// ErrStr returns string from the error, if nil return empty string
	ErrStr() string
}

// JErr implements JobError
type JErr struct {
	reason string
	err    error
}

func JError(reason string, err error) JErr {
	return JErr{reason: reason, err: err}
}

func (err JErr) Reason() string {
	return err.reason
}

func (err JErr) Err() error {
	return err.err
}

func (err JErr) ErrStr() string {
	if err.err != nil {
		return err.err.Error()
	}
	return ""
}
