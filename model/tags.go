package model

import (
	"github.com/afl-lxw/gin-trend/global"
	"runtime/metrics"
)

type Tags struct {
	global.TREND_MODEL
	Name string `json:"name"`
	metrics.Description
}
