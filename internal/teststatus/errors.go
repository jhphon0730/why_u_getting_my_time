package teststatus

import "errors"

var (
	ErrNotFoundTestCase = errors.New("테스트 케이스를 찾을 수 없습니다.")
)
