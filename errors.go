package turbosms

import "github.com/errors"

const (
	// ErrEmptyUserName will return when provided username is empty
	ErrEmptyUserName = errors.Error("provided username is empty. It can't be an empty")
	// ErrEmptyPassword will return when provided password is empty
	ErrEmptyPassword = errors.Error("provided password is empty. It can't be an empty")
)
