/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import (
	"gin/ay"
	"time"
)

type UserMasterLogModel struct {
}

type UserMasterLog struct {
	BaseModel
	Uid      int64 `gorm:"column:uid" json:"uid"`
	MasterId int64 `gorm:"column:master_id" json:"master_id"`
}

func (UserMasterLog) TableName() string {
	return "sm_user_master_log"
}

func (con UserMasterLogModel) Save(uid, masterId int64) bool {

	row := false

	var res UserMasterLog
	ay.Db.Where("uid = ? AND master_id = ?", uid, masterId).First(&res)

	if res.Id == 0 {
		if err := ay.Db.Create(&UserMasterLog{
			Uid:      uid,
			MasterId: masterId,
		}).Error; err == nil {
			row = true
		}
	} else {
		stamp, _ := time.ParseInLocation("2006-01-02 15:04:05", time.Now().Format("2006-01-02 15:04:05"), time.Local)
		res.CreatedAt = MyTime{Time: stamp}
		if err := ay.Db.Save(&res).Error; err == nil {
			row = true
		}
	}
	return row
}
