package contextutils

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/middleware"
)

// GetUserID 함수는 사용자 ID를 가져옵니다.
func GetUserID(c *gin.Context) (uint, bool) {
	userIDValue, exists := c.Get(middleware.UserIDCTXKey(middleware.USER_ID_CTX_KEY))
	if !exists {
		return 0, false
	}

	return userIDValue.(uint), true
}

// GetUserIDIDByParam 함수는 사용자 ID를 가져옵니다.
func GetUserIDIDByParam(c *gin.Context) (uint, bool) {
	userIDParam := c.Param("userID")
	userID, err := strconv.ParseUint(userIDParam, 10, 64)
	if err != nil {
		return 0, false
	}

	return uint(userID), true
}

// GetProjectIDByParam 함수는 프로젝트 ID를 가져옵니다.
func GetProjectIDByParam(c *gin.Context) (uint, bool) {
	projectIDParam := c.Param("projectID")
	projectID, err := strconv.ParseUint(projectIDParam, 10, 64)
	if err != nil {
		return 0, false
	}

	return uint(projectID), true
}
