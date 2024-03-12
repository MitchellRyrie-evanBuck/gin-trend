package initialize

import (
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/utils"
)

func OtherInit() {
	dr, err := utils.ParseDuration(global.TREND_CONFIG.JWT.ExpiresTime)
	if err != nil {
		panic(err)
	}
	_, err = utils.ParseDuration(global.TREND_CONFIG.JWT.BufferTime)
	if err != nil {
		panic(err)
	}
	println("dr------------>", dr)
	//global.BlackCache = local_cache.NewCache(
	//	local_cache.SetDefaultExpire(dr),
	//)
}
