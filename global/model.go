package global

import (
	"gorm.io/gorm"
	"time"
)

type GVA_MODEL struct {
	ID        uint           `json:"id" gorm:"primaryKey"` // 主键ID
	CreatedAt time.Time      // 创建时间
	UpdateAt  time.Time      `gorm:"type:DATETIME; default:NULL"` // 更新时间
	DeleteAt  gorm.DeletedAt `json:"-" gorm:"index"`              //删除时间
}
