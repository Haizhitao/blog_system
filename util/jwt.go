package util

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	Username string
	UserID   uint
}

// GenerateToken 生成JWT令牌
func GenerateToken(userID uint, username string, jwtSecret string) (string, error) {
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()), // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()), // 立即生效
			Issuer:    "haizhitao-blog-system",
		},
		Username: username,
		UserID:   userID,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}

// ParseToken 验证JWT令牌
func ParseToken(tokenString string, jwtSecret string) (*CustomClaims, error) {
	//解析Token
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		//验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})
	// 3. 处理解析错误
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, errors.New("token格式错误")
		} else if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, errors.New("token已过期")
		} else if errors.Is(err, jwt.ErrTokenNotValidYet) {
			return nil, errors.New("token尚未生效")
		}
		return nil, errors.New("token验证失败: " + err.Error())
	}

	// 4. 验证claims
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("无效的token")
}
