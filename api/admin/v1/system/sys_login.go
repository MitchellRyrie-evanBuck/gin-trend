package system

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/model/common/response"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (b *CapApi) Login(c *gin.Context) {
	key := c.ClientIP()
	global.TREND_LOG.Info("IP来自", zap.String("ip---->", key))
	err := configSystemUserService.HandleSystemUser()
	if err != nil {

	}
	response.Ok(c)
}
