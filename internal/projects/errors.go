package projects

import "errors"

var (
	ErrInvalidRequest = errors.New("잘못된 요청입니다.")
	ErrUnauthorized   = errors.New("인증되지 않은 사용자입니다.")

	ErrRecordNotFound = errors.New("사용자를 찾을 수 없습니다.")
	ErrAlreadyMember  = errors.New("이미 프로젝트에 참여한 사용자입니다.")

	ErrInvalidUserID    = errors.New("올바른 사용자 ID가 아닙니다.")
	ErrInvalidProjectID = errors.New("올바른 프로젝트 ID가 아닙니다.")

	ErrLastManager = errors.New("마지막 관리자는 삭제할 수 없습니다.")
	ErrNotMember   = errors.New("프로젝트에 참여하지 않은 사용자입니다.")
)
