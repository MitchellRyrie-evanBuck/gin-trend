package core

import (
	"fmt"
	"github.com/afl-lxw/gin-trend/global"
	"github.com/afl-lxw/gin-trend/initialize"
	"go.uber.org/zap"
	"moul.io/banner"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	if global.GVA_CONFIG.System.UseMultipoint || global.GVA_CONFIG.System.UseRedis {
		// 初始化redis服务
		initialize.Redis()
	}
	if global.GVA_CONFIG.System.UseMongo {
		err := initialize.Mongo.Initialization()
		if err != nil {
			zap.L().Error(fmt.Sprintf("%+v", err))
		}
	}
	// 从db加载jwt数据
	//if global.GVA_DB != nil {
	//	system.LoadAll()
	//}
	Router := initialize.Routers()

	Router.Static("/form-generator", "./resource/page")

	//address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer(address, Router)

	time.Sleep(10 * time.Microsecond)
	fmt.Println(banner.Inline("welcome to GO Trend."))
	global.GVA_LOG.Info("server run success on ", zap.String("address", address))
	fmt.Printf(`
	默认自动化文档地址:http://127.0.0.1%s/swagger/index.html
	开始运行 运行地址为 http://127.0.0.1%s 
	`,
		address, address)
	global.GVA_LOG.Error(s.ListenAndServe().Error())
}
