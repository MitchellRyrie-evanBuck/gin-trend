package global

import (
	"github.com/afl-lxw/gin-trend/config"
	"github.com/afl-lxw/gin-trend/utils/timer"
	"github.com/qiniu/qmgo"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"golang.org/x/sync/singleflight"
	"gorm.io/gorm"
	"sync"
)

var (
	TREND_DB     *gorm.DB
	TREND_DBList map[string]*gorm.DB
	TREND_REDIS  *redis.Client
	TREND_MONGO  *qmgo.QmgoClient
	TREND_CONFIG config.Server
	TREND_VP     *viper.Viper
	// TREND_LOG    *oplogging.Logger
	TREND_LOG                 *zap.Logger
	TREND_Timer               timer.Timer = timer.NewTimerTask()
	TREND_Concurrency_Control             = &singleflight.Group{}

	//BlackCache local_cache.Cache
	lock sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return TREND_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := TREND_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
