package system

import (
	"gin-vue-admin/global"
	"gin-vue-admin/model/common/request"
	"gin-vue-admin/model/system"
)

type OperationRecordService struct {
}

// CreateSysOperationRecord 创建记录
func (o *OperationRecordService) CreateSysOperationRecord(sysOperationRecord system.SysOperationRecord) (err error) {
	err = global.GVA_DB.Create(&sysOperationRecord).Error
	return err
}

// DeleteSysOperationRecordByIds 批量删除记录
func (o *OperationRecordService) DeleteSysOperationRecordByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]system.SysOperationRecord{}, "id in (?)", ids.Ids).Error
	return err
}

// DeleteSysOperationRecord 删除操作记录
func (o *OperationRecordService) DeleteSysOperationRecord() {

}
