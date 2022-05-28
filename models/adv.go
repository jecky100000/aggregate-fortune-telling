/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import "aggregate-fortune-telling/ay"

type AdvModel struct {
}

type Adv struct {
	BaseModel
	Image string `gorm:"column:image" json:"image"`
	Link  string `gorm:"column:link" json:"link"`
	Sort  int    `gorm:"column:sort" json:"sort"`
	Type  int    `gorm:"column:type" json:"type"`
}

func (Adv) TableName() string {
	return "sm_adv"
}

func (con AdvModel) GetType(val int) (res []Adv) {
	ay.Db.Where("type = ?", val).Order("sort asc").Find(&res)
	return
}
