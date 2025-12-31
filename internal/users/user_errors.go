package users

import "errors"

var (
	ErrInternal       = errors.New("내부 서버 오류")
	ErrDuplicateEmail = errors.New("이미 존재하는 이메일입니다.")
	ErrNotFound       = errors.New("사용자를 찾을 수 없습니다.")
	ErrInvalidRequest = errors.New("잘못된 요청입니다.")
)
