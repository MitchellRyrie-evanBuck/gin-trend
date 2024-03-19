package system

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseSystemAreasAPI struct {
}

func (t *BaseSystemAreasAPI) GetSystemAreas(c *gin.Context) {
	key := c.ClientIP()
	global.TREND_LOG.Info("IP来自", zap.String("ip---->", key))
	data, err := configSystemUserService.SystemAreasServices(c)
	if err != nil {
		global.TREND_LOG.Error("查找失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	response.Result(0, data, "操作成功", c)

}
