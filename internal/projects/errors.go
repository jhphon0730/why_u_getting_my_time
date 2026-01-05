package projects

import "errors"

var (
	ErrInvalidRequest = errors.New("잘못된 요청입니다.")
	ErrUnauthorized   = errors.New("인증되지 않은 사용자입니다.")

	ErrRecordNotFound = errors.New("사용자를 찾을 수 없습니다.")
	ErrAlreadyMember  = errors.New("이미 프로젝트에 참여한 사용자입니다.")
)
