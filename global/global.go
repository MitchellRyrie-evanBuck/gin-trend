package global

import (
	"gorm.io/gorm"
	"sync"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
	//GVA_REDIS  *redis.Client
	//GVA_MONGO  *qmgo.QmgoClient
	//GVA_CONFIG config.Server
	//GVA_VP     *viper.Viper
	//// GVA_LOG    *oplogging.Logger
	//GVA_LOG                 *zap.Logger
	//GVA_Timer               timer.Timer = timer.NewTimerTask()
	//GVA_Concurrency_Control             = &singleflight.Group{}
	//
	//BlackCache local_cache.Cache
	lock sync.RWMutex
)

// GetGlobalDBByDBName 通过名称获取db list中的db
func GetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	return GVA_DBList[dbname]
}

// MustGetGlobalDBByDBName 通过名称获取db 如果不存在则panic
func MustGetGlobalDBByDBName(dbname string) *gorm.DB {
	lock.RLock()
	defer lock.RUnlock()
	db, ok := GVA_DBList[dbname]
	if !ok || db == nil {
		panic("db no init")
	}
	return db
}
