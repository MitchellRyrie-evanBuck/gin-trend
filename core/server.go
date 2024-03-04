package core

import (
	"github.com/afl-lxw/gin-trend/initialize"
	"time"
)

type server interface {
	ListenAndServe() error
}

func RunServer() {
	Router := initialize.Routers()
	//address := fmt.Sprintf(":%d", global.GVA_CONFIG.System.Addr)
	s := initServer("8888", Router)
	time.Sleep(10 * time.Microsecond)

	s.ListenAndServe()
}
