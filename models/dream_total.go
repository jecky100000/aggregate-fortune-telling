/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import (
	"gin/ay"
)

type DreamTotalModel struct {
}

type DreamTotal struct {
	Id      int `gorm:"primaryKey" json:"id"`
	DreamId int `gorm:"column:dream_id" json:"dream_id"`
	Num     int `gorm:"column:num" json:"num"`
}

func (DreamTotal) TableName() string {
	return "sm_dream_total"
}

func (con DreamTotalModel) GetTotal(id int) (res DreamTotal) {
	ay.Db.Where("dream_id", id).Find(&res)
	if res.Num == 0 {
		con.InsertTotal(&DreamTotal{DreamId: id, Num: 1})
	} else {
		con.UpdateTotal(&res)
	}
	return
}

func (con DreamTotalModel) UpdateTotal(res *DreamTotal) bool {
	res.Num = res.Num + 1
	if err := ay.Db.Save(res).Error; err != nil {
		return false
	} else {
		return true
	}
}

func (con DreamTotalModel) InsertTotal(res *DreamTotal) bool {

	if err := ay.Db.Create(res).Error; err != nil {
		return false
	} else {
		return true
	}
}
