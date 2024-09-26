package system

import (
	"context"
	sysModel "gin-vue-admin/model/system"
	"gin-vue-admin/service/system"
	"gin-vue-admin/utils"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

const initOrderAuthority = initOrderCasbin + 1

type initAuthority struct{}

func init() {
	system.RegisterInit(initOrderAuthority, &initAuthority{})
}

func (i *initAuthority) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&sysModel.SysAuthority{})
}

func (i *initAuthority) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&sysModel.SysAuthority{})
}

func (i initAuthority) InitializerName() string {
	return sysModel.SysAuthority{}.TableName()
}

func (i *initAuthority) InitializeData(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	entities := []sysModel.SysAuthority{
		{AuthorityID: 888, AuthorityName: "普通用户", ParentID: utils.Pointer[uint](0), DefaultRouter: "dashboard"},
		{AuthorityID: 9528, AuthorityName: "测试用户", ParentID: utils.Pointer[uint](0), DefaultRouter: "dashboard"},
		{AuthorityID: 88881, AuthorityName: "普通用户子角色", ParentID: utils.Pointer[uint](888), DefaultRouter: "dashboard"},
	}
	if err := db.Create(&entities).Error; err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败", sysModel.SysAuthority{}.TableName())
	}
	// data authority
	if err := db.Model(&entities[0]).Association("DataAuthorityID").Replace(
		[]*sysModel.SysAuthority{
			{AuthorityID: 888},
			{AuthorityID: 9528},
			{AuthorityID: 8881},
		}); err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败", db.Model(&entities[0]).Association("DataAuthorityID").Relationship.JoinTable.Name)
	}
	if err := db.Model(&entities[1]).Association("DataAuthorityID").Replace([]*sysModel.SysAuthority{
		{AuthorityID: 9528},
		{AuthorityID: 8881},
	}); err != nil {
		return ctx, errors.Wrapf(err, "%s表数据初始化失败", db.Model(&entities[1]).Association("DataAuthorityID").Relationship.JoinTable.Name)
	}
	next := context.WithValue(ctx, i.InitializerName(), entities)
	return next, nil
}

func (i *initAuthority) DataInserted(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	if errors.Is(db.Where("authority_id = ?", "8881").First(&sysModel.SysAuthority{}).Error, gorm.ErrRecordNotFound) { // 判断是否有数据
		return false
	}
	return true
}
