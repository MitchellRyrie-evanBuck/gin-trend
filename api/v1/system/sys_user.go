package system

import (
	"github.com/afl-lxw/gin-trend/model/common/response"
	"github.com/gin-gonic/gin"
)

func (b *CapApi) Login(c *gin.Context) {
	println("进入接口")
	err := systemUserConfigService.HandleSystemUser()
	if err != nil {

	}
	response.Ok(c)
}
