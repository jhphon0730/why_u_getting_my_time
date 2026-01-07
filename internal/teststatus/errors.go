package teststatus

import "errors"

var (
	ErrInvalidRequest = errors.New("잘못된 요청입니다.")

	ErrNotFoundTestCase = errors.New("테스트 케이스를 찾을 수 없습니다.")
)
