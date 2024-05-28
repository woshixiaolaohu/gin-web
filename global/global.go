package global

import (
	"gin-vue-admin/config"
	"gin-vue-admin/utils/timer"
	"github.com/songzhibin97/gkit/cache/local_cache"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	GVA_DB     *gorm.DB
	GVA_DBList map[string]*gorm.DB
	GVA_CONFIG config.Server
	GVA_VP     *viper.Viper
	BlackCache local_cache.Cache
	GVA_LOG    *zap.Logger
	GVA_Timer  timer.Timer = timer.NewTimerTask()
)
