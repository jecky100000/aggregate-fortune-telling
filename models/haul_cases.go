/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

import "gin/ay"

type HaulCasesModel struct {
}

type HaulCases struct {
	Id    int64  `gorm:"primaryKey" json:"id"`
	Type  string `gorm:"column:type" json:"type"`
	Link  string `gorm:"column:link" json:"link"`
	Cover string `gorm:"column:cover" json:"cover"`
	Sort  string `gorm:"column:sort" json:"sort"`
	Name  string `gorm:"column:name" json:"name"`
}

func (HaulCases) TableName() string {
	return "sm_haul_cases"
}

func (con HaulCasesModel) GetType(types int) (res []HaulCases) {
	ay.Db.Order("sort asc").Find(&res, "type = ?", types)
	return
}
