package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jhphon0730/action_manager/internal/auth"
	"github.com/jhphon0730/action_manager/internal/response"
)

const (
	USER_ID_CTX_KEY = iota
)

type UserIDCTXKey uint

// AuthMiddleware는 JWT 쿠키를 검증하고 userID를 context에 저장
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 쿠키에서 token 추출
		token, err := c.Cookie("token")
		if err != nil || token == "" {
			response.RespondError(c, http.StatusUnauthorized, "unauthorized")
			c.Abort()
			return
		}

		// JWT 검증
		claims, err := auth.ValidateJWTToken(token)
		if err != nil {
			response.RespondError(c, http.StatusUnauthorized, "invalid token")
			c.Abort()
			return
		}

		// userID를 gin.Context에 저장
		c.Set(UserIDCTXKey(USER_ID_CTX_KEY), claims.UserID)

		c.Next()
	}
}
