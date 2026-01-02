package utils

import (
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
