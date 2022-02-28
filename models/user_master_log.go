/*
 * *
 *  @Author anderyly
 *  @email admin@aaayun.cc
 *  @link http://blog.aaayun.cc/
 *  @copyright Copyright (c) 2022
 *  *
 */

package models

type UserMasterLogModel struct {
}

type UserMasterLog struct {
	Id       int64 `gorm:"primaryKey" json:"id"`
	Uid      int64 `gorm:"column:uid" json:"uid"`
	MasterId int64 `gorm:"column:master_id" json:"master_id"`
}

func (UserMasterLog) TableName() string {
	return "sm_user_master_log"
}
