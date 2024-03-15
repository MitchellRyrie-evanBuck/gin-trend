package utils

import (
	"errors"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/golang-jwt/jwt/v5"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

type JWT struct {
	SigningKey []byte
}

func NewJWT() *JWT {
	return &JWT{
		[]byte(global.TREND_CONFIG.JWT.SigningKey),
	}
}

// NewAccessToken 创建一个新的token
func (j *JWT) NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := accessToken.SignedString(j.SigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (j *JWT) NewRefreshToken(claims jwt.RegisteredClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := refreshToken.SignedString(j.SigningKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// CreateTokenByOldToken 旧token 换新token 使用归并回源避免并发问题
func (j *JWT) CreateTokenByOldToken(oldToken string, claims UserClaims) (string, error) {
	v, err, _ := global.TREND_Concurrency_Control.Do("JWT:"+oldToken, func() (interface{}, error) {
		return j.NewAccessToken(claims)
	})
	return v.(string), err
}

// ParseAccessToken 解析 token
func (j *JWT) ParseAccessToken(accessToken string) (*UserClaims, error) {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	if parsedAccessToken != nil {
		if claims, ok := parsedAccessToken.Claims.(*UserClaims); ok && parsedAccessToken.Valid {
			return claims, nil
		}
		return nil, TokenInvalid

	} else {
		return nil, TokenInvalid
	}

}

func (j *JWT) ParseRefreshToken(refreshToken string) *jwt.RegisteredClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	return parsedRefreshToken.Claims.(*jwt.RegisteredClaims)
}
