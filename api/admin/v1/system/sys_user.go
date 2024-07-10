package system

import (
	"github.com/afl-lxw/gin-trend/model/common/response"
	"github.com/gin-gonic/gin"
)

func (b *CapApi) Login(c *gin.Context) error {
	println("进入接口")
	err := configSystemUserService.HandleSystemUser()
	if err != nil {

	}
	return response.Ok(c)
}
