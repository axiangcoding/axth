package errors

import "errors"

var (
	ErrInternalFailed         = errors.New("axth Internal Failed")
	ErrUserNotExist           = errors.New("user not exist")
	ErrUserPasswordNotMatched = errors.New("user password not matched")
)
