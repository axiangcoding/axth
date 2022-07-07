package axth

import "errors"

var (
	ErrInternalFailed         = errors.New("axth internal failed")
	ErrUserNotExist           = errors.New("user not exist")
	ErrUserPasswordNotMatched = errors.New("user password not matched")
	ErrUserUpdateFailed       = errors.New("user profile update failed")
)
