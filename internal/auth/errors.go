package auth

import "errors"

var (
	ErrInvalidToken = errors.New("올바른 토큰이 아닙니다.")
)
