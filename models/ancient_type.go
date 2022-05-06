/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package models

type AncientTypeModel struct {
}

type AncientType struct {
	Id   int64  `gorm:"primaryKey" json:"id"`
	Pid  int64  `json:"pid"`
	Name string `json:"name"`
}

func (AncientType) TableName() string {
	return "sm_ancient_type"
}
