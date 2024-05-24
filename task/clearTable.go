package task

import (
	"errors"
	"fmt"
	"gin-vue-admin/model/common"
	"gorm.io/gorm"
	"time"
)

func ClearTable(db *gorm.DB) error {
	var ClearTableDetail []common.ClearDB
	ClearTableDetail = append(ClearTableDetail, common.ClearDB{
		TableName:    "sys_operation_records",
		CompareField: "created_at",
		Interval:     "2160h",
	})
	ClearTableDetail = append(ClearTableDetail, common.ClearDB{
		TableName:    "jwt_blacklists",
		CompareField: "created_at",
		Interval:     "168h",
	})
	if db == nil {
		return errors.New("DB Cannot Be Empty")
	}
	for _, detail := range ClearTableDetail {
		duration, err := time.ParseDuration(detail.Interval)
		if err != nil {
			return err
		}
		if duration < 0 {
			return errors.New("parse duration < 0")
		}
		err = db.Debug().Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", detail.TableName, detail.CompareField), time.Now().Add(-duration)).Error
		if err != nil {
			return err
		}
	}
	return nil
}
