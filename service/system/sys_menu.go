package system

import (
	"errors"
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/system"
	"gorm.io/gorm"
)

type MenuService struct{}

var MenuServiceApp = new(MenuService)

// getMenuTreeMap
//
//	@function:		getMenuTreeMap
//	@description:	获取路由总树map
//	@param:			authorityID string
//	@return:		treeMap map[string][]system.SysMenu, err error
func (menuService *MenuService) getMenuTreeMap(authorityID uint) (treeMap map[uint][]system.SysMenu, err error) {
	var allMenus []system.SysMenu
	var baseMenu []system.SysBaseMenu
	var btns []system.SysAuthorityBtn
	treeMap = make(map[uint][]system.SysMenu)

	var SysAuthorityMenus []system.SysAuthorityMenu
	err = global.GVA_DB.Where("sys_authority_authority_id = ?", authorityID).Find(&SysAuthorityMenus).Error
	if err != nil {
		return
	}

	var MenuIds []string
	for i := range SysAuthorityMenus {
		MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuID)
	}
	err = global.GVA_DB.Where("id in (?)", MenuIds).Order("sort").Preload("Parameters").Find(&baseMenu).Error
	if err != nil {
		return
	}
	for i := range baseMenu {
		allMenus = append(allMenus, system.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityID: authorityID,
			MenuID:      baseMenu[i].ID,
			Parameters:  baseMenu[i].Parameters,
		})
	}
	err = global.GVA_DB.Where("authority_id = ?", authorityID).Preload("SysBaseMenuBtn").Find(&btns).Error
	if err != nil {
		return
	}
	var btnMap = make(map[uint]map[string]uint)
	for _, v := range btns {
		if btnMap[v.SysMenuID] == nil {
			btnMap[v.SysMenuID] = make(map[string]uint)
		}
		btnMap[v.SysMenuID][v.SysBaseMenuBtn.Name] = authorityID
	}
	for _, v := range allMenus {
		v.Btns = btnMap[v.SysBaseMenu.ID]
		treeMap[v.ParentID] = append(treeMap[v.ParentID], v)
	}
	return treeMap, err
}

// GetMenuTree
//
//	@function:		GetMenuTree
//	@description:	获取动态菜单树
//	@param:			authorityID string
//	@return:		menus []system.SysMenu, err error
func (menuService *MenuService) GetMenuTree(authorityID uint) (menus []system.SysMenu, err error) {
	menuTree, err := menuService.getMenuTreeMap(authorityID)
	menus = menuTree[0]
	for i := 0; i < len(menus); i++ {
		err = menuService.getChildrenList(&menus[i], menuTree)
	}
	return menus, err
}

// getChildrenList
//
//	@function:		getChildrenList
//	@description:	获取子菜单
//	@param:			menu *model.SysMenu, treeMap map[string][]model.SysMenu
//	@return:		err error
func (menuService *MenuService) getChildrenList(menu *system.SysMenu, treeMap map[uint][]system.SysMenu) (err error) {
	menu.Children = treeMap[menu.MenuID]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// GetInfoList
//
//	@function:		GetInfoList
//	@description:	获取路由分页
//	@return:		list interface{}, total int64,err error
func (menuService *MenuService) GetInfoList() (list interface{}, total int64, err error) {
	var menuList []system.SysBaseMenu
	treeMap, err := menuService.getBaseMenuTreeMap()
	menuList = treeMap[0]
	for i := 0; i < len(menuList); i++ {
		err = menuService.getBaseChildrenList(&menuList[i], treeMap)
	}
	return menuList, total, err
}

// getBaseChildrenList
//
//	@function:		getBaseChildrenList
//	@description:	获取菜单的子菜单
//	@param:			menu *model.SysBaseMenu, treeMap map[string][]model.SysBaseMenu
//	@return:		err error
func (menuService *MenuService) getBaseChildrenList(menu *system.SysBaseMenu, treeMap map[uint][]system.SysBaseMenu) (err error) {
	menu.Children = treeMap[menu.ID]
	for i := 0; i < len(menu.Children); i++ {
		err = menuService.getBaseChildrenList(&menu.Children[i], treeMap)
	}
	return err
}

// AddBaseMenu
//
//	@function:		AddBaseMenu
//	@description:	添加基础路由
//	@param:			menu model.SysBaseMenu
//	@return:		error
func (menuService *MenuService) AddBaseMenu(menu system.SysBaseMenu) error {
	if !errors.Is(global.GVA_DB.Where("name = ?", menu.Name).First(&system.SysBaseMenu{}).Error, gorm.ErrRecordNotFound) {
		return errors.New("存在重复name,请修改name")
	}
	return global.GVA_DB.Create(&menu).Error
}

// getBaseMenuTreeMap
//
//	@function:		getBaseMenuTreeMap
//	@description:	获取路由总树map
//	@return:		treeMap map[string][]system.SysBaseMenu, err error
func (menuService *MenuService) getBaseMenuTreeMap() (treeNap map[uint][]system.SysBaseMenu, err error) {
	var allMenus []system.SysBaseMenu
	treeNap = make(map[uint][]system.SysBaseMenu)
	err = global.GVA_DB.Order("sort").Preload("MenuBtn").Preload("Parameters").Find(&allMenus).Error
	for _, v := range allMenus {
		treeNap[v.ParentID] = append(treeNap[v.ParentID], v)
	}
	return treeNap, err
}

// GetBaseMenuTree
//
//	@function:		GetBaseMenuTree
//	@description:	获取基础路由树
//	@return:		menus []system.SysBaseMenu, err error
func (menuService *MenuService) GetBaseMenuTree() (menus []system.SysBaseMenu, err error) {
	treeMap, err := menuService.getBaseMenuTreeMap()
	menus = treeMap[0]
	for i := 0; i < len(menus); i++ {
		err = menuService.getBaseChildrenList(&menus[i], treeMap)
	}
	return menus, err
}

// AddMenuAuthority
//
//	@function:		AddMenuAuthority
//	@description:	为角色增加menu树
//	@param:			menus []model.SysBaseMenu, authorityID string
//	@return:		err error
func (menuService *MenuService) AddMenuAuthority(menus []system.SysBaseMenu, authorityID uint) (err error) {
	var auth system.SysAuthority
	auth.AuthorityID = authorityID
	auth.SysBaseMenus = menus
	err = AuthorityServiceApp.SetMenuAuthority(&auth)
	return err
}

// GetMenuAuthority
//
//	@function:		GetMenuAuthority
//	@description:	查看当前角色树
//	@param:			info *request.GetAuthorityId
//	@return:		menus []system.SysMenu, err error
func (menuService *MenuService) GetMenuAuthority(info *request.GetAuthorityId) (menus []system.SysMenu, err error) {
	var baseMenu []system.SysBaseMenu
	var SysAuthorityMenus []system.SysAuthorityMenu
	err = global.GVA_DB.Where("sys_authority_authority_id = ?", info.AuthorityID).Find(&SysAuthorityMenus).Error
	if err != nil {
		return
	}
	var MenuIds []string
	for i := range SysAuthorityMenus {
		MenuIds = append(MenuIds, SysAuthorityMenus[i].MenuID)
	}
	err = global.GVA_DB.Where("id in (?)", MenuIds).Order("sort").Find(&baseMenu).Error
	for i := range baseMenu {
		menus = append(menus, system.SysMenu{
			SysBaseMenu: baseMenu[i],
			AuthorityID: info.AuthorityID,
			MenuID:      baseMenu[i].ID,
			Parameters:  baseMenu[i].Parameters,
		})
	}
	return menus, err
}

// UserAuthorityDefaultRouter
//
//	@function:		UserAuthorityDefaultRouter
//	@description:	用户角色默认路由检查
//	@param:			user *system.SysUser
func (menuService *MenuService) UserAuthorityDefaultRouter(user *system.SysUser) {
	var menuIds []string
	err := global.GVA_DB.Model(&system.SysAuthorityMenu{}).Where("sys_authority_authority_id = ?", user.AuthorityID).Pluck("sys_base_menu_id", &menuIds).Error
	if err != nil {
		return
	}
	var am system.SysBaseMenu
	err = global.GVA_DB.First(&am, "name = ? and id in (?)", user.Authority.DefaultRouter, menuIds).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		user.Authority.DefaultRouter = "404"
	}
}
