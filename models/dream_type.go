/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import "aggregate-fortune-telling/ay"

type DreamTypeModel struct {
}

type DreamType struct {
	Id   int    `gorm:"primaryKey" json:"id"`
	Pid  int    `gorm:"column:pid" json:"pid"`
	Name string `gorm:"column:name" json:"name"`
}

func (DreamType) TableName() string {
	return "sm_dream_type"
}

func (con DreamTypeModel) GetAllType() (res []DreamType) {
	ay.Db.Find(&res)
	return
}

func (con DreamTypeModel) InsertType(res *DreamType) bool {

	if err := ay.Db.Create(res).Error; err != nil {
		return false
	} else {
		return true
	}
}
