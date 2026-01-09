package testcases

import "errors"

var (
	ErrInvalidRequest = errors.New("잘못된 요청입니다.")
	ErrNotInProjectTestStatus = errors.New("프로젝트에 속하지 않은 테스트케이스 상태 데이터입니다.")
)
