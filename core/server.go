package core

import (
	"fmt"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routers()
	//address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)
	time.Sleep(10 * time.Microsecond)
	fmt.Printf("开始运行 运行地址为 http://127.0.0.1%s", address)

	//global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`默认自动化文档地址:http://127.0.0.1%s/swagger/index.html`, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
