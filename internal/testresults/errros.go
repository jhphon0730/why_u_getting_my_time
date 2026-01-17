package testresults

import "errors"

var (
	ErrNotFoundTestCase = errors.New("테스트케이스를 찾을 수 없습니다.")
	ErrInvalidResult    = errors.New("올바르지 않은 결과입니다.")
)
