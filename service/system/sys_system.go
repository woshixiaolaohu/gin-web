package system

import (
	"gin-vue-admin/config"
	"gin-vue-admin/global"
	"gin-vue-admin/model/system"
	"gin-vue-admin/utils"
	"go.uber.org/zap"
)

type SystemConfigService struct{}

// GetSystemConfig
// @function: GetSystemConfig
// @description: 读取配置文件
// @return: conf config.Server, err error
func (systemConfigService *SystemConfigService) GetSystemConfig() (conf config.Server, err error) {
	return global.GVA_CONFIG, nil
}

// SetSystemConfig
// @function: SetSystemConfig
// @description: 设置配置文件
// @param: system model.System
// @return: err error
func (systemConfigService *SystemConfigService) SetSystemConfig(system system.System) (err error) {
	cs := utils.StructToMap(system.Config)
	for k, v := range cs {
		global.GVA_VP.Set(k, v)
	}
	err = global.GVA_VP.WriteConfig()
	return err
}

// GetServerInfo
// @function: GetServerInfo
// @description: 获取服务器信息
// @return: server *utils.Server, err error
func (systemConfigService *SystemConfigService) GetServerInfo() (server *utils.Server, err error) {
	var s utils.Server
	s.Os = utils.InitOS()
	if s.Cpu, err = utils.InitCPU(); err != nil {
		global.GVA_LOG.Error("func utils.InitCPU() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Ram, err = utils.InitRAM(); err != nil {
		global.GVA_LOG.Error("func utils.InitRAM() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	if s.Disk, err = utils.InitDisk(); err != nil {
		global.GVA_LOG.Error("func utils.InitDisk() Failed", zap.String("err", err.Error()))
		return &s, err
	}
	return &s, nil
}
