package xerrors

import "errors"

var ErrInvalidCredetionals = errors.New("invalid credetionals provided")
var ErrExpiredToken = errors.New("JWT token is alredy expired")
var ErrInvalidToken = errors.New("JWT token is invalid")
