package user

import (
	"fmt"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/model/admin/system"
	userReq "github.com/afl-lxw/gin-trend/model/app/user/request"
	systemRes "github.com/afl-lxw/gin-trend/model/app/user/response"
	"github.com/afl-lxw/gin-trend/model/common/response"
	"github.com/afl-lxw/gin-trend/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

type BaseAppUser struct {
}

// UserLogin 微信授权
func (t *BaseAppUser) UserLogin(c *gin.Context) {
	var l userReq.Login
	err := c.ShouldBindJSON(&l)
	key := c.ClientIP()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	fmt.Println("==============", key)
}

func (t *BaseAppUser) TokenNext(c *gin.Context, user system.SysUser) {
	j := &utils.JWT{SigningKey: []byte(global.TREND_CONFIG.JWT.SigningKey)} // 唯一签名

	claims := j.CreateClaims(utils.BaseClaims{
		UUID:        user.UUID,
		ID:          user.ID,
		NickName:    user.NickName,
		Username:    user.Username,
		AuthorityId: user.AuthorityId,
	})
	token, err := j.NewAccessToken(claims)
	if err != nil {
		global.TREND_LOG.Error("获取token失败!", zap.Error(err))
		response.FailWithMessage("获取token失败", c)
		return
	}

	if !global.TREND_CONFIG.System.UseMultipoint {
		utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
		response.OkWithDetailed(systemRes.LoginResponse{
			User:      user,
			Token:     token,
			ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
		}, "登录成功", c)
		return
	}

	//if jwtStr, err := jwtService.GetRedisJWT(user.Username); err == redis.Nil {
	//	if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
	//		global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
	//		response.FailWithMessage("设置登录状态失败", c)
	//		return
	//	}
	//	utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	//	response.OkWithDetailed(systemRes.LoginResponse{
	//		User:      user,
	//		Token:     token,
	//		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	//	}, "登录成功", c)
	//} else if err != nil {
	//	global.GVA_LOG.Error("设置登录状态失败!", zap.Error(err))
	//	response.FailWithMessage("设置登录状态失败", c)
	//} else {
	//	var blackJWT system.JwtBlacklist
	//	blackJWT.Jwt = jwtStr
	//	if err := jwtService.JsonInBlacklist(blackJWT); err != nil {
	//		response.FailWithMessage("jwt作废失败", c)
	//		return
	//	}
	//	if err := jwtService.SetRedisJWT(token, user.Username); err != nil {
	//		response.FailWithMessage("设置登录状态失败", c)
	//		return
	//	}
	//	utils.SetToken(c, token, int(claims.RegisteredClaims.ExpiresAt.Unix()-time.Now().Unix()))
	//	response.OkWithDetailed(systemRes.LoginResponse{
	//		User:      user,
	//		Token:     token,
	//		ExpiresAt: claims.RegisteredClaims.ExpiresAt.Unix() * 1000,
	//	}, "登录成功", c)
	//}
}

func (t *BaseAppUser) Register(c *gin.Context) {
	//validators.ValidateCreateUser()
}
