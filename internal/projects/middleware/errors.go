package middleware

import "errors"

var (
	UnauthorizedUser    = errors.New("인증되지 않은 사용자입니다.")
	UnauthorizedAuth    = errors.New("인증되지 않은 권한입니다.")
	UnauthorizedProject = errors.New("인증되지 않은 프로젝트입니다.")
	Unauthorized        = errors.New("인증되지 않은 사용자입니다.")
	PermissionRequired = errors.New("권한이 필요합니다.")
	ProjectNotFound = errors.New("프로젝트를 찾지 못했습니다.")
)
