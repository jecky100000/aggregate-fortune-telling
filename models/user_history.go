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

type UserHistoryModel struct {
}

type UserHistory struct {
	BaseModel
	Type int    `json:"type"`
	Uid  int64  `json:"uid"`
	Cid  int64  `json:"cid"`
	Data string `json:"data"`
	Des  string `json:"des"`
}

func (UserHistory) TableName() string {
	return "sm_user_history"
}

func (con UserHistoryModel) GetAllPage(uid int64, vtype, page int) (res []UserHistory) {
	ay.Db.Debug().Where("uid = ? and type = ?", uid, vtype).Limit(10).Offset(page * 10).Find(&res)
	return
}

func (con UserHistoryModel) Save(uid int64, vtype int, cid int64, data, des string) bool {

	row := false

	var res UserHistory
	ay.Db.Where("uid = ? AND type = ?", uid, vtype).First(&res)

	if res.Id == 0 || vtype == 6 {
		if err := ay.Db.Create(&UserHistory{
			Type: vtype,
			Uid:  uid,
			Cid:  cid,
			Data: data,
			Des:  des,
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
