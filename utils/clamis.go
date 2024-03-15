package utils

import (
	//"github.com/afl-lxw/gin-trend/global"
	//systemReq "github.com/afl-lxw/gin-trend/model/system/request"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"os"

	//"github.com/gofrs/uuid/v5"
	"net"
)

type UserClaims struct {
	Id    string `json:"id"`
	First string `json:"first"`
	Last  string `json:"last"`
	jwt.RegisteredClaims
}

func NewAccessToken(claims UserClaims) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func NewRefreshToken(claims jwt.RegisteredClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ParseAccessToken(accessToken string) *UserClaims {
	parsedAccessToken, _ := jwt.ParseWithClaims(accessToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	return parsedAccessToken.Claims.(*UserClaims)
}

func ParseRefreshToken(refreshToken string) *jwt.RegisteredClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("TOKEN_SECRET")), nil
	})

	return parsedRefreshToken.Claims.(*jwt.RegisteredClaims)
}

// ------------------------------------------------------------------------------------------------------

func ClearToken(c *gin.Context) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	c.SetCookie("x-token", "", -1, "/", host, false, false)
}

func SetToken(c *gin.Context, token string, maxAge int) {
	// 增加cookie x-token 向来源的web添加
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}
	c.SetCookie("x-token", token, maxAge, "/", host, false, false)
}

func GetToken(c *gin.Context) string {
	token, _ := c.Cookie("x-token")
	if token == "" {
		token = c.Request.Header.Get("x-token")
	}
	return token
}

//func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
//	token := GetToken(c)
//	j := NewJWT()
//	claims, err := j.ParseToken(token)
//	if err != nil {
//		global.GVA_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
//	}
//	return claims, err
//}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
//func GetUserID(c *gin.Context) uint {
//	if claims, exists := c.Get("claims"); !exists {
//		if cl, err := GetClaims(c); err != nil {
//			return 0
//		} else {
//			return cl.BaseClaims.ID
//		}
//	} else {
//		waitUse := claims.(*systemReq.CustomClaims)
//		return waitUse.BaseClaims.ID
//	}
//}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
//func GetUserUuid(c *gin.Context) uuid.UUID {
//	if claims, exists := c.Get("claims"); !exists {
//		if cl, err := GetClaims(c); err != nil {
//			return uuid.UUID{}
//		} else {
//			return cl.UUID
//		}
//	} else {
//		waitUse := claims.(*systemReq.CustomClaims)
//		return waitUse.UUID
//	}
//}
//
//// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
//func GetUserAuthorityId(c *gin.Context) uint {
//	if claims, exists := c.Get("claims"); !exists {
//		if cl, err := GetClaims(c); err != nil {
//			return 0
//		} else {
//			return cl.AuthorityId
//		}
//	} else {
//		waitUse := claims.(*systemReq.CustomClaims)
//		return waitUse.AuthorityId
//	}
//}
//
//// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
//func GetUserInfo(c *gin.Context) *systemReq.CustomClaims {
//	if claims, exists := c.Get("claims"); !exists {
//		if cl, err := GetClaims(c); err != nil {
//			return nil
//		} else {
//			return cl
//		}
//	} else {
//		waitUse := claims.(*systemReq.CustomClaims)
//		return waitUse
//	}
//}
//
//// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
//func GetUserName(c *gin.Context) string {
//	if claims, exists := c.Get("claims"); !exists {
//		if cl, err := GetClaims(c); err != nil {
//			return ""
//		} else {
//			return cl.Username
//		}
//	} else {
//		waitUse := claims.(*systemReq.CustomClaims)
//		return waitUse.Username
//	}
//}
