package utils

import (
	"github.com/afl-lxw/gin-trend/global"
	"time"

	//"github.com/afl-lxw/gin-trend/global"
	//systemReq "github.com/afl-lxw/gin-trend/model/system/request"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"

	"github.com/segmentio/ksuid"
	"net"
)

type UserClaims struct {
	First      string `json:"first"`
	Last       string `json:"last"`
	BufferTime int64
	jwt.RegisteredClaims
	BaseClaims
}

type BaseClaims struct {
	UUID        ksuid.KSUID
	ID          uint
	Username    string
	NickName    string
	AuthorityId uint
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

func (j *JWT) CreateClaims(baseClaims BaseClaims) UserClaims {
	bf, _ := ParseDuration(global.TREND_CONFIG.JWT.BufferTime)
	ep, _ := ParseDuration(global.TREND_CONFIG.JWT.ExpiresTime)
	claims := UserClaims{
		BaseClaims: baseClaims,
		BufferTime: int64(bf / time.Second), // 缓冲时间1天 缓冲时间内会获得新的token刷新令牌 此时一个用户会存在两个有效令牌 但是前端只留一个 另一个会丢失
		RegisteredClaims: jwt.RegisteredClaims{
			Audience:  jwt.ClaimStrings{"GVA"},                   // 受众
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000)), // 签名生效时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(ep)),    // 过期时间 7天  配置文件
			Issuer:    global.TREND_CONFIG.JWT.Issuer,            // 签名的发行者
		},
	}
	return claims
}

func GetClaims(c *gin.Context) (*UserClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseAccessToken(token)
	if err != nil {
		global.TREND_LOG.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

// GetUserID 从Gin的Context中获取从jwt解析出来的用户ID
func GetUserID(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.ID
		}
	} else {
		waitUse := claims.(*UserClaims)
		return waitUse.BaseClaims.ID
	}
}

// GetUserUuid 从Gin的Context中获取从jwt解析出来的用户UUID
func GetUserUuid(c *gin.Context) ksuid.KSUID {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ksuid.KSUID{}
		} else {
			return cl.UUID
		}
	} else {
		waitUse := claims.(*UserClaims)
		return waitUse.UUID
	}
}

// GetUserAuthorityId 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserAuthorityId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.AuthorityId
		}
	} else {
		waitUse := claims.(*UserClaims)
		return waitUse.AuthorityId
	}
}

// GetUserInfo 从Gin的Context中获取从jwt解析出来的用户角色id
func GetUserInfo(c *gin.Context) *UserClaims {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return nil
		} else {
			return cl
		}
	} else {
		waitUse := claims.(*UserClaims)
		return waitUse
	}
}

// GetUserName 从Gin的Context中获取从jwt解析出来的用户名
func GetUserName(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Username
		}
	} else {
		waitUse := claims.(*UserClaims)
		return waitUse.Username
	}
}
