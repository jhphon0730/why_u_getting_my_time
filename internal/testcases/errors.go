package testcases

import "errors"

var (
	ErrNotFound     = errors.New("테스트케이스를 찾을 수 없습니다.")
	ErrSameStatus   = errors.New("현재 상태와 동일한 상태로 변경하려고 합니다.")
	ErrUnauthorized = errors.New("권한이 없습니다.")

	ErrInvalidRequest         = errors.New("잘못된 요청입니다.")
	ErrNotInProjectTestStatus = errors.New("프로젝트에 속하지 않은 테스트케이스 상태 데이터입니다.")
	ErrNotInProjectMember     = errors.New("프로젝트에 속하지 않은 테스트케이스 할당자 데이터입니다.")
)
