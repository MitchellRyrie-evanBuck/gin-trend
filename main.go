package main

import (
	"fmt"
	"github.com/afl-lxw/gin-trend/core"
	"github.com/afl-lxw/gin-trend/global"
)

func main() {
	fmt.Println("开始执行==》")
	global.GVA_VP = core.Viper() // 初始化Viper

	core.RunServer()
}
