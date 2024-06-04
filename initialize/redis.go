package initialize

import (
	"context"
	"gin-vue-admin/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

func Redis() {
	redisConfig := global.GVA_CONFIG.Redis
	var client redis.UniversalClient
	// 使用集群模式
	if redisConfig.UseCluster {
		client = redis.NewClusterClient(&redis.ClusterOptions{
			Addrs:    redisConfig.ClusterAddrs,
			Password: redisConfig.Password,
		})
	} else {
		// 使用单例模式
		client = redis.NewClient(&redis.Options{
			Addr:     redisConfig.Addr,
			Password: redisConfig.Password,
			DB:       redisConfig.DB,
		})
	}
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.GVA_LOG.Error("redis connect ping failed, err:", zap.Error(err))
		panic(err)
	} else {
		global.GVA_LOG.Info("redis connect ping response:", zap.String("pong", pong))
		global.GVA_REDIS = client
	}
}
