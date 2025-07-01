package xerrors

import "errors"

var ErrInvalidUserId = errors.New("invalid user ID provided")
var ErrUserNotFound = errors.New("user not found")
