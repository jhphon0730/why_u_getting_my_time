package auth

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/jhphon0730/action_manager/internal/config"
)

// TokenClaims는 JWT 토큰에 포함될 사용자 정보를 나타냅니다.
type TokenClaims struct {
	UserID uint `json:"user_id"`
	jwt.RegisteredClaims
}

// getJWTSecret는 JWT 시크릿을 안전하게 가져옵니다.
func getJWTSecret() []byte {
	cfg := config.GetConfig()
	return []byte(cfg.JWT_SECRET)
}

// GenerateJWTToken 함수는 JWT 토큰을 생성합니다.
func GenerateJWTToken(userID uint) (string, error) {
	claims := TokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(1 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(getJWTSecret())
}

// ValidateJWTToken 함수는 JWT 토큰을 검증합니다.
func ValidateJWTToken(tokenString string) (*TokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&TokenClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// 서명 알고리즘 검증
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, ErrInvalidToken
			}
			return getJWTSecret(), nil
		},
	)

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*TokenClaims)
	if !ok || !token.Valid {
		return nil, ErrInvalidToken
	}

	return claims, nil
}
