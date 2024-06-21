package system

import (
	"context"
	"gin-vue-admin/global"
	"gin-vue-admin/model/system"
	"gin-vue-admin/utils"
	"go.uber.org/zap"
)

type JwtService struct {
}

// JsonInBlackList 拉黑jwt
func (j *JwtService) JsonInBlackList(jwtList system.JwtBlacklist) (err error) {
	err = global.GVA_DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

// ISBlackList 判断 jwt 是否在黑名单内部
func (j *JwtService) ISBlackList(jwt string) bool {
	_, ok := global.BlackCache.Get(jwt)
	return ok
}

// GetRedisJWT jwt 存入 redis 并设置过期时间
func (j *JwtService) GetRedisJWT(userName string) (redisJWT string, err error) {
	redisJWT, err = global.GVA_REDIS.Get(context.Background(), userName).Result()
	return redisJWT, err
}

// SetRedisJWT jwt 存入 redis 并设置过期时间
func (j *JwtService) SetRedisJWT(jwt, userName string) (err error) {
	// 此处过期时间等于 jwt 过期时间
	dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
	if err != nil {
		return err
	}
	timer := dr
	err = global.GVA_REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func LoadAll() {
	var data []string
	err := global.GVA_DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.GVA_LOG.Error("加载数据可 jwt 黑名单失败！", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		//jwt 加入黑名单 加入 BlackCache 中
		global.BlackCache.SetDefault(data[i], struct{}{})
	}
}
