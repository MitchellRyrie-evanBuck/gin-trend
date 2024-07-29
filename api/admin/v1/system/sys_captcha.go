package system

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CapApi struct {
}

func (b *CapApi) Captcha(c *gin.Context) {
	data, err := configSystemCaptchaService.GenerateCaptchaHandler(c)
	// 这里进行数据返回
	if err != nil {
		global.TREND_LOG.Error("查找失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)

	}
	response.Result(0, data, "操作成功", c)
}
