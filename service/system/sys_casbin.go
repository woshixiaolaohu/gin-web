package system

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/system/request"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"sync"
)

type CasbinService struct{}

var CasbinServiceApp = new(CasbinService)

// UpdateCasbin
// @function: UpdateCasbin
// @description: 更新casbin权限
// @param: authorityId string, casbinInfos []request.CasbinInfo
// @return: error
func (casbinService *CasbinService) UpdateCasbin(AuthorityID uint, casbinInfos []request.CasbinInfo) error {
	authorityID := strconv.Itoa(int(AuthorityID))
	casbinService.ClearCasbin(0, authorityID)
	rules := [][]string{}
	// 权限去重
	deduplicateMap := make(map[string]bool)
	for _, v := range casbinInfos {
		key := authorityID + v.Path + v.Method
		if _, ok := deduplicateMap[key]; !ok {
			deduplicateMap[key] = true
			rules = append(rules, []string{authorityID, v.Path, v.Method})
		}
	}
	e := casbinService.Casbin()
	success, _ := e.AddPolicies(rules)
	if !success {
		return errors.New("存在相同API，添加失败，请联系管理员")
	}
	return nil
}

// UpdateCasbinApi
// @function: UpdateCasbinApi
// @description: API更新随动
// @param: oldPath string, newPath string, oldMethod string, newMethod string
// @return: error
func (casbinService *CasbinService) UpdateCasbinApi(oldPath, newPath, oldMethod, newMethod string) error {
	err := global.GVA_DB.Model(&gormadapter.CasbinRule{}).Where("v1 = AND v2 = ?", oldPath, oldMethod).Updates(map[string]interface{}{
		"v1": newPath,
		"v2": newMethod,
	}).Error
	e := casbinService.Casbin()
	err = e.LoadPolicy()
	if err != nil {
		return err
	}
	return err
}

// GetPolicyPathByAuthorityID
// @function: GetPolicyPathByAuthorityID
// @description: 获取权限列表
// @param: authorityId string
// @return: pathMaps []request.CasbinInfo
func (casbinService *CasbinService) GetPolicyPathByAuthorityID(AuthorityID uint) (pathMaps []request.CasbinInfo) {
	e := casbinService.Casbin()
	authorityID := strconv.Itoa(int(AuthorityID))
	list, _ := e.GetFilteredPolicy(0, authorityID)
	for _, v := range list {
		pathMaps = append(pathMaps, request.CasbinInfo{
			Path:   v[1],
			Method: v[2],
		})
	}
	return pathMaps
}

// ClearCasbin
// @function: ClearCasbin
// @description: 清除匹配的权限
// @param: v int, p ...string
// @return: bool
func (casbinService *CasbinService) ClearCasbin(v int, p ...string) bool {
	e := casbinService.Casbin()
	success, _ := e.RemoveFilteredPolicy(v, p...)
	return success

}

// RemoveFilteredPolicy
// @function: RemoveFilteredPolicy
// @description: 使用数据库方法清理筛选的politicy 此方法需要调用FreshCasbin方法才可以在系统中即刻生效
// @param: db *gorm.DB, authorityId string
// @return: error
func (casbinService *CasbinService) RemoveFilteredPolicy(db *gorm.DB, authorityID string) error {
	return db.Delete(&gormadapter.CasbinRule{}, "v0 = ? ", authorityID).Error
}

// SyncPolicy
// @function: SyncPolicy
// @description: 同步目前数据库的policy 此方法需要调用FreshCasbin方法才可以在系统中即刻生效
// @param: db *gorm.DB, authorityId string, rules [][]string
// @return: error
func (casbinService *CasbinService) SyncPolicy(db *gorm.DB, authorityID string, rules [][]string) error {
	err := casbinService.RemoveFilteredPolicy(db, authorityID)
	if err != nil {
		return err
	}
	return casbinService.AddPolicies(db, rules)
}

// AddPolicies
// @function: AddPolicies
// @description: 添加权限
// @param: db *gorm.DB, rules [][]string
// @return: error
func (casbinService *CasbinService) AddPolicies(db *gorm.DB, rules [][]string) error {
	var casbinRules []gormadapter.CasbinRule
	for i := range rules {
		casbinRules = append(casbinRules, gormadapter.CasbinRule{
			Ptype: "p",
			V0:    rules[i][0],
			V1:    rules[i][1],
			V2:    rules[i][2],
		})
	}
	return db.Create(&casbinRules).Error
}

// FreshCasbin
// @function: FreshCasbin
// @description: 刷新 Casbin
// @return: error
func (casbinService *CasbinService) FreshCasbin() (err error) {
	e := casbinService.Casbin()
	err = e.LoadPolicy()
	return err
}

var (
	syncedCachedEnforcer *casbin.SyncedCachedEnforcer
	once                 sync.Once
)

// Casbin
// @function: Casbin
// @description: 持久化到数据库  引入自定义规则
// @return: *casbin.SyncedCachedEnforcer
func (casbinService *CasbinService) Casbin() *casbin.SyncedCachedEnforcer {
	once.Do(func() {
		a, err := gormadapter.NewAdapterByDB(global.GVA_DB)
		if err != nil {
			zap.L().Error("适配数据库失败，请检查 casbin 表是否为InnoDB引擎！", zap.Error(err))
			return
		}
		text := `
		[request_definition]
		r = sub, obj, act
		
		[policy_definition]
		p = sub, obj, act
		
		[role_definition]
		g = _, _
		
		[policy_effect]
		e = some(where (p.eft == allow))
		
		[matchers]
		m = r.sub == p.sub && keyMatch2(r.obj,p.obj) && r.act == p.act
		`
		m, err := model.NewModelFromString(text)
		if err != nil {
			zap.L().Error("字符串加载模型失败！", zap.Error(err))
			return
		}
		syncedCachedEnforcer, _ = casbin.NewSyncedCachedEnforcer(m, a)
		syncedCachedEnforcer.SetExpireTime(60 * 60)
		_ = syncedCachedEnforcer.LoadPolicy()
	})
	return syncedCachedEnforcer
}
