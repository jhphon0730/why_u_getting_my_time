package attachments

import "errors"

var (
	ErrInvalidTargetType  = errors.New("올바른 타입이 아닙니다.")
	ErrNoFilesProvided    = errors.New("파일이 제공되지 않았습니다.")
	ErrNotFoundTestCase   = errors.New("테스트 케이스를 찾을 수 없습니다.")
	ErrNotFoundTestResult = errors.New("테스트 결과를 찾을 수 없습니다.")
	ErrInvalidRequest     = errors.New("잘못된 요청입니다.")
)
