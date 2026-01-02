package projects

import "errors"

var (
	ErrInvalidRequest = errors.New("잘못된 요청입니다.")
	ErrUnauthorized   = errors.New("인증되지 않은 사용자입니다.")
)
