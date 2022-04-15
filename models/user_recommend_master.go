/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type UserRecommendMasterModel struct {
}

type UserRecommendMaster struct {
	BaseModel
	Uid      int64 `gorm:"column:uid" json:"uid"`
	MasterId int64 `gorm:"column:master_id" json:"master_id"`
}

func (UserRecommendMaster) TableName() string {
	return "sm_user_recommend_master"
}
